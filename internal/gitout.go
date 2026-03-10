package internal

type GitErr int

const (
	MissingRemoteChanges GitErr = iota
	CommitAfterMerge
	MissingPerms
	InterruptedOperation
	InvalidRemoteCredentials
	UnrelatedHists
	TooBig
	NotARepo
	Unknown
)

type CommitaErr struct {
	Raw  string
	Type GitErr
	Msg  string
}

func (e *CommitaErr) Error() string {
	return e.Msg
}
