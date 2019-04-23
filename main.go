package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// func convertToHEX(ipv4 string) string {

// }

func convertToBIN(ipv4 string) string {
	parts := strings.Split(ipv4, ".")

	ipv4BIN := []string{}

	for _, part := range parts {
		temp, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}

		str := ""
		for temp != 0 {
			bin := temp % 2
			temp = temp / 2
			str = strconv.Itoa(bin) + str
		}

		ipv4BIN = append(ipv4BIN, str)
	}

	return strings.Join(ipv4BIN, ".")
}

func main() {
	filePath := os.Args[1]
	rowNumber := os.Args[2]

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(file), "\n")

	for _, row := range rows {
		rowArr := strings.Split(row, " ")

		number := rowArr[0]
		ipv4 := rowArr[1]

		if number == rowNumber {
			fmt.Println("DEC: ", ipv4)
			fmt.Println("HEX: ", ipv4)
			fmt.Println("BIN: ", convertToBIN(ipv4))

			return
		}
	}
}
