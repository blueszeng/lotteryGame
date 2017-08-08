package app

import (
	"lotteryGame/utils"
)

type DrawLottery struct {
	FirstGroup  string
	SecondGroup string
	ThirdGroup  string
	SelectRatio func()
}

func (self *DrawLottery) Draw() map[string]string {
	self.SelectRatio()
	result := map[string]string{}
	self.FirstGroup = utils.RandomNumberString(3)
	self.SecondGroup = utils.RandomNumberString(3)
	self.ThirdGroup = utils.RandomNumberString(3)

	result["firstGroup"] = self.FirstGroup
	result["secondGroup"] = self.SecondGroup
	result["thirdGroup"] = self.ThirdGroup
	return result
}
