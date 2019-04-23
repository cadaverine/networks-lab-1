package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func convertToHEX(ipv4 string) string {
	ipv4HEX := []string{}
	segments := strings.Split(ipv4, ".")

	for _, segment := range segments {
		temp, err := strconv.Atoi(segment)
		if err != nil {
			panic(err)
		}

		var str, hex string

		for temp != 0 {
			switch hex = strconv.Itoa(temp % 16); hex {
			case "10":
				hex = "A"
			case "11":
				hex = "B"
			case "12":
				hex = "C"
			case "13":
				hex = "D"
			case "14":
				hex = "E"
			case "15":
				hex = "F"
			}

			str = hex + str
			temp = temp / 16
		}

		ipv4HEX = append(ipv4HEX, str)
	}

	return strings.Join(ipv4HEX, ".")
}

func convertToBIN(ipv4 string) string {
	ipv4BIN := []string{}
	segments := strings.Split(ipv4, ".")

	for _, segment := range segments {
		temp, err := strconv.Atoi(segment)
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
			fmt.Println("HEX: ", convertToHEX(ipv4))
			fmt.Println("BIN: ", convertToBIN(ipv4))

			return
		}
	}
}
