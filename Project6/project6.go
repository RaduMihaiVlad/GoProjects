package main


import (
	"fmt"
)

func IsDigit(val rune) bool {
	
	zeroVal := int('0')
	nineVal := int('9')
	digVal := int(val)
	if zeroVal <= digVal && digVal <= nineVal {
		return true
	}
	return false
}

func Normalize(phoneNumber string) string {

	normalizedPhoneNumber := ""

	for _, val := range phoneNumber {
		if ok := IsDigit(val); ok {
			normalizedPhoneNumber = normalizedPhoneNumber + string(val)
		}
	}
	return normalizedPhoneNumber
}

func main() {
	
	fmt.Println(Normalize("123-456-789"))

}