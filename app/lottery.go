package app

import "log"

const (
	IDE = iota
	RUN
	STOP
)

type Lottery struct {
	status   int
	task     *Task
	runState chan bool
}

func New(count, roundTime, calculatTime, firstGroup, scondGroup,
	thirdGroup int64, selectRatio func()) *Lottery {
	return &Lottery{
		status:   IDE,
		task:     TaskNew(count, roundTime, calculatTime, firstGroup, scondGroup, thirdGroup, selectRatio),
		runState: make(chan bool),
	}
}

func (self *Lottery) run() {
	if self.status != IDE {
		return
	}
	self.status = RUN
	go func() {
		self.task.Start()
	}()
	select {
	case <-self.runState:
		log.Println("fuck")
		return
	}
}

func (self *Lottery) stop() {
	if self.status != STOP {
		return
	}
	self.status = STOP
	self.task.Status <- true
	defer func() {
		self.runState <- true
	}()
}
