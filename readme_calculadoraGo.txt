# 🧮 Calculadora em Go

Uma calculadora de terminal desenvolvida em Go, com suporte a operações básicas, tratamento de erros e histórico de cálculos.

---

## 🚀 Funcionalidades

- ➕ Adição
- ➖ Subtração
- ✖️ Multiplicação
- ➗ Divisão (com tratamento de divisão por zero)
- 📜 Histórico de cálculos
- 📥 Entrada de dados com validação
- 🔁 Loop interativo

---

## 🧠 Conceitos aplicados

- Estrutura de pacotes em Go
- Separação de responsabilidades
- Tratamento de erros (`error`)
- Manipulação de entrada com `bufio`
- Slice para armazenamento de dados
- Modularização do código

---

## 📁 Estrutura do projeto

calculadora/
│
├── entrada/
│ └── entrada.go
│
├── operacoes/
│ └── operacoes.go
│
├── historico/
│ └── historico.go
│
├── main.go
└── go.mod
---

## ▶️ Como executar

No termianl digite: 
go run main.go

Exemplo de funcionamento:

======== Calculadora ========
1 - Soma
2 - Subtração
3 - Multiplicação
4 - Divisão
5 - Histórico
6 - Sair

Escolha: 1
Digite o primeiro número: 5
Digite o segundo número: 2

Resultado: 25.00 + 12.00 = 37.00


Autor: Daniel da Cruz Mendes.

Obs: Este é o meu primeiro código, fiz o básico e depois fui obtendo mais funções a inserir no código.


