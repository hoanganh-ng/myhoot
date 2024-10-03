package user

type User struct {
	name    string
	isAdmin bool
}

func (u User) IsAdmin() bool {
	return u.isAdmin
}

func (u *User) ChangeName(newName string) {
	u.name = newName
}

func (u User) Name() string {
	return u.name
}
