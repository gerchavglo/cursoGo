package main

import (
"bufio"
"fmt"
"strings"
"strconv"
"os"
)
func main(){
	scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Calculadora Ver 0.1, implementa fx suma como func")

	scanner.Scan()
	operacion := scanner.Text()

    resultado := sumar(operacion)

	fmt.Println(resultado)



}

func sumar(operacion string) int {

	valores := strings.Split(operacion, "+")
	resultado := 0


	for i := 0; i< len(valores); i++ {
		num,_ := strconv.Atoi(valores [i])
		resultado += num
	}

	return resultado
}
