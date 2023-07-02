package job

type Job interface {
	GetJId() int64
	DoWait() (Result, error)
	Cancel() error
	GetStatus() (Status, error)
	Wait() (Status, error)
}

type EndState int

const (
	NotEnd EndState = iota
	Locally
	Globally
)

type Status int

const (
	Initializing Status = iota
	Created
	Scheduled
	Running
	Failing
	Failed
	Cancelling
	Cancelled
	Finished
	Restarting
	Suspended
	Reconciling
)

func (s Status) GetEndState() EndState {
	switch s {
	case Initializing, Created, Scheduled, Running, Cancelling, Failing, Restarting, Reconciling:
		return NotEnd
	case Failed, Cancelled, Finished:
		return Globally
	case Suspended:
		return Locally
	}
	return NotEnd
}

func (s Status) IsGloballyEndState() bool {
	return s.GetEndState() == Globally
}

func (s Status) IsEndState() bool {
	return s.GetEndState() != NotEnd
}

type Result struct {
	Status Status
	Error  error
}
