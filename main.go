// to add:
// 1) clean up global variables
// 2) move make color variables into a method
package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"math"
)

const uSt float64 = 0.19

const versand float64 = 5.50

const paypalFix float64 = 0.35
const paypalVar float64 = 0.0299

const colorReset string  = "\033[0m"
const colorRed string    = "\033[31m"
const colorGreen string  = "\033[32m"
const colorYellow string = "\033[33m"
const colorBlue string   = "\033[34m"
const colorPurple string = "\033[35m"
const colorCyan string   = "\033[36m"
const colorGray string   = "\033[37m"
const colorWhite string  = "\033[97m"

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

var ekString string 
var ek float64

var frachtMargeString string
var frachtMarge float64

var mengeString string 
var menge float64

var kat string
	
var vkEbayString string
var vkEbay float64

var provision float64
	
var provisionBoote float64
var provisionGarten float64
	
var netEbay float64
var rawEbay float64

var vkShopCalc float64

var einstand float64

var discount [5]float64 = [5]float64{0, 5, 10, 15, 20}


func greeting() {
	fmt.Println("______________.                  _________        .__          ")
	fmt.Println("\\_   _____/\\_ |__ _____  ___.__. \\_   ___ \\_____  |  |   ____  ")
	fmt.Println(" |    __)_  | __ \\\\__  \\<   |  | /    \\  \\/\\__  \\ |  | _/ ___\\ ")
	fmt.Println(" |        \\ | \\_\\ \\/ __ \\\\___  | \\     \\____/ __ \\|  |_\\  \\___ ")
	fmt.Println("/_______  / |___  (____  / ____|  \\______  (____  /____/\\___  >")
	fmt.Println("        \\/      \\/     \\/\\/              \\/     \\/          \\/ ")

	fmt.Println(colorYellow+"Send q to quit"+colorReset)
}

func newInput(question string) string {
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

func calcOutput() {
	provisionBoote = min(990, vkEbay) * 11/100 + max(0, vkEbay - 990) * 2/100
	provisionGarten = min(200, vkEbay) * 12/100 + max(0, vkEbay - 200) * 2/100
	
	switch kat {
		case "b": provision = provisionBoote
		case "g": provision = provisionGarten
	}

	netEbay = vkEbay / (1 + uSt)
	rawEbay = netEbay - versand - provision

	vkShopCalc = (rawEbay + paypalFix) / ((1 - paypalVar) / (1 + uSt))

	einstand = ek + (ek * frachtMarge/100)
}

func writeOutput() {
	fmt.Printf("Shop Preis: %.2f", vkShopCalc)
	
	fmt.Println("-----------------------------------------")
	fmt.Printf("| %8v | %6v | %5v | %9v |\n", "Discount", "Gewinn", "Marge", "Breakeven")
	fmt.Println("-----------------------------------------")
	for i := 0; i < len(discount); i++ {

		gewinnCalc := rawEbay * (1 - discount[i] / 100) - einstand
		gewinn := fmt.Sprintf("%.2f", gewinnCalc)

		margeCalc := gewinnCalc * 100 / rawEbay 
		marge := fmt.Sprintf("%.2f", margeCalc)

		var margeColor string	
		switch {
		case margeCalc >= 50: margeColor = colorGreen+marge+colorReset
		case margeCalc >= 30: margeColor = colorYellow+marge+colorReset
		case margeCalc < 30: margeColor = colorRed+marge+colorReset
		}

		breakeven := math.Ceil((einstand * menge) / gewinnCalc)

		fmt.Printf("| %8v | %6v | %5v | %9v |\n", discount[i], gewinn, margeColor, breakeven)
	}
	fmt.Println("-----------------------------------------")
}

func main() {
	greeting()

	ekString = newInput("Einkaufspreis")
	ek = toFloat(ekString)

	frachtMargeString = newInput("Frachtmarge")
	frachtMarge = toFloat(frachtMargeString)

	mengeString = newInput("Menge")
	menge = toFloat(mengeString)

	fmt.Println(colorYellow+"Send b for boats or g for garden categories"+colorReset)
	kat = newInput("Kategorie")

	vkEbayString = newInput("Ebay Preis")
	vkEbay = toFloat(vkEbayString)

	calcOutput()

	writeOutput()

	for {
		fmt.Println(colorYellow+"Send ek to start with a new product"+colorReset)
		vkEbayString = newInput("Ebay Preis")

		if vkEbayString == "ek" {
			ekString = newInput("Einkaufspreis")
			ek = toFloat(ekString)

			mengeString = newInput("Menge")
			menge = toFloat(mengeString)

			vkEbayString = newInput("Ebay Preis")
			vkEbay = toFloat(vkEbayString)

			calcOutput()

			writeOutput()

			continue
		}
		vkEbay = toFloat(vkEbayString)

		calcOutput()

		writeOutput()
	}
}
