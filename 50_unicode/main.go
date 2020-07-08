package main

import "fmt"

func main() {
	var symbol rune = 'a'
	var autoSymbol = 'a'
	unicodeSymbol := '¤'
	unicodeSymbolByNumber := '\u2318'

	println(symbol, autoSymbol, unicodeSymbol, unicodeSymbolByNumber)

	str1 := "Привет, Мир!"
	fmt.Println("ru: ", str1, len(str1))
	for index, runeValue := range str1 {
		fmt.Printf("%#U at position %d\n", runeValue, index)
	}
	println(str1[1])

	bin := []byte(str1)
	fmt.Println("binary ru: ", bin, len(bin))
	for idx, val := range bin {
		fmt.Printf("raw binary idx: %v, oct: %v, hex: %x\n", idx, val, val)
	}
}
