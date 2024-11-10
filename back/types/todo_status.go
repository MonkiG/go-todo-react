package types

type TodoStatus int

const (
	TODO TodoStatus = iota
	IN_PROGRESS
	DONE
)
