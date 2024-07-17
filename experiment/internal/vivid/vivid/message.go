package vivid

import "github.com/kercylan98/minotaur/experiment/internal/vivid/prc"

type Message = prc.Message

const (
	onSuspendMailbox onSuspendMailboxMessage = 0
	onResumeMailbox  onResumeMailboxMessage  = 0
)

type (
	onSuspendMailboxMessage int8
	onResumeMailboxMessage  int8
)
