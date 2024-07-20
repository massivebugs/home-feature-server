package handler

import "database/sql"

type Handlers struct {
	PingHandler   *PingHandler
	RepeatHandler *RepeatHandler
	AuthHandler   *AuthHandler
}

func NewAPIHandlers(db *sql.DB) *Handlers {
	return &Handlers{
		PingHandler:   NewPingHandler(),
		RepeatHandler: NewRepeatHandler(),
		AuthHandler:   NewAuthHandler(db),
	}
}
