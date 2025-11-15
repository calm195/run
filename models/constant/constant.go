package constant

const (
	ERROR   = 7
	SUCCESS = 0
)

const (
	OneHundred = iota
	TwoHundred
	FourHundred
	EightHundred
	OneThousand
	FifteenHundred
	SixteenHundred
	ThreeThousand
	FourThousand
	FiveThousand
	TenThousand
	HalfMarathon
	ThirtyThousand
	Marathon
)

var GameTypes = map[int16]string{
	OneHundred:     "100m",
	TwoHundred:     "200m",
	FourHundred:    "400m",
	EightHundred:   "800m",
	OneThousand:    "1km",
	FifteenHundred: "1500m",
	SixteenHundred: "1600m",
	ThreeThousand:  "3km",
	FourThousand:   "4km",
	FiveThousand:   "5km",
	TenThousand:    "10km",
	HalfMarathon:   "半马",
	ThirtyThousand: "30km",
	Marathon:       "全马",
}

func GetGameTypeName(value int16) string {
	return GameTypes[value]
}

func IfGameTypeNotExist(gameType int16) bool {
	_, ok := GameTypes[gameType]
	return !ok
}
