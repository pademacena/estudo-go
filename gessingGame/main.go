package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Advinhacao")
	fmt.Println(
		"Um numero aleatorio sera sorteado. Tente acertar. O numero é um inteiro entre 0 e 100",
	)

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	x := rand.Int64N(101)

	for i := range chutes {
		fmt.Println("Qual o seu chute")
		scanner.Scan()
		chute := scanner.Text()
		// TrimSpace -> remove os espaços da frente e de tras
		chute = strings.TrimSpace(chute)
		// abaixo converte em inteiro
		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O seu chute tem que ser um inteiro")
			return
		}
		switch {
		case chuteInt < x:
			fmt.Println("voce errou, o numero sorteado é maior que ", chute)
		case chuteInt > x:
			fmt.Println("voce errou, o numero sorteado é menor que ", chute)
		case chuteInt == x:
			fmt.Printf(
				"Parabens, o numero era %d \n"+
					" voce acertou em %d tentativas .\n"+
					" essas foram as suas tentativas: %v\n",
				x, i+1, chutes[:i], // chutes[:i] -> imprime ate o indice atual do array
			)
			return
		}
		chutes[i] = chuteInt
	}
	fmt.Printf(
		"Infelizmente, voce nao acertou o numero, o valor era %d \n"+
			" voce teve 10 tentativas .\n"+
			" essas foram as suas tentativas: %v\n",
		x, chutes,
	)
}
