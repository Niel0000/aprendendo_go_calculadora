package historico

import "fmt"

// Adicionar recebe o histórico atual e um novo registro,
// e retorna o histórico atualizado
func Adicionar(historico []string, registro string) []string {
	historico = append(historico, registro)
	return historico
}

// Listar mostra todos os itens do histórico
func Listar(historico []string) {
	fmt.Println("==== Últimos Cálculos ====")

	if len(historico) == 0 {
		fmt.Println("Nenhum cálculo realizado ainda.")
		return
	}

	for i, item := range historico {
		fmt.Printf("%d) %s\n", i+1, item)
	}
}
