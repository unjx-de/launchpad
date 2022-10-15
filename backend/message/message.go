package message

type Response struct {
	Message string `json:"message" validate:"required"`
}

type (
	Responses uint
)

const (
	MissingInformation Responses = iota
	CannotCreate
	CannotUpdate
	CannotDelete
	CannotFind
	StillInUse
	CannotOpen
	CannotStart
	CannotGet
	NotLoggedIn
	CannotProcess
)

func (e Responses) String() string {
	switch e {
	case MissingInformation:
		return "missing information"
	case CannotCreate:
		return "cannot be created"
	case CannotUpdate:
		return "cannot be updated"
	case CannotDelete:
		return "cannot be deleted"
	case CannotFind:
		return "cannot be found"
	case StillInUse:
		return "still in use"
	case CannotOpen:
		return "cannot be opened"
	case CannotStart:
		return "cannot start"
	case CannotGet:
		return "cannot get information"
	case NotLoggedIn:
		return "unauthorized, please provide token"
	default:
		return "cannot be processed"
	}
}
