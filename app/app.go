package app

import "log"

var lottery *Lottery

func init() {
	lottery = New(0, 60, 5, 110, 121, 131, func() {
		log.Println("message")
	})
}

func Run() {
	lottery.run()

}
