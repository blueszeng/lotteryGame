package app

import (
	"fmt"
	"log"
)

import "github.com/robfig/cron"

type Task struct {
	Count        int64        // 计数器
	RoundTime    int64        // 一轮时长
	CalculatTime int64        //计算倒计时长
	drawLoty     *DrawLottery // 开奖
	Status       chan bool
}

func TaskNew(count, roundTime, calculatTime, firstGroup, scondGroup,
	thirdGroup int64, selectRatio func()) *Task {
	return &Task{
		Count:        count,
		RoundTime:    roundTime,
		CalculatTime: calculatTime,
		Status:       make(chan bool),
		drawLoty: &DrawLottery{
			FirstGroup:  firstGroup,
			SecondGroup: scondGroup,
			ThirdGroup:  thirdGroup,
			SelectRatio: selectRatio,
		},
	}
}

func (self *Task) Start() {
	cron := cron.New()
	cron.Start()
	defer cron.Stop()
	cron.AddFunc("*/1 * * * * *",
		func() {
			self.Count++
			countDown := (self.RoundTime - (self.Count % self.RoundTime))
			if countDown == self.RoundTime-self.CalculatTime {
				log.Println("计算一次")
				result := self.drawLoty.Draw()
				for key, value := range result {
					fmt.Println(key, value)
				}
			}
		})
	<-self.Status
}
