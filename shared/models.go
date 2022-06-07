package shared

type Command struct {
	Id  int64  `json:"id"`
	Cmd string `json:"command"`
}
