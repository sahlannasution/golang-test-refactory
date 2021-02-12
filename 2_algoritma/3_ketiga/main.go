/*
	Maaf kiranya saya salah mengerti dengan maksud soalnya dan maaf jika masih manual pembagian arraynya
	krna tidak disebutkan harus dengan perulangan atau harus bagi 3 sama rata array nya
*/

package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

func sorting(w string) (string, []string) {
	s := strings.Split(w, "")
	sort.Strings(s)

	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return strings.Join(s, ""), list
}

func countCharacters(w string, data []string) {
	for i := 0; i < len(data); i++ {
		alpha := strings.Count(w, data[i])
		fmt.Printf("Karakter %s sebanyak %d kali\n", data[i], alpha)
	}
}

// func findCharacter(char []string, val string) int || string{
// 	for j, item := range char {
// 		if item == val {
// 			return j + 5
// 		}
// 	}
// 	ret
// }
func main() {
	fmt.Println("OK")
	s := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras interdum mi eu magna fermentum, vel luctus tellus semper. Nunc dignissim eleifend ipsum, nec viverra mauris pellentesque non. Fusce auctor ex id mauris egestas, quis luctus nunc pharetra. Sed in dignissim nisi. Aliquam sed tempor urna, nec aliquam mi. Aenean eu feugiat lacus, vel dictum eros. Nulla condimentum porttitor aliquet. Vestibulum vehicula elit non arcu auctor maximus. Quisque est eros, maximus nec diam faucibus, mollis odio."
	small := strings.ToLower(s)
	s = strings.ToUpper(s)

	// fmt.Println(s)
	// fmt.Printf("\n")

	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	onlyAlphabets := (reg.ReplaceAllString(s, ""))
	// fmt.Println(onlyAlphabets)
	// fmt.Printf("\n")

	srt, unique := sorting(onlyAlphabets)
	// fmt.Println(srt)
	// fmt.Println(unique)
	fmt.Println("Jumlah per Karakter muncul :")
	countCharacters(srt, unique)
	fmt.Printf("\n")

	character := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	fmt.Println("Setiap karakter geser 5 karakter : (kecuali v w x y z krna mereka adalah 5 karakter terakhir)")
	count := strings.Split(small, "")
	// fmt.Println(count)
	result := ""
	panjang := len(character)
	fmt.Println(strings.Join(count, ""))
	fmt.Printf("\n")

	for i := 0; i < len(count); i++ {
		temp := 0
		for j := 0; j < panjang; j++ {
			if character[j] == count[i] {
				// temp = j + 5
				if j+5 >= panjang {
					temp = 0
				} else {
					// fmt.Println(j)
					temp = j + 5
				}

				break
			}
		}

		if temp != 0 {
			result += character[temp]

		} else {
			result += count[i]
		}
	}

	fmt.Println(result)
}
