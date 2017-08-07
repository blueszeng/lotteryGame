package app

const (
	IDE = iota
	START
	RUN
	STOP
)

type lottery struct {
	status int
}

func (self *lottery) run() {

}

func (self *lottery) stop() {

}
