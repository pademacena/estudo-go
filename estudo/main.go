package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
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

	estudoArray()

	aulaLoops()
	explicacao()
	ifConditions()
	switchStatements()
	aulaDeDefer()

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

func estudoArray() {
	// o tempo tem que ser constante
	// nao se pode colocar um tamanho a mais no array( rola erro de compilaçao)

	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	// definindo valor em um indice especifico
	arr2 := [10]int{5: 400, 7: 300}
	fmt.Println(arr2)

	arr3 := [10]string{4: "foo"}
	fmt.Println(arr3)

	// posso atribuir um valor de uma constante para um array, porem nao posso atribuir o valor de uma variavel comum
	// temos arrays dinamicos, porem eles sao Slices e nao sao teratados como um array normal
	const x = 10
	arr4 := [x]string{2: "ola"}
	fmt.Println(arr4)
}

// atualizar o go : baixar o instalador e sobreescrever (windows e mac)
// linux: remover o antigo e instalar o novo (oooh bosta)
func aulaLoops() {
	// 3 statements
	// padrao
	var res int
	for i := 0; i < 10; i++ {
		res++
	}
	fmt.Println(res)

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for range arr {
		fmt.Println("dentro")
	}

	// retorna o indice
	for i := range arr {
		fmt.Println(i)
	}

	// indice e elemento
	for i, elem := range arr {
		fmt.Println(i, elem)
	}

	// _, -> blank identifier: no caso, retorna somente o elemento
	for _, elem := range arr {
		fmt.Println(elem)
	}

	for range 10 {
		fmt.Println("dentro")
	}

	// so posso declarar uma variavel nesse caso
	for i := range 10 {
		fmt.Println(i)
	}
}

func explicacao() {
	const n = 10
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < n; i++ {
		// i := i -> shadow
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}

func ifConditions() {
	if 1 < 2 {
		fmt.Println("sim")
	}

	// declarando valor variavel durante o if
	// variavel so existe dentro do escopo da if condition
	// short statement
	if x := math.Sqrt(16); x < 5 {
		fmt.Println("caiu no if")
	} else if x > 4 {
		fmt.Println("caiu no else")
	}

	meuBool := true
	if meuBool {
		fmt.Println("meuBool true")
	}

	if err := doError(); err != nil {
		fmt.Println("deu erro")
	}
}

func doError() error {
	return errors.New("error")
}

func switchStatements() {
	do(1)
	do(2)
	do(3)
	do2(1)
	do3(1)
	fmt.Println(isWeekend(time.Now()))
	do4()
	do5(time.Now())
	do6("")
}

func do(x int) {
	// adicionar o Break é redundante
	switch x {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println("outra coisa")
	}
}

func do2(x int) {
	// adicionar o Break é redundante
	switch x {
	case 1:
		fmt.Println(1)
		fallthrough // vai pro proximo caso
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println("outra coisa")
	}
}

func do3(x int) {
	switch {
	case x == 1:
		fmt.Println(1)
	case "abc" == "foo":
		fmt.Println(2)
	default:
		fmt.Println("outra coisa")
	}
}

func isWeekend(x time.Time) bool {
	switch {
	case x.Weekday() > 0 && x.Weekday() < 6:
		return false
	default:
		return true
	}
}

func do4() {
	switch x := math.Sqrt(4); x {
	case 2:
		fmt.Println("resultado é 2")
	default:
		fmt.Println("algo de errado")
	}
}

func do5(x time.Time) bool {
	switch x.Weekday() {
	case time.Sunday, time.Saturday:
		return true
	default:
		return false
	}
}

func do6(x any) {
	switch t := x.(type) {
	case string:
		takeString2(t)
	case int:
	case nil:
	}
}

func takeString2(s string) {
	fmt.Println(s)
}

func aulaDeDefer() {
	// defer statement adia uma chamada de funcao ate a funcao principal retornar
	doDefer()

	x := doDefer2()
	fmt.Println(x)

	doDefer3()
	doDefer4()
	doDefer5()
}

func doDefer() {
	defer fmt.Println(" world")
	fmt.Println("Hello")
}

func doDefer2() int {
	defer fmt.Println(" world")
	fmt.Println("Hello")
	return 10
}

func doDefer3() {
	defer fmt.Println(3)
	defer fmt.Println(2)
	fmt.Println(1)
}

func doDefer4() {
	x := 10
	defer func(y int) {
		fmt.Println(y)
	}(x)

	x = 50
	fmt.Println(x)
}

func doDefer5() {
	x := 10
	defer func() {
		fmt.Println(x)
	}()

	x = 50
	fmt.Println(x)
}

// func doDefer6() {
// 	file, err := os.Open("foo.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	// roda mesmo depois do panic
// 	defer file.Close()

// 	if err := x(); err != nil {
// 		file.Close()
// 		panic(err)
// 	}
// }
