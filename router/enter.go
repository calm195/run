package router

import "run/api"

var RootGroup = new(Group)

type Group struct {
	Game GameRouter
}

var (
	gameApi = api.RootGroup.GameApi
)
