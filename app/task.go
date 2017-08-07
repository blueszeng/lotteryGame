package app

import (
	"log"

	"github.com/robfig/cron"
)

type Task struct {
	Count        int64 // 计数器
	RoundTime    int64 // 一轮时长
	CalculatTime int64 //计算倒计时长
	drawLottery  DrawLottery
}

func New(count, roundTime, calculatTime int64) *Task {
	return &Task{
		Count:        count,
		RoundTime:    roundTime,
		CalculatTime: calculatTime,
	}
}

func (self *Task) Start() {
	cron.AddFunc("*/1 * * * * *",
		func() {
			self.Count++
			countDown := (self.RoundTime - (self.Count % self.RoundTime))
			if countDown == self.RoundTime-self.CalculatTime {
				log.Println("计算一次")
				self.drawLottery.Draw()
			}
		})

}
