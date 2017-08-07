package main

import (
	"lotteryGame/app"
	// "github.com/robfig/cron"
)

// const OneSecond = 100*time.Second + 20*time.Millisecond

// func funcPanicRecovery() {
// 	cron := cron.New()
// 	cron.Start()
// 	defer cron.Stop()
// 	var count int64 = 0
// 	cron.AddFunc("*/1 * * * * *",
// 		func() {
// 			count++
// 			log.Println(60 - (count % 60))

// 			if 60-(count%60) == 55 {
// 				log.Println("计算一次")
// 			}
// 		})

// 	select {
// 	case <-time.After(OneSecond):
// 		log.Println("fuck")
// 		return
// 	}
// }

func main() {
	app.Run()

}
