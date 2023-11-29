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

type LoginForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=3,max=50"`
}

type RegisterForm struct {
	Name     string `json:"name" binding:"required,min=3,max=20,name"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=3,max=50"`
}
