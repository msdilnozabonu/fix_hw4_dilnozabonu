package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"product/internal/product"
)

func main() {
	var (
		productInfo product.Product
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Product info: ")
	productInfo.Name, _ = reader.ReadString('\n')

	productInfo.Name = strings.TrimSpace(productInfo.Name)

	fmt.Print("Brand: ")
	productInfo.Brand, _ = reader.ReadString('\n')
	productInfo.Brand = strings.TrimSpace(productInfo.Brand)
	fmt.Print("Price: ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSpace(priceStr)
	priceStr = strings.ReplaceAll(priceStr, " ", "")

	if !isValidPriceString(priceStr) {
		fmt.Println("вы вели не правильную сумму")
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("вы вели не правильную сумму")
		return
	}

	fmt.Print("In stock? (0-false,1-true): ")
	stockStr, _ := reader.ReadString('\n')
	stockStr = strings.TrimSpace(stockStr)
	switch stockStr {
	case "1":
		productInfo.InStock = true
	case "0":
		productInfo.InStock = false
	default:
		inStock, err := strconv.ParseBool(stockStr)
		if err != nil {
			productInfo.InStock = false
		} else {
			productInfo.InStock = inStock
		}
	}
	productInfo.Price = int(price * tiinToSum)

	calculatedAmount := product.Calculate(productInfo.Price)

	converted := product.Converter(productInfo, calculatedAmount)
	fmt.Println(converted)
}

func isValidPriceString(s string) bool {
	if s == "" {
		return false
	}
	dotCount := 0
	start := 0
	if s[0] == '-' {
		start = 1
	}
	if start == len(s) {
		return false
	}
	for i := start; i < len(s); i++ {
		c := s[i]
		if c == '.' {
			dotCount++
			if dotCount > 1 {
				return false
			}
		} else if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

const (
	tiinToSum = 100
)
