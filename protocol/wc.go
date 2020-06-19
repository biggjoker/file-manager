package protocol

type CmdInfo struct {
	Path  string `json:"path"`
	Cmd string   `json:"cmd"`
}

type CmdRespone struct {
	Body  string `json:"body"`
}