package signal

type Effect struct {
	run     func()
	depends []SignalInterface
}

func NewEffect(run func(), signals ...SignalInterface) *Effect {
	e := &Effect{
		run:     run,
		depends: signals,
	}

	listener := func(any) {
		e.runEffect()
	}

	for _, s := range signals {
		s.Subscribe(listener)
	}

	e.runEffect()

	return e
}

func (e *Effect) runEffect() {
	e.run()
}
