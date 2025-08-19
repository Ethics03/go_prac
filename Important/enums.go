package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)
// there isnt any default enum type in golang
// so we make a custom type
// then use iota -> basically increments starting from 0
// so each state has a value like ->
// Pending -> 0
// Completed -> 1
// Rejected -> 2
type CheckStatus int

const (
	Pending CheckStatus = iota
	Completed
	Rejected
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func enums() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	ns2 := transition(ns)
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:

		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
