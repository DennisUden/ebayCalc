package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"github.com/DennisUden/GoLib"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func newInput(question string) string {
	fmt.Printf("%v: ", question)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	return answer
}

func toFloat(a string) float64 {
	aFloat, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println(err)
	}
	return aFloat
}

func main() {
	// In eine Schleife packen
	ekString := newInput("Einkaufspreis")
	ek := toFloat(ekString)

	frachtMargeString := newInput("Frachtmarge")
	frachtMarge := toFloat(frachtMargeString)

	kat := newInput("Kategorie")
	fmt.Println(kat)

	mengeString := newInput("Menge")
	menge := toFloat(mengeString)

	vkEbayString := newInput("Ebay Preis")
	vkEbay := toFloat(vkEbayString)

	uSt := 0.19
	versand := 5.50
	// Kategorien berücksichtigen
	provision := min(99, (vkEbay * 12/100) + 0.35)
//	netEbay := vkEbay / (1 + uSt)
//	rawEbay := netEbay - versand - provision
//	fmt.Println(rawEbay)

	vkShopCalc := vkEbay - provision - versand
	vkShop := GoLib.Round(vkShopCalc, 2)

	netShop := vkShopCalc / (1 + uSt)

	einstand := ek + (ek * frachtMarge/100)

	discount := [5]float64{0, 5, 10, 15, 20}

	gewinnCalc := netShop - einstand
	gewinn := GoLib.Round(gewinnCalc, 2)

	margeCalc := gewinnCalc * 100 / netShop
	marge := GoLib.Round(margeCalc, 2)

	// Muss aufrunden
	breakeven := GoLib.RoundUp((einstand * menge) / gewinnCalc)

	fmt.Println("Shop Preis:", vkShop)

	fmt.Println("-----------------------------------------")
	fmt.Printf("| %8v | %6v | %5v | %9v |\n", "Discount", "Gewinn", "Marge", "Breakeven")
	fmt.Println("-----------------------------------------")
	fmt.Printf("| %8v | %6v | %5v | %9v |\n", discount[0], gewinn, marge, breakeven)
}
