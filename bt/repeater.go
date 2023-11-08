package bt

type Repeater struct {
	child INode
	count int8
}

func NewRepeater(child INode, count int8) *Repeater {
	return &Repeater{
		child: child,
		count: count,
	}
}

func (r *Repeater) Exec(db IBlackboard) Status {
	for r.count > 0 {
		r.child.Exec(db)
		r.count -= 1
	}
	return Success
}
