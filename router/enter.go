package router

import "run/api"

var RootGroup = new(Group)

type Group struct {
	Game   GameRouter
	Record RecordRouter
}

var (
	gameApi   = api.RootGroup.GameApi
	recordApi = api.RootGroup.RecordApi
)
