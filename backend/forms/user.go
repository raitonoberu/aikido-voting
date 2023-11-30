package forms

type UserForm struct{}

func (f *UserForm) Login(err error) string {
	// TODO
	return err.Error()
}

func (f *UserForm) Register(err error) string {
	// TODO
	return err.Error()
}

func (f *UserForm) Update(err error) string {
	// TODO
	return err.Error()
}

type LoginForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=3,max=50"`
}

type RegisterForm struct {
	Name     string `json:"name" binding:"required,min=3,max=50,name"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=3,max=50"`
}

type UpdateUserForm struct {
	// all fields are optional
	Name     string `json:"name" binding:"omitempty,min=3,max=50,name"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty,min=3,max=50"`
}
