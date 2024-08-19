package requests

type MessageRequest struct {
	Message string
	ToUserID uint
}

type ReadMessageRequest struct {
	FromUserID uint
	ToUserID uint
}