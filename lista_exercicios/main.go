package main

import "fmt"

type no struct {
	valor   int
	proximo *no
}

type lista struct {
	inicio *no
}

func (l *lista) adicionarInicio(valor int) {
	novo := &no{valor: valor}
	novo.proximo = l.inicio
	l.inicio = novo
}

func (l *lista) adicionarFim(valor int) {
	novo := &no{valor: valor}
	if l.inicio == nil {
		l.inicio = novo
		return
	}
	atual := l.inicio
	for atual.proximo != nil {
		atual = atual.proximo
	}
	atual.proximo = novo
}

func main() {
	minhaLista := lista{}

	// Adicionando os 5 valores misturados
	minhaLista.adicionarInicio(10)
	minhaLista.adicionarFim(20)
	minhaLista.adicionarInicio(5)
	minhaLista.adicionarFim(30)
	minhaLista.adicionarInicio(1)

	// Conferindo o resultado com um for simples
	fmt.Println("Minha lista depois das insercoes:")
	atual := minhaLista.inicio
	for atual != nil {
		fmt.Print(atual.valor, " -> ")
		atual = atual.proximo
	}
	fmt.Println("nil")
}
