package entrada

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

/*Esta função lê um número decimal do usuário. e só segue quando recolhe uma entrada válida*/
func LerNumeroDecimal(reader *bufio.Reader) float64 {
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ReplaceAll(input, ",", ".")

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número válido.")
			continue
		}
		return (num)
	}
}

/*Esta função lê um número inteiro do usuário. e só segue quando recolhe uma entrada válida*/
func LerNumeroInteiro(reader *bufio.Reader) int {
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número válido.")
			continue
		}
		return int(num)
	}
}
