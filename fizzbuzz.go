package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// fb takes in an int value and returns string value of fizz/buzz/fizzbuzz or an error
func fb(a int) string {
	if a%3 == 0 && a%5 == 0 {
		return "fizzbuzz"
	} else if a%5 == 0 {
		return "buzz"
	} else if a%3 == 0 {
		return "fizz"
	}
	return "You dont know how to play"

}
func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter number to fizz buzz: ")
	reader.Scan()
	int_text, err := strconv.Atoi(reader.Text()) //convert string to int
	if err != nil {
		fmt.Println("Enter a number :)") //If still string return err
	}
	fmt.Println("Fizzbuzz tests result is - " + fb(int_text))
}
