package bt

import (
	"sync"
)

type Parallel struct {
	children []INode
}

func NewParallel(children ...INode) *Parallel {
	return &Parallel{
		children: children,
	}
}

func (r *Parallel) Exec(db IBlackboard) Status {
	wg := sync.WaitGroup{}
	for _, child := range r.children {
		wg.Add(1)
		go func(node INode) {
			defer wg.Done()
			node.Exec(db)
		}(child)
	}
	wg.Wait()
	return Success
}
