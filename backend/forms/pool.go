package forms

type PoolForm struct{}

func (f *PoolForm) Create(err error) string {
	// TODO
	return err.Error()
}

type CreatePoolForm struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	GroupID     int64    `json:"group_id" binding:"required"`
	IsAnonymous *bool    `json:"is_anonymous" binding:"required"`
	Options     []string `json:"options" binding:"required"`
	OpenFor     int64    `json:"open_for" binding:"required"`
}
