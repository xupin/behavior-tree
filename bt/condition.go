package bt

type Condition struct {
	f func(IBlackboard) bool
}

func NewCondition(f func(IBlackboard) bool) *Condition {
	return &Condition{
		f: f,
	}
}

func (r *Condition) Exec(db IBlackboard) Status {
	if v := r.f(db); !v {
		return Failure
	}
	return Success
}
