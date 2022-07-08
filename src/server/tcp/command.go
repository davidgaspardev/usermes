package tcp

type commandID int

const (
	CMD_REGISTER commandID = iota
	CMD_RECORD
)

type _Command struct {
	id     commandID
	client *_Client
	args   []string
}
