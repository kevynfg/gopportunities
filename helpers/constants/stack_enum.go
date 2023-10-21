package constants

type StackEnum uint

const (
	FRONTEND StackEnum = iota + 1
	BACKEND
	DEVOPS
	QA
)

func (s StackEnum) String() string {
	return [...]string{"Frontend", "Backend", "DevOps", "QA"}[s-1]
}