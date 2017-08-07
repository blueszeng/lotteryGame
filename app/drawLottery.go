package app

type DrawLottery struct {
	FirstGroup  int64
	SecondGroup int64
	ThirdGroup  int64
	SelectRatio func()
}

func (self *DrawLottery) Draw() map[string]int64 {
	self.SelectRatio()
	result := map[string]int64{}
	result["firstGroup"] = self.FirstGroup
	result["secondGroup"] = self.SecondGroup
	result["thirdGroup"] = self.ThirdGroup
	return result
}
