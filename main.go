// to add:
package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"math"
)

func color(s string, c string) string {
	m := map[string]string{
		"reset": "\033[0m",
		"red": "\033[31m",
		"green": "\033[32m",
		"yellow": "\033[33m",
		"blue": "\033[34m",
		"purple": "\033[35m",
		"cyan": "\033[36m",
		"gray": "\033[37m",
		"white": "\033[97m",
	}

	return m[c]+s+m["reset"]
}

func greeting() {
	fmt.Println("______________.                  _________        .__          ")
	fmt.Println("\\_   _____/\\_ |__ _____  ___.__. \\_   ___ \\_____  |  |   ____  ")
	fmt.Println(" |    __)_  | __ \\\\__  \\<   |  | /    \\  \\/\\__  \\ |  | _/ ___\\ ")
	fmt.Println(" |        \\ | \\_\\ \\/ __ \\\\___  | \\     \\____/ __ \\|  |_\\  \\___ ")
	fmt.Println("/_______  / |___  (____  / ____|  \\______  (____  /____/\\___  >")
	fmt.Println("        \\/      \\/     \\/\\/              \\/     \\/          \\/ ")

	fmt.Println(color("Send q to quit", "yellow"))
}

func newInput(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v: ", question)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	answer = strings.ReplaceAll(answer, ",", ".")
	if answer == "q" {
		os.Exit(0)
	}
	return answer
}

func toFloat(a string) float64 {
	aFloat, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println(err)
	}
	return aFloat
}

func calcOutput(kat string, ek, frachtMarge, menge, vkEbay float64, rawEbay, vkShopCalc, einstand *float64) {
	provisionBoote := min(990, vkEbay) * 11/100 + max(0, vkEbay - 990) * 2/100
	provisionGarten := min(200, vkEbay) * 12/100 + max(0, vkEbay - 200) * 2/100
	
	var provision float64
	switch kat {
		case "b": provision = provisionBoote
		case "g": provision = provisionGarten
	}

	const uSt float64 = 0.19
	const versand float64 = 5.50
	const paypalFix float64 = 0.35
	const paypalVar float64 = 0.0299

	netEbay := vkEbay / (1 + uSt)
	*rawEbay = netEbay - versand - provision

	*vkShopCalc = (*rawEbay + paypalFix) / ((1 - paypalVar) / (1 + uSt))

	*einstand = ek + (ek * frachtMarge/100)
}

func writeOutput(menge, rawEbay, vkShopCalc, einstand float64) {
	fmt.Printf("Shop Preis: %.2f\n", vkShopCalc)
	
	fmt.Println("-----------------------------------------")
	fmt.Printf("| %8v | %6v | %5v | %9v |\n", "Discount", "Gewinn", "Marge", "Breakeven")
	fmt.Println("-----------------------------------------")

	var discount [5]float64 = [5]float64{0, 5, 10, 15, 20}
	for i := 0; i < len(discount); i++ {

		gewinnCalc := rawEbay * (1 - discount[i] / 100) - einstand
		gewinn := fmt.Sprintf("%.2f", gewinnCalc)

		margeCalc := gewinnCalc * 100 / rawEbay 
		marge := fmt.Sprintf("%.2f", margeCalc)

		var margeColor string
		switch {
		case margeCalc >= 50: margeColor = color(marge, "green")
		case margeCalc >= 30: margeColor = color(marge, "yellow")
		case margeCalc < 30: margeColor = color(marge, "red")
		}

		breakeven := math.Ceil((einstand * menge) / gewinnCalc)

		fmt.Printf("| %8v | %6v | %5v | %9v |\n", discount[i], gewinn, margeColor, breakeven)
	}
	fmt.Println("-----------------------------------------")
}

func main() {
	greeting()

	ekString := newInput("Einkaufspreis")
	ek := toFloat(ekString)

	frachtMargeString := newInput("Frachtmarge")
	frachtMarge := toFloat(frachtMargeString)

	mengeString := newInput("Menge")
	menge := toFloat(mengeString)

	fmt.Println(color("Send b for boats or g for garden categories", "yellow"))
	kat := newInput("Kategorie")

	vkEbayString := newInput("Ebay Preis")
	vkEbay := toFloat(vkEbayString)

	var rawEbay float64

	var vkShopCalc float64

	var einstand float64

	calcOutput(kat, ek, frachtMarge, menge, vkEbay, &rawEbay, &vkShopCalc, &einstand)

	writeOutput(menge, rawEbay, vkShopCalc, einstand)

	for {
		fmt.Println(color("Send ek to start with a new product", "yellow"))
		vkEbayString = newInput("Ebay Preis")

		if vkEbayString == "ek" {
			ekString = newInput("Einkaufspreis")
			ek = toFloat(ekString)

			mengeString = newInput("Menge")
			menge = toFloat(mengeString)

			vkEbayString = newInput("Ebay Preis")
			vkEbay = toFloat(vkEbayString)

			calcOutput(kat, ek, frachtMarge, menge, vkEbay, &rawEbay, &vkShopCalc, &einstand)

			writeOutput(menge, rawEbay, vkShopCalc, einstand)

			continue
		}
		vkEbay = toFloat(vkEbayString)

		calcOutput(kat, ek, frachtMarge, menge, vkEbay, &rawEbay, &vkShopCalc, &einstand)

		writeOutput(menge, rawEbay, vkShopCalc, einstand)
	}
}
