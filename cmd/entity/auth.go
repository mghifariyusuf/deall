package entity

type Authorization struct {
	User    User   `json:"user"`
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}
