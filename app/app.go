package app

import "log"

var lottery *Lottery

func init() {
	lottery = New(0, 10, 5, 50, "110", "121", func() {
		log.Println("message")
	})
}

func Run() {
	lottery.start()
}
