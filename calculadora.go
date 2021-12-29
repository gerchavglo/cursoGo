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
	scanner.Scan()
	operacion := scanner.Text()

valores := strings.Split(operacion, "+")

operador1,_ := strconv.Atoi(valores[0])
operador2,_ := strconv.Atoi(valores[1])

	fmt.Println(operador1+operador2)

}
