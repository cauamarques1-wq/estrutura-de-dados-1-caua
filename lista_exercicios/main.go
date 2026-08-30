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

func (l *lista) valorNaPosicao(posicaoProcurada int) (int, bool) {
	if posicaoProcurada < 0 {
		return 0, false
	}
	atual := l.inicio
	pos := 0
	for atual != nil {
		if pos == posicaoProcurada {
			return atual.valor, true
		}
		atual = atual.proximo
		pos++
	}
	return 0, false
}

func (l *lista) tamanho() int {
	atual := l.inicio
	count := 0
	for atual != nil {
		count++
		atual = atual.proximo
	}
	return count
}

func (l *lista) imprimir() {
	atual := l.inicio
	for atual != nil {
		fmt.Print(atual.valor, " -> ")
		atual = atual.proximo
	}
	fmt.Println("nil")
}

func main() {
	minhaLista := lista{}
	var op int

	for {
		fmt.Println("\n--- MENU DA LISTA ---")
		fmt.Println("1 - Adicionar inicio")
		fmt.Println("2 - Adicionar fim")
		fmt.Println("3 - Adicionar posicao")
		fmt.Println("4 - Remover inicio")
		fmt.Println("5 - Remover fim")
		fmt.Println("6 - Remover posicao")
		fmt.Println("7 - Buscar por valor")
		fmt.Println("8 - Buscar por posicao")
		fmt.Println("9 - Tamanho")
		fmt.Println("10 - Imprimir lista")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha: ")
		fmt.Scan(&op)

		if op == 0 {
			break
		}

		switch op {
		case 1:
			var v int
			fmt.Print("Valor: ")
			fmt.Scan(&v)
			minhaLista.adicionarInicio(v)
		case 2:
			var v int
			fmt.Print("Valor: ")
			fmt.Scan(&v)
			minhaLista.adicionarFim(v)
		case 3:
			var v, p int
			fmt.Print("Valor e posicao: ")
			fmt.Scan(&v, &p)
			ok := minhaLista.adicionarPosicao(v, p)
			if !ok {
				fmt.Println("erro ao add, posicao invalida")
			}
		case 4:
			v, ok := minhaLista.removerInicio()
			if ok {
				fmt.Println("removeu", v)
			} else {
				fmt.Println("lista vazia")
			}
		case 5:
			v, ok := minhaLista.removerFim()
			if ok {
				fmt.Println("removeu o", v)
			} else {
				fmt.Println("deu ruim, lista vazia")
			}
		case 6:
			var p int
			fmt.Print("Qual posicao remover: ")
			fmt.Scan(&p)
			v, ok := minhaLista.removerPosicao(p)
			if ok {
				fmt.Println("tirou o", v)
			} else {
				fmt.Println("posicao n existe")
			}
		case 7:
			var v int
			fmt.Print("Procurar qual numero: ")
			fmt.Scan(&v)
			p, ok := minhaLista.posicao(v)
			if ok {
				fmt.Println("ta na posicao", p)
			} else {
				fmt.Println("nao achei")
			}
		case 8:
			var p int
			fmt.Print("Ver qual posicao: ")
			fmt.Scan(&p)
			v, ok := minhaLista.valorNaPosicao(p)
			if ok {
				fmt.Println("valor eh", v)
			} else {
				fmt.Println("posicao invalida")
			}
		case 9:
			fmt.Println("tamanho da lista:", minhaLista.tamanho())
		case 10:
			minhaLista.imprimir()
		default:
			fmt.Println("opcao errada vei")
		}
	}
}