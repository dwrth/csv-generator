package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func generateRandomName() string {
	// Generate two random strings between 3-8 characters for first/last name
	firstName := generateRandomString(rand.Intn(6) + 3)
	lastName := generateRandomString(rand.Intn(6) + 3)
	return firstName + " " + lastName
}

func generateRandomEmail(name, domain string) string {
	// Convert to lowercase and replace spaces with dots
	emailName := strings.ToLower(strings.ReplaceAll(name, " ", "."))
	// Add random number to make it unique
	randomNum := fmt.Sprintf("%03d", rand.Intn(1000))
	return fmt.Sprintf("%s%s@%s", emailName, randomNum, domain)
}

func main() {
	fmt.Fprint(os.Stdout, "How many entries do you need? > ")

	countString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}

	count, err := strconv.Atoi(strings.TrimSpace(countString[:len(countString)-1]))
	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, "Which domain should the emails use? (default: example.com) > ")

	domainInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}

	domain := strings.TrimSpace(domainInput[:len(domainInput)-1])

	if len(domain) == 0 {
		domain = "example.com"
	}

	file, err := os.Create("generated_data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"name", "email"})

	// Generate 7000 entries
	for range count {
		name := generateRandomName()
		email := generateRandomEmail(name, domain)
		writer.Write([]string{name, email})
	}
}
