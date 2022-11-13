package main

import "fmt"

type Data struct {
	Name string
	Age  int
}

func main() {
	var a, b float64
	fmt.Scanf("%f + %f", &a, &b)
	fmt.Printf("%.2f + %.2f = %.2f", a, b, a+b)

}
