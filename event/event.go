package event

type (
	Event interface {
		Name() string
		Fire() interface{}
	}
)
