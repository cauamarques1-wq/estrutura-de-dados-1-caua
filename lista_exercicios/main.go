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

func (l *lista) removerInicio() (int, bool) {
	if l.inicio == nil {
		return 0, false
	}
	removido := l.inicio.valor
	l.inicio = l.inicio.proximo
	return removido, true
}

func (l *lista) removerFim() (int, bool) {
	if l.inicio == nil {
		return 0, false
	}
	if l.inicio.proximo == nil {
		rem := l.inicio.valor
		l.inicio = nil
		return rem, true
	}
	atual := l.inicio
	for atual.proximo.proximo != nil {
		atual = atual.proximo
	}
	rem := atual.proximo.valor
	atual.proximo = nil
	return rem, true
}

func (l *lista) removerPosicao(pos int) (int, bool) {
	if pos < 0 || l.inicio == nil {
		return 0, false
	}
	if pos == 0 {
		return l.removerInicio()
	}
	
	atual := l.inicio
	contador := 0
	for atual != nil && contador < pos-1 {
		atual = atual.proximo
		contador++
	}
	
	if atual == nil || atual.proximo == nil {
		return 0, false
	}
	
	removido := atual.proximo.valor
	atual.proximo = atual.proximo.proximo
	return removido, true
}

func (l *lista) posicao(valorProcurado int) (int, bool) {
	atual := l.inicio
	pos := 0
	for atual != nil {
		if atual.valor == valorProcurado {
			return pos, true
		}
		atual = atual.proximo
		pos++
	}
	return 0, false
}

func main() {
	minhaLista := lista{}

	minhaLista.adicionarFim(10)
	minhaLista.adicionarFim(20)
	minhaLista.adicionarFim(30)
	minhaLista.adicionarFim(40)

	fmt.Println("lista atual:")
	atual := minhaLista.inicio
	for atual != nil {
		fmt.Print(atual.valor, " -> ")
		atual = atual.proximo
	}
	fmt.Println("nil")

	p, ok := minhaLista.posicao(30)
	if ok {
		fmt.Println("achou o 30 na posicao:", p)
	}

	p2, ok2 := minhaLista.posicao(99)
	if !ok2 {
		fmt.Println("buscou o 99 e retornou false certinho")
	}
}