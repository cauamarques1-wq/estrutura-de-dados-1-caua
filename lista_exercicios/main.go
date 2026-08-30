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

func (l *lista) adicionarPosicao(valor, pos int) bool {
	if pos < 0 {
		return false
	}
	if pos == 0 {
		l.adicionarInicio(valor)
		return true
	}
	
	novo := &no{valor: valor}
	atual := l.inicio
	contador := 0
	
	for atual != nil && contador < pos-1 {
		atual = atual.proximo
		contador++
	}
	
	if atual == nil {
		return false
	}
	
	novo.proximo = atual.proximo
	atual.proximo = novo
	return true
}

func main() {
	minhaLista := lista{}

	minhaLista.adicionarInicio(10)
	minhaLista.adicionarFim(20)
	minhaLista.adicionarInicio(5)
	minhaLista.adicionarFim(30)
	minhaLista.adicionarInicio(1)

	fmt.Println("lista antes:")
	atual := minhaLista.inicio
	for atual != nil {
		fmt.Print(atual.valor, " -> ")
		atual = atual.proximo
	}
	fmt.Println("nil")

	ok := minhaLista.adicionarPosicao(15, 3)
	if ok {
		fmt.Println("conseguiu colocar o 15 na posicao 3")
	}

	ok2 := minhaLista.adicionarPosicao(99, 20)
	if !ok2 {
		fmt.Println("erro ao add, posicao invalida demais")
	}

	fmt.Println("lista depois:")
	atual = minhaLista.inicio
	for atual != nil {
		fmt.Print(atual.valor, " -> ")
		atual = atual.proximo
	}
	fmt.Println("nil")
}