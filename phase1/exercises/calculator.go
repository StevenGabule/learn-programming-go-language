package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("==================CALCULATOR==================\n")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter expression (e.g., 5 + 3) or 'quit': ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		result, err := calculate(input)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		fmt.Printf("Result: %.2f\n", result)
	}
}

func calculate(expr string) (float64, error) {
	parts := strings.Fields(expr)

	if len(parts) != 3 {
		return 0, fmt.Errorf("Expected format: num op num")
	}

	a, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, err
	}

	b, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, err
	}

	switch parts[1] {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operator: %f", parts[1])
	}
}
