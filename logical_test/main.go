package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Expected: False, Result: %t\n", validatePincode("17283"))
	fmt.Printf("Expected: True, Result: %t\n", validatePincode("172839"))
	fmt.Printf("Expected: False, Result: %t\n", validatePincode("111822"))
	fmt.Printf("Expected: True, Result: %t\n", validatePincode("112762"))
	fmt.Printf("Expected: False, Result: %t\n", validatePincode("123743"))
	fmt.Printf("Expected: False, Result: %t\n", validatePincode("321895"))
	fmt.Printf("Expected: True, Result: %t\n", validatePincode("124578"))
	fmt.Printf("Expected: False, Result: %t\n", validatePincode("112233"))
	fmt.Printf("Expected: False, Result: %t\n", validatePincode("882211"))
	fmt.Printf("Expected: True, Result: %t\n", validatePincode("887712"))
}

func validatePincode(pincode string) bool {
	_, err := strconv.Atoi(pincode)
	if err != nil {
		return false
	}
	if len(pincode) < 6 {
		return false
	}
	samenumCount := 0
	pairCount := 0

	sortCount := 1
	sortReverseCount := 1

	var temp int
	for i, char := range pincode {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		if i == 0 {
			temp = num
			continue
		}
		if temp == num {
			samenumCount++
			pairCount++
		} else {
			samenumCount = 0
		}
		if temp-num == 1 {
			sortReverseCount++
		} else {
			sortReverseCount = 1
		}
		if temp-num == -1 {
			sortCount++
		} else {
			sortCount = 1
		}

		if samenumCount > 2 {
			return false
		}
		if pairCount > 2 {
			return false
		}
		if sortCount > 2 {
			return false
		}
		if sortReverseCount > 2 {
			return false
		}
		temp = num
	}
	return true
}
