package service

var RootGroup = new(Group)

type Group struct {
	GameService
	RecordService
	EventService
	StandardService
}
