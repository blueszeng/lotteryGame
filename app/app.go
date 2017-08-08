package app

import "log"

var lottery *Lottery

func init() {
	lottery = New(0, 2, 1, 0, "110", "121", func() {
		log.Println("message")
	})
}

func Run() {
	lottery.start()
}
