package entity

type User struct {
	Uid      string `json:"uid"`
	Gid      string `json:"gid"`
	Username string `json:"username"`
	Name     string `json:"name"`
	HomeDir  string `json:"home_dir"`
}
