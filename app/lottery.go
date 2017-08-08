package app

const (
	IDE = iota
	RUN
	STOP
)

type Lottery struct {
	status int
	task   *Task
	finish chan bool
}

func New(count, roundTime, calculatTime, nextperiod int64, firstGroup, scondGroup string, selectRatio func()) *Lottery {
	return &Lottery{
		status: IDE,
		task:   TaskNew(count, roundTime, calculatTime, nextperiod, firstGroup, scondGroup, selectRatio),
	}
}

func (self *Lottery) start() {
	if self.status != IDE {
		return
	}
	self.status = RUN
	self.finish = make(chan bool)
	self.task.Start(self.finish)
}

func (self *Lottery) stop() {
	if self.status != STOP {
		return
	}
	self.status = STOP
	self.finish <- true
}
