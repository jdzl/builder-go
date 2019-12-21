package builder

type Executer interface {
	LoadActions(actions []int)
	BuildPart()
}