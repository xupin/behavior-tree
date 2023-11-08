package bt

type Blackboard struct {
	data map[string]interface{}
}

func NewBlackboard() *Blackboard {
	return &Blackboard{
		data: make(map[string]interface{}, 0),
	}
}

func (r *Blackboard) Set(key string, value interface{}) {
	r.data[key] = value
}

func (r *Blackboard) Get(key string) interface{} {
	return r.data[key]
}
