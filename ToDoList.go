package main

import (
"bufio"
"fmt"
"strings"
"strconv"
"os"
)

var keeprunning bool = false

func main(){

   for !keeprunning { 

         menu() 
   }
   fmt.Println (" Program Stoped! ")
   
}


func menu() {
	scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("To-Do List Ver 0.1, Available options: (A)add, (L)list, D(Delete), N(New), Q(Quit) ")

	scanner.Scan()
	operacion := scanner.Text()

    resultado := sumar(operacion)

		if strings.Contains(operacion, "A"){
		   resultado = sumar(operacion)
		}else if strings.Contains(operacion, "L"){
			resultado = sumar(operacion)
		}else if strings.Contains(operacion, "D"){
			resultado = sumar(operacion)
		}else if strings.Contains(operacion, "N"){
			resultado = sumar(operacion)
		}else if strings.Contains(operacion, "Q"){
			keeprunning = true
		}else{
			fmt.Println("Error!, NOT suitable option given")
		}
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
