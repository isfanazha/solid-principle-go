package entity

type User struct {
	ID       int
	Username string
	Password string
}

type Profile struct {
	UserID       int
	FullName     string
	PictureURL   string
	Bio          string
	ContactEmail string
}
