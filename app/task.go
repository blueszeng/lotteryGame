package app

import (
	"fmt"
	"log"
	"time"

	"lotteryGame/common/db"

	"github.com/robfig/cron"
)

type Task struct {
	Nextperiod   int64        // 下一期
	Count        int64        // 计数器
	RoundTime    int64        // 一轮时长
	CalculatTime int64        //计算倒计时长
	DrawLoty     *DrawLottery // 开奖
}

type PeriodLottery struct {
	Currperiod  string `db:"periodNo"`
	FirstGroup  string `db:"drawFirstNumber"`
	SecondGroup string `db:"drawSecondNumber"`
	ThirdGroup  string `db:"drawThirdNumber"`
}

func TaskNew(count, roundTime, calculatTime, nextperiod int64, firstGroup, scondGroup string, selectRatio func()) *Task {
	return &Task{
		Nextperiod:   nextperiod,
		Count:        count,
		RoundTime:    roundTime,
		CalculatTime: calculatTime,
		DrawLoty: &DrawLottery{
			FirstGroup:  firstGroup,
			SecondGroup: scondGroup,
			SelectRatio: selectRatio,
		},
	}
}

func (self *Task) Start(status chan bool) {

	cron := cron.New()
	cron.Start()
	defer cron.Stop()
	cron.AddFunc("*/1 * * * * *",
		func() {
			self.Count++
			countDown := (self.RoundTime - (self.Count % self.RoundTime))
			if countDown == self.RoundTime-self.CalculatTime {
				log.Println("计算一次")
				result := self.DrawLoty.Draw()

				currDate := time.Now().Format("20060102")
				plottery := &PeriodLottery{
					fmt.Sprintf("%s%d", currDate, self.Nextperiod),
					result["firstGroup"],
					result["secondGroup"],
					result["thirdGroup"],
				}
				writeDrawLotteryData(plottery)
				self.Nextperiod++
			}
		})
	<-status

}

func writeDrawLotteryData(plottry *PeriodLottery) {
	db.GetDb().MustExec(
		"INSERT INTO drawlotterys (periodNo, drawFirstNumber, drawSecondNumber, drawThirdNumber) VALUES (?, ?, ?, ?)",
		plottry.Currperiod, plottry.FirstGroup, plottry.SecondGroup, plottry.ThirdGroup)

}
