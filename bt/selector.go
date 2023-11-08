package bt

type Selector struct {
	children []INode
}

func NewSelector(children ...INode) *Selector {
	return &Selector{
		children: children,
	}
}

func (r *Selector) Exec(db IBlackboard) Status {
	for _, child := range r.children {
		if result := child.Exec(db); result == Success {
			return Success
		}
	}
	return Failure
}
