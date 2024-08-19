package websocket

import (
	"Daveslist/responses"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections by default
		return true
	},
}

type Station struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

var stations = make(map[int]*Station)
var stationsMu sync.Mutex

type serviceWebsocket struct{}

func NewServiceWebsocket() *serviceWebsocket {
	return &serviceWebsocket{}
}

func (s *serviceWebsocket) HandleConnectionsWebsocket(c *gin.Context) {
	stationID, err := strconv.Atoi(c.Params.ByName("stationID"))
	if err != nil {
		newError := responses.NewAppErr(400, err.Error())
		c.JSON(http.StatusBadRequest, newError)
		return
	}
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}

	stationsMu.Lock()
	if stations[stationID] == nil {
		stations[stationID] = &Station{
			clients: make(map[*websocket.Conn]bool),
		}
	}
	station := stations[stationID]
	station.mu.Lock()
	station.clients[ws] = true
	station.mu.Unlock()
	stationsMu.Unlock()

	defer func() {
		station.mu.Lock()
		delete(station.clients, ws)
		station.mu.Unlock()
		ws.Close()
	}()

	s.readMessages(ws, stationID)
}

type JsonMessage struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func (s *serviceWebsocket) readMessages(ws *websocket.Conn, stationID int) {
	v := 0
	for {
		var msg JsonMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println("Error reading JSON:", err)
			}
			break
		}
		v += 1
		msg.Type = fmt.Sprintf("%s %v", msg.Type, v)
		fmt.Println("Received message:", msg)

		s.BroadcastMessage(msg, stationID)
	}
}

func (s *serviceWebsocket) BroadcastMessage(msg interface{}, stationID int) error {
	stationsMu.Lock()
	station, ok := stations[stationID]
	stationsMu.Unlock()

	if !ok {
		errStr := fmt.Sprintf(`websocket %d station not found`, stationID)
		return errors.New(errStr)
	}

	station.mu.Lock()
	defer station.mu.Unlock()

	for client := range station.clients {
		err := client.WriteJSON(msg)
		if err != nil {
			client.Close()
			delete(station.clients, client)
			return err
		}
	}

	return nil
}
