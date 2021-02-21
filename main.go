package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func cipher(operation string, text string, key int) string {

	shift, offset := rune(key), rune(26)

	runes := []rune(text)

	for index, char := range runes {

		switch operation {
		case "decrypt":
			if char >= 'a'+shift && char <= 'z' ||
				char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift ||
				char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		case "encrypt":
			if char >= 'a' && char <= 'z'-shift ||
				char >= 'A' && char <= 'Z'-shift {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' ||
				char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		}

		runes[index] = char
	}

	return string(runes)
}

func main() {

	var command string
	var t string
	var text string

	_, err := fmt.Scan(&command)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Scan(&t)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()

	text2 := strings.Split(text, "-")

	key := strings.Split(text2[1], " ")
	i, err := strconv.Atoi(key[1])
	if err != nil {
		log.Fatal(err)
	}

	res := cipher(command, text2[0], i)

	print(res)

}
