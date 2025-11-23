package constant

type EventMeta struct {
	ID       int16
	Name     string
	Distance int32
}

var EventList = []EventMeta{
	{ID: 1, Name: "100m", Distance: 100},
	{ID: 2, Name: "200m", Distance: 200},
	{ID: 3, Name: "400m", Distance: 400},
	{ID: 4, Name: "800m", Distance: 800},
	{ID: 5, Name: "1000m", Distance: 1000},
	{ID: 6, Name: "1500m", Distance: 1500},
	{ID: 7, Name: "1600m", Distance: 1600},
	{ID: 8, Name: "3km", Distance: 3000},
	{ID: 9, Name: "4km", Distance: 4000},
	{ID: 10, Name: "5km", Distance: 5000},
	{ID: 11, Name: "10km", Distance: 10000},
	{ID: 12, Name: "半马", Distance: 21097},
	{ID: 13, Name: "30km", Distance: 30000},
	{ID: 14, Name: "全马", Distance: 42195},
}

// GameTypes 兼容现有代码
var GameTypes = func() map[int16]string {
	m := make(map[int16]string, len(EventList))
	for _, e := range EventList {
		m[e.ID] = e.Name
	}
	return m
}()

func GetGameTypeName(value int16) string {
	return GameTypes[value]
}

func IfGameTypeNotExist(gameType int16) bool {
	_, ok := GameTypes[gameType]
	return !ok
}
