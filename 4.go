package main

import (
	"fmt"
	"math"
)

type Geometria struct {
	B, H, L1, L2, op
}

func (g Geometria) CalculaArea() float64 {
	return (g.B*g.H)
}

func (g Geometria) CalculaPerimetro() float64 {
	return (g.L1*2+g.L2*2)
}

func main() {
	fmt.Println("Qual operação deseja fazer?")
	fmt.Println("Digite 1 para calcular a área, 2 para calcular o perímetro e 3 para sair")
	for _, op := op {
		switch op {
		case 1:
			fmt.Println("Digite a base do retângulo: ")
			fmt.Scan(&B)
			fmt.Println("Digite a altura do retângulo")
			fmt.Scan(&H)
			fmt.Println("A área desse retângulo é: ", CalculaArea())
		
		case 2:
			fmt.Println("Digite o lado maior do retângulo: ")
			fmt.Scan(&L1)
			fmt.Println("Digite o lado menor do retângulo: ")
			fmt.Scan(&L2)
			fmt.Println("O perímetro desse retângulo é: ", CalculaPerimetro())
		case 3:
			fmt.Println("Saindo do programa...")
			return 
		default:
			fmt.Println("Digite um valor válido.")
		}
	}
}

