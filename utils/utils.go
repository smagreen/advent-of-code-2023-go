package utils

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func getSessionCookie() string {
	sessionCookie := os.Getenv("AOC_SESSION_COOKIE")
	if sessionCookie == "" {
		log.Fatal("Please set the AOC_SESSION_COOKIE environment variable")
	}

	return sessionCookie
}

func GetAoCInputHTTP(day int) []string {
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	sessionCookie := "session=" + getSessionCookie()
	req.Header.Set("Cookie", sessionCookie)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// The body variable contains your puzzle input
	//log.Println(string(body))

	var lines []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func GetAoCInputFile(location string) []string {
	file, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
