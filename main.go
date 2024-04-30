// to add:
// input loop cleanup
// maybe switch to fmt.scanf
// marge and breakeven sometimes break

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

func getColor() map[string]string {
	color := map[string]string{
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
	return color
}

func color(text string, color string) string {
	return getColor()[color]+text+getColor()["reset"]
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

func toFloat(a string) (float64, error) {
	aFloat, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, err
	} 
	return aFloat, nil
}

func writeOutput(inputs inputs) {

	provisionBoote := min(990, inputs.vkEbay) * 12/100 + max(0, inputs.vkEbay - 990) * 2/100 + 0.35
	provisionGarten := min(200, inputs.vkEbay) * 12/100 + max(0, inputs.vkEbay - 200) * 2/100 + 0.35

	var provision float64

	switch inputs.kat {
	case "b": provision = provisionBoote
	case "g": provision = provisionGarten
	}

	const uSt float64 = 0.19
	const versand float64 = 5.50
	const paypalFix float64 = 0.39
	const paypalVar float64 = 0.0299

	netEbay := inputs.vkEbay / (1 + uSt)
	rawEbay := netEbay - versand - provision

	vkShopCalc := (rawEbay + paypalFix) / (1 / (1 + uSt) - paypalVar)
	vkShopNet := vkShopCalc / (1 + uSt)

	einstand := inputs.ek + (inputs.ek * inputs.fracht/100)

	fmt.Println("-------------------------------------------")
	fmt.Printf("Shop Preis: %.2f\n", vkShopCalc)

	fmt.Println("-------------------------------------------")
	fmt.Printf("| %8v | %6v | %7v | %9v |\n", "Discount", "Gewinn", "Marge", "Breakeven")
	fmt.Println("-------------------------------------------")

	discount := [5]float64{0, 5, 10, 15, 20}

	for i := 0; i < len(discount); i++ {
		vkShopDisc := vkShopNet * (1 - discount[i]/100)
		paypalGes := paypalFix + paypalVar * vkShopDisc

		gewinnCalc := vkShopDisc - paypalGes - einstand
		gewinn := fmt.Sprintf("%.2f", gewinnCalc)

		margeCalc := gewinnCalc * 100 / rawEbay
		var marge string
		if gewinnCalc < 0 {
			marge = ""
		} else {
			marge = fmt.Sprintf("%.2f", margeCalc)
		}

		var margeColor string
		switch {
		case margeCalc >= 50: margeColor = getColor()["green"]
		case margeCalc >= 30: margeColor = getColor()["yellow"]
		case margeCalc < 30: margeColor = getColor()["red"]
		}

		breakevenCalc := math.Ceil((einstand * inputs.menge) / (vkShopDisc - paypalGes))
		var breakeven string
		if breakevenCalc > inputs.menge || breakevenCalc < 0 {
			breakeven = ""
		} else {
			breakeven = fmt.Sprintf("%v", breakevenCalc)
		}

		fmt.Printf("| %8v | %6v | %s%7v%s | %9v |\n", discount[i], gewinn, margeColor, marge, getColor()["reset"],breakeven)
	}
	fmt.Println("-------------------------------------------")
}
// this section looks messy. i hope to improve it in the near future
func main() {
	greeting()

	inputs := inputs{}

	// work in progress: wrapping each input-prompt in an error handler and a loop
	for {
		var err error

		ekString := newInput("Einkaufspreis")
		inputs.ek, err = toFloat(ekString)

		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
		

	// frachtMargeString := newInput("Frachtmarge")
	// inputs.fracht = toFloat(frachtMargeString)

	// mengeString := newInput("Menge")
	// inputs.menge = toFloat(mengeString)

	fmt.Println(color("Send ", "yellow")+"b"+color(" for boats or ", "yellow")+"g"+color(" for garden categories", "yellow"))
	inputs.kat = newInput("Kategorie")

	vkEbayString := newInput("Ebay Preis")
	// inputs.vkEbay = toFloat(vkEbayString)
	
	writeOutput(inputs)

	for {
		fmt.Println(color("Send ", "yellow")+"ek"+color(" to start with a new product", "yellow"))
		// vkEbayString = newInput("Ebay Preis")

		if vkEbayString == "ek" {
			// ekString = newInput("Einkaufspreis")
			// inputs.ek = toFloat(ekString)

			// mengeString = newInput("Menge")
			// inputs.menge = toFloat(mengeString)

			// vkEbayString = newInput("Ebay Preis")
			// inputs.vkEbay = toFloat(vkEbayString)

			writeOutput(inputs)

			continue
		}
		// inputs.vkEbay = toFloat(vkEbayString)

		writeOutput(inputs)
	}
}
