package log

import "strings"

type stackError struct {
	err error
}

type stackErrorTracks []string

func (s stackErrorTracks) String() string {
	return strings.Join(s, "\n")
}
