package main

import (
	// "fmt"
	"lotteryGame/app"
	// "lotteryGame/common/db"
	// "lotteryGame/common/redis"
	// "github.com/robfig/cron"
)

func main() {
	// redisClient := redis.GetRedis()
	// pong, err := redisClient.Ping().Result()
	// fmt.Println(pong, err)
	// err = redisClient.Set("key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := redisClient.Get("key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	// mysqlDb := db.GetDb()

	// // rows, err := mysqlDb.Query("SELECT * FROM rates")
	// var rateData []struct {
	// 	Id         int64   `db:"id"`
	// 	SceneId    int64   `db:"sceneId"`
	// 	Select     int64   `db:"select"`
	// 	Ratio      float64 `db:"ratio"`
	// 	SelectName string  `db:"selectName"`
	// }
	// err = mysqlDb.Select(&rateData, "SELECT * FROM rates")
	// fmt.Println(rateData)
	app.Run()

}
