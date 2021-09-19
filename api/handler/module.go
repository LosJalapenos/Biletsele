package handler

import (
	"github.com/losjalapenos/biletsele/api/handler/connect"
	"github.com/losjalapenos/biletsele/api/handler/create_room"
	"github.com/losjalapenos/biletsele/api/handler/join_room"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	connect.Handler,
	create_room.Handler,
	join_room.Handler,
)
