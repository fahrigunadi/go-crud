package entities

type User struct {
	Id       int64
	Name     string `validate:"required"`
	Username string `validate:"required"`
	Phone    string `validate:"required" label:"Phone Number"`
}
