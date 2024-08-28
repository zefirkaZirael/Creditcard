package main

import (
	"fmt"
	"os"
	"time"
	"unicode"

	//"sort"
	"bufio"
	"math/rand"
	"strconv"
	"strings"
)

// /// VALIDATE
func luhnCheck(cardNumber string) bool {
	if len(cardNumber) < 13 {
		return false
	}

	var sum int
	alt := false

	for i := len(cardNumber) - 1; i >= 0; i-- {

		num := int(cardNumber[i] - '0')
		if alt {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
		// fmt.Println(sum)
		alt = !alt
	}
	return sum%10 == 0
}

func validate(cards []string) {
	for _, card := range cards {
		if luhnCheck(card) {
			fmt.Println("OK")
		} else {
			fmt.Println("INCORRECT")
			os.Exit(1)
		}
	}
	os.Exit(0)
}

// /////// GENERATE
func validate2(card string) bool {
	for i, r := range card {
		if !unicode.IsDigit(r) {
			// Allow '*' symbols only at the end of the card number
			if r == '*' {
				// Ensure that all characters after this are also '*'
				for j := i; j < len(card); j++ {
					if card[j] != '*' {
						return false
					}
				}
				break // No need to continue the loop once we find the first '*'
			}
			return false
		}
	}
	return true
}

func luhnCheck2(cardNumber string, j int) int {
	var sum int
	alt := false
	if j%2 != 0 {
		alt = true
	}

	for i := len(cardNumber) - 1; i >= 0; i-- {
		num := int(cardNumber[i] - '0')
		if alt {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
		alt = !alt
	}

	res := 10 - (sum % 10)
	return res
}

func generate(cardNumber string, peek bool) {
	j := 0
	runes := []rune(cardNumber) // Convert string to []rune to allow modification
	if len(runes) < 13 {
		fmt.Println("INCORRECT")
		os.Exit(1)
	}
	// Loop through the cardNumber string
	if validate2(cardNumber) {
		for i := len(runes) - 1; i > 0; i-- {
			if runes[i] == '*' {
				j++
				runes[i] = '0' // Replace '*' with '0'
			}
		}

		if j > 4 || j == 0 {
			fmt.Fprintln(os.Stderr, "Error: Invalid number of asterisks.")
			os.Exit(1)
		}

		var generated []string
		if j == 1 {
			base := string(runes[:len(runes)-1])
			res := luhnCheck2(base, j) // res type int
			if res < 10 {
				new_res := strconv.Itoa(res)
				base = base + new_res
				generated = append(generated, base)
			}
		} else if j == 2 {
			for i1 := 0; i1 < 9; i1++ {
				runes[len(runes)-2] = rune(i1 + '0')
				base := string(runes[:len(runes)-1])
				res := luhnCheck2(base, j) // res type int
				if res < 10 {
					new_res := strconv.Itoa(res)
					base = base + new_res

					generated = append(generated, base)
				}
			}
		} else if j == 3 {
			for i1 := 0; i1 < 99; i1++ {
				base := string(runes[:len(runes)-3])
				newbase := base + fmt.Sprintf("%02d", i1)
				res := luhnCheck2(newbase, j) // res type int
				if res < 10 {
					new_res := strconv.Itoa(res)
					base = newbase + new_res
					generated = append(generated, base)
				}
			}
		} else if j == 4 {
			for i1 := 0; i1 < 999; i1++ {
				base := string(runes[:len(runes)-4])
				newbase := base + fmt.Sprintf("%03d", i1)
				res := luhnCheck2(newbase, j) // res type int
				if res < 10 {
					new_res := strconv.Itoa(res)
					base = newbase + new_res
					exists := false
					for _, r := range generated {
						if r == base {
							exists = true
							break
						}
					}
					// If 'newRune' is not in the slice, append it
					if !exists {
						generated = append(generated, base)
					}
				}
			}
		}
		if !peek {
			for _, number := range generated {
				fmt.Println(number)
				if number == "0" {
					fmt.Println()
				}
			}
		} else {
			rand.Seed(time.Now().UnixNano())
			randomIndex := rand.Intn(len(generated))
			fmt.Println(generated[randomIndex])
		}
	} else {
		fmt.Println("INCORRECT card Number")
		os.Exit(1)
	}
}

// /////// INFORMATION

func loadPrefixes2(filename string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	prefixes := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			prefixes[parts[1]] = parts[0]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	return prefixes
}

func findBrandAndIssuer(cardNumber string, brands, issuers map[string]string) (string, string) {
	brand, issuer := "-", "-"

	for prefix, name := range brands {
		if strings.HasPrefix(cardNumber, prefix) {

			brand = name
			break
		}
	}

	for prefix, name := range issuers {
		if strings.HasPrefix(cardNumber, prefix) {

			issuer = name
			break
		}
	}

	return brand, issuer
}

func information(cardNumber, brandsFile, issuersFile string) {
	brands := loadPrefixes2(brandsFile)
	issuers := loadPrefixes2(issuersFile)

	valid := luhnCheck(cardNumber)
	brand, issuer := findBrandAndIssuer(cardNumber, brands, issuers)

	fmt.Println(cardNumber)
	if valid {
		fmt.Println("Correct: yes")
		fmt.Printf("Card Brand: %s\n", brand)
		fmt.Printf("Card Issuer: %s\n", issuer)
	} else {
		fmt.Println("Correct: no")
		fmt.Printf("Card Brand: -\n")
		fmt.Printf("Card Issuer: -\n")
	}
}

///////// ISSUE
func loadPrefixes(filename string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	prefixes := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			prefixes[parts[0]] = parts[1]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	return prefixes
}

func randomDigit() int {
	return rand.Intn(10)
}

// Function to complete the card number using Luhn's algorithm
func completeLuhn(cardNumber string) string {
	sum := 0
	alt := true
	for i := len(cardNumber) - 1; i >= 0; i-- {
		num := int(cardNumber[i] - '0')
		if alt {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		sum += num
		alt = !alt
	}

	// Find the last digit that makes the card valid
	lastDigit := (10 - (sum % 10)) % 10
	return cardNumber + strconv.Itoa(lastDigit)
}

func generateCardNumber(prefix string, length int) string {
	randomPartLength := length - len(prefix) - 1
	randomPart := ""

	for i := 0; i < randomPartLength; i++ {
		randomPart += strconv.Itoa(randomDigit())
	}

	partialCardNumber := prefix + randomPart
	return completeLuhn(partialCardNumber)
}

func main() {
	getValueAfterEquals := func(arg string) (string, error) {
		parts := strings.Split(arg, "=")
		if len(parts) < 2 {
			return "", fmt.Errorf("argument '%s' does not contain '=' or is malformed", arg)
		}
		return parts[1], nil
	}

	if len(os.Args) > 2 {
		if os.Args[1] == "validate" {
			// Handle "--stdin" for validation
			if len(os.Args) > 2 && os.Args[2] == "--stdin" {
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					line := scanner.Text()
					cards := strings.Fields(line) // Split by space to handle multiple cards
					validate(cards)
				}
			} else {
				// Validate passed card numbers from command line
				validate(os.Args[2:])
			}
		} else if os.Args[1] == "generate" {
			// Handle "--pick" option
			if len(os.Args) > 3 && os.Args[2] == "--pick" {
				generate(os.Args[3], true) // Generate with the --pick flag
			} else if len(os.Args) > 2 {
				// Generate without the --pick flag
				generate(os.Args[2], false)
			} else {
				fmt.Println("Usage: ./creditcard generate [--pick] <card_pattern>")
				os.Exit(1)
			}
		} else if os.Args[1] == "information" {
			// Handle "--pick" option
			if len(os.Args) > 4 {
				var err error
				var brandsFile, issuersFile, cardNumber string

				if brandsFile, err = getValueAfterEquals(os.Args[2]); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				if issuersFile, err = getValueAfterEquals(os.Args[3]); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				brandsFile = strings.Split(os.Args[2], "=")[1]
				issuersFile = strings.Split(os.Args[3], "=")[1]
				cardNumber = os.Args[4]
				information(cardNumber, brandsFile, issuersFile)
			} else {
				fmt.Println("Usage: ./creditcard information --brands=brands.txt --issuers=issuers.txt <card_number>")
				os.Exit(1)
			}
		} else if os.Args[1] == "issue" {
			// Handle "--pick" option
			if len(os.Args) > 5 {
				// Function to safely extract the value after '='

				var err error
				var brandsFile, issuersFile, brand, issuer string

				if brandsFile, err = getValueAfterEquals(os.Args[2]); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				if issuersFile, err = getValueAfterEquals(os.Args[3]); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if brand, err = getValueAfterEquals(os.Args[4]); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				if issuer, err = getValueAfterEquals(os.Args[5]); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				brands := loadPrefixes(brandsFile)
				issuers := loadPrefixes(issuersFile)

				_, brandExists := brands[brand]               // Correct usage of map key (index)
				issuerPrefix, issuerExists := issuers[issuer] // Correct usage of map key (index)

				if !brandExists || !issuerExists {
					fmt.Fprintf(os.Stderr, "Error: Brand or Issuer not found.\n")
					os.Exit(1)
				}

				rand.Seed(time.Now().UnixNano()) // Seed random number generator

				// Assuming standard credit card length is 16 digits
				cardNumber := generateCardNumber(issuerPrefix, 16)

				fmt.Println(cardNumber)
			} else {
				fmt.Fprintf(os.Stderr, "Usage: ./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=<BRAND> --issuer=<ISSUER>\n")
				os.Exit(1)
			}
		}
	}
}
