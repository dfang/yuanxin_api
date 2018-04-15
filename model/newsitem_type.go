package model

type NewsItemType int

const (
	Zero NewsItemType = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
)

func (s NewsItemType) String() string {
	return [...]string{"所有", "产业报道", "厂商动态", "数码相机/摄像机 ", "智能家电", "智能手机", "电脑", ""}[s]
}

func (s NewsItemType) Name() NewsItemType {
	return s
}

func (s NewsItemType) Value() int {
	return int(s)
}
