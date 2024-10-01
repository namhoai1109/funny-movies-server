package link

import (
	dbutil "funnymovies/util/db"
	websocketutil "funnymovies/util/websocket"

	"gorm.io/gorm"
)

type Link struct {
	db             *gorm.DB
	linkRepository LinkRepository
	ws             WebSocket
}

func New(
	db *gorm.DB,
	linkRepository LinkRepository,
	ws WebSocket,
) *Link {
	return &Link{
		db:             db,
		linkRepository: linkRepository,
		ws:             ws,
	}
}

type LinkRepository interface {
	dbutil.Intf
}

type WebSocket interface {
	BroadcastMessage(msg websocketutil.Message)
}
