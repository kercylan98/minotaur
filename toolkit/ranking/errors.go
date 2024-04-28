package ranking

import "errors"

var (
	ErrNotExistCompetitor = errors.New("ranking not exist competitor")
	ErrIndexErr           = errors.New("ranking index error")
	ErrNonexistentRanking = errors.New("nonexistent ranking")
)
