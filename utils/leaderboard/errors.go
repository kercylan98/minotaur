package leaderboard

import "errors"

var (
	ErrNotExistCompetitor = errors.New("leaderboard not exist competitor")
	ErrIndexErr           = errors.New("leaderboard index error")
	ErrNonexistentRanking = errors.New("nonexistent ranking")
)
