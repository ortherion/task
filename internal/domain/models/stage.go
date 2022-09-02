package models

type Stage int

const (
	Undefined Stage = iota
	Accept
	Reject
	InProcess
)

func (s Stage) String() string {
	switch s {
	case Undefined:
		return "undefined"
	case Accept:
		return "accept"
	case Reject:
		return "reject"
	case InProcess:
		return "in process"
	default:
		return "unknown type"
	}
}
