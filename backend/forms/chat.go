package forms

type ChatForm struct{}

func (f *ChatForm) WriteMessage(err error) string {
	// TODO
	return err.Error()
}

type WriteMessageForm struct {
	Text string `json:"text" binding:"required"`
}
