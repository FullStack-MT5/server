package benchttp

type User struct {
	ID    int64  `json:"-"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
