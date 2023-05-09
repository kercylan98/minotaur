package builtin

import "errors"

var (
	ErrRankingListNotExistCompetitor = errors.New("ranking list not exist competitor")
	ErrRankingListIndexErr           = errors.New("ranking list index error")
	ErrRankingListNonexistentRanking = errors.New("nonexistent ranking")
)
