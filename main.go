package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"calculadora/entrada"
	"calculadora/historico"
	"calculadora/operacoes"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var listaHistorico []string

	for {
		fmt.Println("======== Calculadora ========")
		time.Sleep(1 * time.Second)
		fmt.Println("1 - Somar")
		fmt.Println("2 - Subtração")
		fmt.Println("3 - Multiplicação")
		fmt.Println("4 - Divisão")
		fmt.Println("5 - Histórico")
		fmt.Println("6 - Sair")

		choice := entrada.LerNumeroInteiro(reader)

		switch choice {

		case 1, 2, 3, 4:
			fmt.Print("Digite o primeiro número: ")
			num1 := entrada.LerNumeroDecimal(reader)

			fmt.Print("Digite o segundo número: ")
			num2 := entrada.LerNumeroDecimal(reader)

			var resultado float64
			var operacao string

			switch choice {
			case 1:
				resultado = operacoes.Somar(num1, num2)
				operacao = "+"
			case 2:
				resultado = operacoes.Subtrair(num1, num2)
				operacao = "-"
			case 3:
				resultado = operacoes.Multiplicar(num1, num2)
				operacao = "*"
			case 4:
				res, err := operacoes.Dividir(num1, num2)
				if err != nil {
					fmt.Println("Erro:", err)
					continue
				}
				resultado = res
				operacao = "/"
			}

			registro := fmt.Sprintf("%.2f %s %.2f = %.2f", num1, operacao, num2, resultado)

			listaHistorico = historico.Adicionar(listaHistorico, registro)

			fmt.Println("Resultado:", registro)

		case 5:
			historico.Listar(listaHistorico)

		case 6:
			fmt.Println("Saindo... 👋")
			time.Sleep(1 * time.Second)
			fmt.Print("Até a próxima! ")
			return

		default:
			fmt.Println("Opção inválida.")
		}

		time.Sleep(2 * time.Second)
	}
}
