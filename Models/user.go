package Models

type User struct {
	Name     string `json:name`
	Age      int    `json:age`
	Token    string `json:token`
	Password string `json:password`
}
