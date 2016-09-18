package core

var registered []func() = []func(){}

func Register(fn ...func()) {
	registered = append(registered, fn...)
}

func ExecRegistred() {
	for _, fn := range registered {
		fn()
	}
}
