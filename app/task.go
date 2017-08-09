package app

import (
	"fmt"
	"log"
	"lotteryGame/common/db"
	"lotteryGame/common/redis"
	"time"

	"github.com/robfig/cron"
)

type Task struct {
	Nextperiod   int64        // 下一期
	Counter      int64        // 计数器
	RoundTime    int64        // 一轮时长
	CalculatTime int64        //计算倒计时长
	DrawLoty     *DrawLottery // 开奖
}

type PeriodLottery struct {
	Currperiod  string `db:"periodNo"`
	FirstGroup  string `db:"drawFirstNumber"`
	SecondGroup string `db:"drawSecondNumber"`
	ThirdGroup  string `db:"drawThirdNumber"`
	Result      string `db:"result"`
}

func TaskNew(count, roundTime, calculatTime, nextperiod int64, firstGroup, scondGroup string, selectRatio func()) *Task {
	return &Task{
		Nextperiod:   nextperiod,
		Counter:      count,
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
			self.Counter++
			updateCounter(self.Counter) // 更新定时器
			countDown := (self.RoundTime - (self.Counter % self.RoundTime))
			if countDown == self.RoundTime-self.CalculatTime {
				result := self.DrawLoty.Draw() //生成投注组
				currDate := time.Now().Format("20060102")
				value, err := sizeScensResult(result["thirdGroup"])
				if err != nil {
					return
				}
				plottery := &PeriodLottery{
					fmt.Sprintf("%s%d", currDate, self.Nextperiod),
					result["firstGroup"],
					result["secondGroup"],
					result["thirdGroup"],
					value,
				}
				writeDrawLotteryData(plottery)
				calculateUserBets(plottery)
				self.Nextperiod++
			}
		})
	<-status

}

func updateCounter(counter int64) {
	err := redis.GetRedis().Set("global_counter", counter, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func calculateUserBets(plottry *PeriodLottery) {
	var BetData []struct {
		Id         int64   `db:"id"`
		UserId     int64   `db:"userId"`
		RateId     int64   `db:"rateId"`
		PeriodNo   string  `db:"periodNo"`
		Money      int     `db:"money"`
		Win        float64 `db:"win"`
		State      int64   `db:"state"`
		Ratio      float64 `db:"ratio"`
		UpdateTime string  `db:"updateTime"`
	}
	// 读取出来
	err := db.GetDb().Select(&BetData, "SELECT bets.id as id,  userId,  bets.periodNo as periodNo, money, win, state ratio,"+
		"updateTime FROM bets JOIN rates ON bets.rateId = rates.id WHERE bets.periodNo = ? ORDER BY bets.updateTime desc ", "201708090")
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(BetData) == 0 {
		return
	}
	// 计算结果
	for i := 0; i < len(BetData); i++ {
		buySelect, err := rateConfig[BetData[i].RateId]
		if err == false {
			continue
		}

		tdLen := len(plottry.ThirdGroup)
		selectLen := len(buySelect)
		if selectLen >= 3 {
			result, err := sizeScensResult(plottry.ThirdGroup)
			if err != nil {
				continue
			}
			if result == buySelect {
				BetData[i].Win = float64(BetData[i].Money) * BetData[i].Ratio
				BetData[i].State = 1 // 中奖
			}

		} else {
			compThiredGroup := plottry.ThirdGroup[(tdLen - selectLen):]
			fmt.Println(compThiredGroup, buySelect)
			if compThiredGroup == buySelect {
				BetData[i].Win = float64(BetData[i].Money) * BetData[i].Ratio
				BetData[i].State = 1 // 中奖
			}
		}

		BetData[i].State = 2 //没中奖
	}
	// 写回数据库
	tx := db.GetDb().MustBegin()
	for i := 0; i < len(BetData); i++ {
		tx.MustExec("UPDATE bets SET win = ?, State = ? WHERE id = ?", BetData[i].Win, BetData[i].State, BetData[i].Id)
		if BetData[i].State == 1 { //中奖才更改金钱
			tx.MustExec("UPDATE users SET money = money + ? WHERE id = ?", BetData[i].Win, BetData[i].UserId)
		}
	}
	tx.Commit()

}

func writeDrawLotteryData(plottry *PeriodLottery) {
	result := db.GetDb().MustExec(
		"INSERT INTO drawlotterys (periodNo, drawFirstNumber, drawSecondNumber, drawThirdNumber, result) VALUES (?, ?, ?, ?, ?)",
		plottry.Currperiod, plottry.FirstGroup, plottry.SecondGroup, plottry.ThirdGroup, plottry.Result)
	_, err := result.RowsAffected()
	if err != nil {
		log.Fatal("insert error", err)
	}

	key := "global_drawLottery_push"
	pipe := redis.GetRedis().Pipeline()
	pipe.LPush(key, plottry.Currperiod)
	mapValue := map[string]interface{}{
		"periodNo":         plottry.Currperiod,
		"drawFirstNumber":  plottry.FirstGroup,
		"drawSecondNumber": plottry.SecondGroup,
		"drawThirdNumber":  plottry.ThirdGroup,
		"result":           plottry.Result,
	}
	pipe.HMSet(plottry.Currperiod, mapValue)
	_, err = pipe.Exec()
	if err != nil {
		log.Fatal(err)
	}

}
