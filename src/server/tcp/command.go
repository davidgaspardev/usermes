package tcp

type commandID int

const (
	CMD_RECORD commandID = iota
)

type _Command struct {
	id     commandID
	client *_Client
	args   []string
}
