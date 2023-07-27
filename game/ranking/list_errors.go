package ranking

import "errors"

var (
	ErrListNotExistCompetitor = errors.New("ranking list not exist competitor")
	ErrListIndexErr           = errors.New("ranking list index error")
	ErrListNonexistentRanking = errors.New("nonexistent ranking")
)
