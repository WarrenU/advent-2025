package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getMaximalBatteryValue(chargeStr string, batterySize int) int {
	if len(chargeStr) < batterySize {
		return 0
	}

	x := int(chargeStr[0] - '0')
	y := int(chargeStr[1] - '0')

	total := len(chargeStr) - 1

	for i := range total {
		b := int(chargeStr[i+1] - '0')

		if b > x && i+1 != total {
			x = b
			y = int(chargeStr[i+2] - '0')
		} else if b > y {
			y = b
		}
	}

	s := fmt.Sprintf("%d%d", x, y)
	charge, _ := strconv.Atoi(s)
	return charge
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	chargesRunningSum := 0
	for scanner.Scan() {
		chargesRunningSum += getMaximalBatteryValue(scanner.Text(), 2)
	}
	fmt.Printf("Sum of all battery values is: %v\n", chargesRunningSum)
}
