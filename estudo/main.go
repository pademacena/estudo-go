package main

import "fmt"

// variaves escopo de paocote
// var idade int = 27
// se atribuir o valor nao preciso necessariamente colocar o tipo
var idade = 27

func main() {
	fmt.Println("Hello, world")

	//FUNCOES
	digaOi()
	fmt.Println(somar(1, 2))
	a, b := swap(10, 20)
	fmt.Println(a, b)
	res, rem := dividir(10, 3)
	fmt.Println(res, rem)
	f := novaSoma(2)
	x := f(1)
	fmt.Println(x)
	fmt.Println(somarIndefinido(1, 2, 3, 4, 5, 6, 7))

	//VARIAVEIS

	//variaveis declaradas sem valor atribuido recebem 0 ou ""
	// var nome, sobrenome string = "Paulo", "Pereira"
	// var (
	// 	nome      string = "Paulo"
	// 	sobrenome string = "Pereira"
	// )
	// tanto faz dos modos de cima ou do modo abaixo

	nome := "Paulo"
	sobrenome := "Pereira"
	// o modo acima nao é valido em escopo de pacote, somente de funcao
	fmt.Println(nome, sobrenome, idade)

	// na declaracao curta o tipo sempre ser inferido, nao posso atribuir
	foo, bar := "foo", 50
	fmt.Println(foo, bar)

	// y := 10
	// para tipos especificos o modelo acima nao funciona, preciso fazer a declaracao padrao
	var y int8 = 10
	fmt.Println(y)
}

func digaOi() {
	fmt.Println("Oi")
}

func somar(a int, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

// := -> declarando uma variavel
// = -> atalizando ou atribuindo um valor
func dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	// retornando variaveis ja declaradas na fuuncao (nao é indicado) : naked return
	// return
	return res, rem
}

// funcao clousre
func novaSoma(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func somarIndefinido(nums ...int) int {
	var out int
	for _, n := range nums {
		out += n
	}
	return out
}
