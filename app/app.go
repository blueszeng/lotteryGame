package app

import (
	"log"
	"lotteryGame/common/db"
	"strconv"
)

var lottery *Lottery
var rateConfig map[int64]string

func init() {

	// init rate config table
	rateConfig = map[int64]string{}
	var rateData []struct {
		Id         int64  `db:"id"`
		SelectName string `db:"selectName"`
	}
	db.GetDb().Select(&rateData, "SELECT id, selectName from rates")
	if len(rateData) == 0 {
		return
	}
	for i := 0; i < len(rateData); i++ {
		rateConfig[rateData[i].Id] = rateData[i].SelectName
	}
	var configData struct {
		Id               int    `db:"id"`
		PeriodNo         string `db:"periodNo"`
		DrawFirstNumber  string `db:"drawFirstNumber"`
		DrawSecondNumber string `db:"drawSecondNumber"`
	}
	db.GetDb().Get(&configData, "select  id, periodNo, drawFirstNumber, drawSecondNumber from drawlotterys order by id desc limit 0,1")
	log.Println(configData)
	if configData.Id < 0 {
		configData.PeriodNo = "20180101000"
		configData.DrawFirstNumber = "000"
		configData.DrawSecondNumber = "000"
	}
	PeriodNoString := configData.PeriodNo[8:]
	PeriodNo, error := strconv.ParseInt(PeriodNoString, 10, 64)
	if error != nil {
		log.Fatal("字符串转换成整数失败")
	}
	log.Println(len(configData.PeriodNo), PeriodNoString, PeriodNo)
	// init game config
	lottery = New(0, 60, 5, PeriodNo+1, configData.DrawFirstNumber, configData.DrawSecondNumber, func() {
		log.Println("message")
	})
}

func Run() {
	lottery.start()
}
