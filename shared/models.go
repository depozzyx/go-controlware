package shared

type Command struct {
	Id  string `json:"id"`
	Cmd string `json:"command"`
}

var Version string = "v0.2.2"
