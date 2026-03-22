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

type AdditionalColumn struct {
	header string
	values []string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func generateRandomName() string {
	firstName := generateRandomString(rand.Intn(6) + 3)
	lastName := generateRandomString(rand.Intn(6) + 3)
	return firstName + " " + lastName
}

func generateRandomEmail(name, domain string) string {
	emailName := strings.ToLower(strings.ReplaceAll(name, " ", "."))
	randomNum := fmt.Sprintf("%03d", rand.Intn(1000))
	return fmt.Sprintf("%s%s@%s", emailName, randomNum, domain)
}

func main() {
	_, err := fmt.Fprint(os.Stdout, "How many entries do you need? > ")
	check(err)

	countString, err := bufio.NewReader(os.Stdin).ReadString('\n')
	check(err)

	count, err := strconv.Atoi(strings.TrimSpace(countString[:len(countString)-1]))
	check(err)

	_, err = fmt.Fprint(os.Stdout, "Which domain should the emails use? (default: example.com) > ")
	check(err)

	domainInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	check(err)

	domain := strings.TrimSpace(domainInput[:len(domainInput)-1])

	if len(domain) == 0 {
		domain = "example.com"
	}

	var additionalColumns []AdditionalColumn

	_, err = fmt.Fprint(os.Stdout, "Need more columns? Provide comma + space separated column headers > ")
	check(err)

	columnHeadersInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	check(err)

	columnHeaders := strings.SplitSeq(columnHeadersInput, ", ")

	for columnHeader := range columnHeaders {
		header := strings.TrimSpace(columnHeader)
		_, err = fmt.Fprint(os.Stdout, "Provide comma + space separated values to randomly choose from for column: "+header+" > ")
		check(err)

		columnValuesInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		check(err)

		columnValues := strings.Split(columnValuesInput, ", ")

		for i, columnValue := range columnValues {
			columnValues[i] = strings.TrimSpace(columnValue)
		}

		newCol := AdditionalColumn{header: header, values: columnValues}
		additionalColumns = append(additionalColumns, newCol)
	}

	filename := "generated_data_" + strconv.Itoa(count) + ".csv"
	file, err := os.Create(filename)
	check(err)

	defer func() {
		err := file.Close()
		check(err)
	}()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"name", "email"}

	if len(additionalColumns) > 0 {
		for _, column := range additionalColumns {
			headers = append(headers, column.header)
		}
	}

	err = writer.Write(headers)
	check(err)

	writer.Flush()

	for range count {
		name := generateRandomName()
		email := generateRandomEmail(name, domain)
		results := []string{name, email}

		for _, column := range additionalColumns {
			randomIndex := rand.Intn(len(column.values))
			results = append(results, column.values[randomIndex])
		}

		err = writer.Write(results)
		check(err)
	}

	wd, err := os.Getwd()
	check(err)

	_, err = fmt.Fprintln(os.Stdout, "Done! File written to: "+wd+"/"+filename)
	check(err)
}
