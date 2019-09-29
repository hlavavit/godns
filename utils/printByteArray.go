package utils

import "fmt"

// PrintByteArray prints byte array to console for inspection
func PrintByteArray(data []byte) {
	fmt.Println("Start of ByteArray")
	for index, dataPart := range data {
		if index%2 == 1 {
			fmt.Printf(" ")
		}
		fmt.Printf("%08b", dataPart)
		if index%2 == 1 || index == len(data)-1 {
			fmt.Printf("\n")
		}
	}
	fmt.Println("End of ByteArray")
}
