package responses

import "time"

type MessageResp struct {
	Message    string
	FromUserID uint
	ToUserID   uint
	CreatedAt  time.Time
}
