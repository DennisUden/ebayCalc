// to add:
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
	
var uSt float64 = 0.19

var versand float64 = 5.50

var provisionBoote float64
var provisionGarten float64
	
var netEbay float64
var rawEbay float64

var paypalFix float64 = 0.35
var paypalVar float64 = 0.0299

var vkShopCalc float64
var vkShop float64

var einstand float64

var discount [5]float64 = [5]float64{0, 5, 10, 15, 20}

var colorReset string  = "\033[0m"
var colorRed string    = "\033[31m"
var colorGreen string  = "\033[32m"
var colorYellow string = "\033[33m"
var colorBlue string   = "\033[34m"
var colorPurple string = "\033[35m"
var colorCyan string   = "\033[36m"
var colorGray string   = "\033[37m"
var colorWhite string  = "\033[97m"

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
	vkShop = GoLib.Round(vkShopCalc, 2)

	einstand = ek + (ek * frachtMarge/100)
}

func writeOutput() {
	fmt.Println("Shop Preis:", vkShop)
	
	fmt.Println("-----------------------------------------")
	fmt.Printf("| %8v | %6v | %5v | %9v |\n", "Discount", "Gewinn", "Marge", "Breakeven")
	fmt.Println("-----------------------------------------")
	for i := 0; i < len(discount); i++ {

		gewinnCalc := rawEbay * (1 - discount[i] / 100) - einstand
		gewinn := GoLib.Round(gewinnCalc, 2)

		margeCalc := gewinnCalc * 100 / rawEbay 
		marge := GoLib.Round(margeCalc, 2)

		margeStr := strconv.FormatFloat(marge, 'f', 2, 64)
		var margeColor string	
		switch {
		case marge >= 50: margeColor = colorGreen+margeStr+colorReset
		case marge >= 30: margeColor = colorYellow+margeStr+colorReset
		case marge < 30: margeColor = colorRed+margeStr+colorReset
		}

		breakeven := GoLib.RoundUp((einstand * menge) / gewinnCalc)

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
