package builder

import "fmt"

type WorkFlow struct {
	Actions []int
}

func (w *WorkFlow) Init() {
	fmt.Println("Configuracion inicial")
}

func (w *WorkFlow) Process() {
	fmt.Println("proceso N+1")
}

func (w *WorkFlow) Finalize() {
	fmt.Println("finalizando proceso")
}

func (w *WorkFlow) LoadActions(actions []int) {
	w.Actions = actions
}

func (w *WorkFlow) BuildPart() {
	for _, v := range w.Actions {
		switch v {
		case 1:
			w.Init()
		case 2:
			w.Init()
		case 3:
			w.Finalize()
		default:
			fmt.Println("evento no registrado")
		}
	}
}