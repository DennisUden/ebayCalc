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
	var ek float64
	var frachtMarge float64
	var menge float64
	var vkEbay float64

	stopLoop := false
	// In eine Schleife packen
	for i := 1; stopLoop == false; i++ {

		ekString := newInput("Einkaufspreis")
		ek = toFloat(ekString)
	
		frachtMargeString := newInput("Frachtmarge")
		frachtMarge = toFloat(frachtMargeString)
	
		kat := newInput("Kategorie")
		fmt.Println(kat)
	
		mengeString := newInput("Menge")
		menge = toFloat(mengeString)
	
		vkEbayString := newInput("Ebay Preis")
		vkEbay = toFloat(vkEbayString)
	
		stopLoop = true
	}

	uSt := 0.19
	versand := 5.50
	// Kategorien berücksichtigen
	provision := min(99, (vkEbay * 12/100) + 0.35)
	netEbay := vkEbay / (1 + uSt)
	rawEbay := netEbay - versand - provision

	paypalFix := 0.35
	paypalVar := 0.0299

	vkShopCalc := (rawEbay + paypalFix) / ((1 - paypalVar) / (1 + uSt))
	vkShop := GoLib.Round(vkShopCalc, 2)

//	netShop := vkShopCalc / (1 + uSt)

	einstand := ek + (ek * frachtMarge/100)

	discount := [5]float64{0, 5, 10, 15, 20}

	fmt.Println("Shop Preis:", vkShop)

	fmt.Println("-----------------------------------------")
	fmt.Printf("| %8v | %6v | %5v | %9v |\n", "Discount", "Gewinn", "Marge", "Breakeven")
	fmt.Println("-----------------------------------------")
	for i := 0; i < len(discount); i++ {

		gewinnCalc := rawEbay * (1 - discount[i] / 100) - einstand
		gewinn := GoLib.Round(gewinnCalc, 2)

		margeCalc := gewinnCalc * 100 / rawEbay 
		marge := GoLib.Round(margeCalc, 2)

		breakeven := GoLib.RoundUp((einstand * menge) / gewinnCalc)

		fmt.Printf("| %8v | %6v | %5v | %9v |\n", discount[i], gewinn, marge, breakeven)
	}
}
