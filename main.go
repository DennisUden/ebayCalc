// to add:
// input loop cleanup
// maybe switch to fmt.scanf

package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"math"
)

type inputs struct {
	ek float64
	fracht float64
	menge float64
	kat string
	vkEbay float64
	}

func color(text string, color string) string {
	colorMap := map[string]string{
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

	return colorMap[color]+text+colorMap["reset"]
}

func greeting() {
	s := `

	 _____ _                    ____      _      
	| ____| |__   __ _ _   _   / ___|__ _| | ___ 
	|  _| | '_ \ / _' | | | | | |   / _' | |/ __|
	| |___| |_) | (_| | |_| | | |__| (_| | | (__ 
	|_____|_.__/ \__,_|\__, |  \____\__,_|_|\___|
			   |___/                     

	`
	fmt.Println(s)
}

func newInput(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v: ", color(question, "purple"))
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

func writeOutput(inputs inputs) {

	provisionBoote := min(990, inputs.vkEbay) * 11/100 + max(0, inputs.vkEbay - 990) * 2/100
	provisionGarten := min(200, inputs.vkEbay) * 12/100 + max(0, inputs.vkEbay - 200) * 2/100

	var provision float64

	switch inputs.kat {
	case "b": provision = provisionBoote
	case "g": provision = provisionGarten
	}

	const uSt float64 = 0.19
	const versand float64 = 5.50
	const paypalFix float64 = 0.35
	const paypalVar float64 = 0.0299

	netEbay := inputs.vkEbay / (1 + uSt)
	rawEbay := netEbay - versand - provision

	if rawEbay <= 0 {
		fmt.Println("Error, rawEbay negative. Set higher price")
		return
	}

	vkShopCalc := (rawEbay + paypalFix) / ((1 - paypalVar) / (1 + uSt))

	einstand := inputs.ek + (inputs.ek * inputs.fracht/100)

	fmt.Println("-----------------------------------------")
	fmt.Printf("Shop Preis: %.2f\n", vkShopCalc)

	fmt.Println("-----------------------------------------")
	fmt.Printf("| %8v | %6v | %5v | %9v |\n", "Discount", "Gewinn", "Marge", "Breakeven")
	fmt.Println("-----------------------------------------")

	discount := [5]float64{0, 5, 10, 15, 20}

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

		breakevenCalc := math.Ceil((einstand * inputs.menge) / gewinnCalc)
		var breakeven string
		if breakevenCalc < 0 {
			breakeven = "never"
		} else {
			breakeven = fmt.Sprintf("%v", breakevenCalc)
		}

		fmt.Printf("| %8v | %6v | %5v | %9v |\n", discount[i], gewinn, margeColor, breakeven)
	}
	fmt.Println("-----------------------------------------")
}
// this section looks messy. i hope to improve it in the near future
func main() {
	greeting()

	inputs := inputs{}

	ekString := newInput("Einkaufspreis")
	inputs.ek = toFloat(ekString)
		

	frachtMargeString := newInput("Frachtmarge")
	inputs.fracht = toFloat(frachtMargeString)

	mengeString := newInput("Menge")
	inputs.menge = toFloat(mengeString)

	fmt.Println(color("Send ", "yellow")+"b"+color(" for boats or ", "yellow")+"g"+color(" for garden categories", "yellow"))
	inputs.kat = newInput("Kategorie")

	vkEbayString := newInput("Ebay Preis")
	inputs.vkEbay = toFloat(vkEbayString)
	
	writeOutput(inputs)

	for {
		fmt.Println(color("Send ", "yellow")+"ek"+color(" to start with a new product", "yellow"))
		vkEbayString = newInput("Ebay Preis")

		if vkEbayString == "ek" {
			ekString = newInput("Einkaufspreis")
			inputs.ek = toFloat(ekString)

			mengeString = newInput("Menge")
			inputs.menge = toFloat(mengeString)

			vkEbayString = newInput("Ebay Preis")
			inputs.vkEbay = toFloat(vkEbayString)

			writeOutput(inputs)

			continue
		}
		inputs.vkEbay = toFloat(vkEbayString)

		writeOutput(inputs)
	}
}
