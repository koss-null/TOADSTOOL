package queue

type (
	Queue interface {
		Push(interface{}) error
		Pop() interface{}
	}
)
