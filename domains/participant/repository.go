package participant

type UserRepository interface {
	Get(string) (*User, error)
}
