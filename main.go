package main

import "fmt"

func main() {

	var chosennb1 float32
	var chosennb2 float32
	var userchoice string
	var tryagain = true
	var useranwsser string

	for tryagain {
		fmt.Println("Enter your first number: ")
		fmt.Scanln(&chosennb1)

		fmt.Println("Enter your second number: ")
		fmt.Scanln(&chosennb2)

		fmt.Println("Enter the operation you want: (+, -, *, /")
		fmt.Scanln(&userchoice)

		switch userchoice {
		case "+":
			fmt.Printf("Resultat: %.2f\n", chosennb1+chosennb2)

		case "-":
			fmt.Printf("Result: %.2f\n", chosennb1-chosennb2)

		case "*":
			fmt.Printf("Result: %.2f\n", chosennb1*chosennb2)

		case "/":
			if chosennb2 != 0 {
				fmt.Printf("Result: %.2f\n", chosennb1/chosennb2)
			} else {
				fmt.Printf("Division by 0")
			}
		default:
			fmt.Println("Invalid operation choice")
		}

		fmt.Println("Do you to play again ? y / n")
		fmt.Scanln(&useranwsser)

		if useranwsser == "y" {
			tryagain = true
		} else {
			tryagain = false
			fmt.Println("Thanks for playing !")
		}
	}
}
