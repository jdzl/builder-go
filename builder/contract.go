package builder

import "fmt"

type Contract struct {
	Actions []int
}

func (v *Contract) GetInfo() {
	fmt.Println("Obtener info contrato")
}

func (v *Contract) NormaX() {
	fmt.Println("Aplicando norma X al contrato")
}

func (v *Contract) NormaY() {
	fmt.Println("Aplicando norma Y al contrato")
}

func (v *Contract) NormaZ() {
	fmt.Println("Aplicando norma Z al contrato")
}

func (v *Contract) Testing() {
	fmt.Println("verificando normas")
}
func (v *Contract) Finalize() {
	fmt.Println("Se ha finalizado el contrato")
}

func (v *Contract) LoadActions(actions []int) {
	v.Actions = actions
}

func (v *Contract) BuildPart() {
	for _, a := range v.Actions {
		switch a {
		case 1:
			v.GetInfo()
		case 2:
			v.NormaX()
		case 3:
			v.NormaY()
		case 4:
			v.NormaZ()
		case 5:
			v.Testing()
		case 6:
			v.Finalize()
		default:
			fmt.Println("evento no registrado")
		}
	}
}