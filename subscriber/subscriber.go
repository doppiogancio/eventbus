package subscriber

type (
	Subscriber interface {
		Name() string
		Receive(arg interface{})
	}
)
