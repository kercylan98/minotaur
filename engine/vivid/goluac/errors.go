package goluac

import (
	"errors"
	"fmt"
)

var (
	ErrorNotLoadGoluacComponent = errors.New("not load goluac component")
	ErrorNotIsLuaMessage        = errors.New("not is lua message")
)

func newDoLuaScriptError(name string, err error) error {
	return &DoLuaScriptError{
		Name: name,
		Err:  err,
	}
}

type DoLuaScriptError struct {
	Name string
	Err  error
}

func (e *DoLuaScriptError) Error() string {
	return fmt.Sprintf("do lua script %s error: %s", e.Name, e.Err)
}
