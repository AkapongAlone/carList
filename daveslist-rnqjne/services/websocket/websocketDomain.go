package websocket

type ServiceWebsocketDomain interface {
	BroadcastMessage(msg interface{}, stationID int) error
}
