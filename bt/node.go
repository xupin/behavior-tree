package bt

type INode interface {
	Exec(IBlackboard) Status
}

type IAction interface {
	INode
	Enter()
	Leave()
}

type IBlackboard interface {
	Set(string, interface{})
	Get(string) interface{}
}

type Status int8

const (
	Success Status = iota
	Failure
)
