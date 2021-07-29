package main
import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func leiaLinha() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}
func main(){
	var texto string
	fmt.Println("Insira um texto:")
	texto = leiaLinha()

	letras := make(map[string]int)

	for _, l := range texto {
		letras[string(l)] += 1
	}

	letrasOrdem := make([]string, 0, len(letras))
	for i := range letras {
		letrasOrdem = append(letrasOrdem, i)
	}
	sort.Strings(letrasOrdem)

	for _, c := range letrasOrdem {
		fmt.Printf("%v -> %v ", c, letras[c])
	}
}