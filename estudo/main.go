package main

import (
	"fmt"
	"strconv"
)

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

	// bool
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr ( o ultimo usado somente em conceito unsafe)
	// byte -> igual a uint8
	// rune -> igual a int32
	// float32 float64
	// complex64 complex128
	// string

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

	var bt uint8 = 10
	takeByte(bt)

	convertIntFloat()

	// CONSTANTE
	// constante sao imutaveis
	// nem todo tipo pode ser contante
	// tipos caracter, string , booleanos e todos os valores numericos podem ser, o resto nao
	// nao pode mos usar shortContext para constante
	// constante eu posso omitir o meu tipo
	// forma de driblar o sistema de tipo do GOLANG
	// nao posso passar um float constante para uma que converte o valor em inteiro

	const testeConstante = 48
	fmt.Println(testeConstante)
	takeInt32(testeConstante)
	takeInt64(testeConstante)

	// constante literal
	takeInt32(11)
	takeInt64(12)
	takeString("foo") // string literal ou anonimo string

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

func takeByte(b byte) {
	fmt.Println(b)
}

func convertIntFloat() {
	var x int = 65
	f := float64(x)
	fmt.Println(f)
	s := string(x) //Transforma o valor em uma runa, nao é a forma certa de conversão
	fmt.Println(s)
	ss := strconv.FormatInt(int64(x), 10) // forma certa de converter um int em uma string
	fmt.Println(ss)
}

func takeInt32(x int32) {
	fmt.Println(x)
}

func takeInt64(x int64) {
	fmt.Println(x)
}

func takeString(x string) {
	fmt.Println(x)
}
