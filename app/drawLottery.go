package app

type DrawLottery struct {
	FirstGroup  int64
	SecondGroup int64
	ThirdGroup  int64
	SelectRatio func()
}

func (self *DrawLottery) Draw() {
	// ratio := self.SelectRatio()

}
