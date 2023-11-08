package bt

// 序列
type Sequence struct {
	children []INode
}

func NewSequence(children ...INode) *Sequence {
	return &Sequence{
		children: children,
	}
}

func (r *Sequence) Exec(db IBlackboard) Status {
	for _, child := range r.children {
		if result := child.Exec(db); result == Failure {
			return Failure
		}
	}
	return Success
}
