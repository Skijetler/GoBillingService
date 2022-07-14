package biz

// User is a user model.
type User struct {
	ID        int64
	UserName  string
	FirstName string
	LastName  string
	Gender    bool
	Email     string
	Password  string
}
