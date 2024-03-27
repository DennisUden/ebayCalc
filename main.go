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

func newInput(q string) string {
	fmt.Printf("%v: ", q)
	a, _ := reader.ReadString('\n')
	return a
}

func main() {
	ek := newInput("Einkaufspreis")
	
	ekNoSpace := strings.TrimSpace(ek)

	ekFloat, err := strconv.ParseFloat(ekNoSpace , 64)
	if err != nil {
		fmt.Println(err)
	}

	x := 2.0

	fmt.Println(x*ekFloat)

}
