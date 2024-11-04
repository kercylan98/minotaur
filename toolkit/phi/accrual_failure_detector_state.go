package phi

import "time"

type accrualFailureDetectorState struct {
	history   accrualFailureDetectorHistory
	timestamp *time.Time
}
