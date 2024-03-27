package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
//	"github.com/DennisUden/GoLib"
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
	ek := newInput("Einkaufspreis")
	ekFloat := toFloat(ek)
	fmt.Println(ekFloat*2)

	marge := newInput("Frachtmarge")
	margeFloat := toFloat(marge)
	fmt.Println(margeFloat*2)

	kat := newInput("Kategorie")
	fmt.Println(kat)

	vk := newInput("Verkaufspreis")
	vkFloat := toFloat(vk)
	fmt.Println(vkFloat*2)

}
