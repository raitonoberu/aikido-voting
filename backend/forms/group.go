package forms

type GroupForm struct{}

func (f *GroupForm) Create(err error) string {
	// TODO
	return err.Error()
}

func (f *GroupForm) Update(err error) string {
	// TODO
	return err.Error()
}

func (f *GroupForm) AddUser(err error) string {
	// TODO
	return err.Error()
}

type CreateGroupForm struct {
	Name string `json:"name" binding:"required"`
}

type UpdateGroupForm struct {
	Name string `json:"name"`
}

type AddUserForm struct {
	ID int64 `json:"id" binding:"required"`
}
