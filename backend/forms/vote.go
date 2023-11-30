package forms

type VoteForm struct{}

func (f *VoteForm) Create(err error) string {
	// TODO
	return err.Error()
}

type CreateVoteForm struct {
	ID int64 `json:"id" binding:"required,min=1"`
}
