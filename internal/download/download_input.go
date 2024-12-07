package download

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// ReadInput retrieves the input for a specific Advent of Code puzzle.
// It first checks if the input exists locally in a file, and if not, downloads it.
// The input is saved to a file in the format: {current_dir}/{year}/{day}/input.txt
//
// Parameters:
//   - year: The year of the Advent of Code puzzle
//   - day: The day of the puzzle (1-25)
//
// Returns the puzzle input as a string and any error encountered.
func ReadInput(year, day int) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fileName := fmt.Sprintf("%s/%d/%02d/input.txt", dir, year, day)
	_, err = os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			input, err := DownloadInput(year, day)
			if err != nil {
				return "", fmt.Errorf("downloading input failed: %v", err)
			}
			err = os.WriteFile(fileName, []byte(input), 0644)
			if err != nil {
				return "", fmt.Errorf("writing input file failed: %v", err)
			}
			return input, nil
		} else {
			return "", err
		}
	} else {
		input, err := os.ReadFile(fileName)
		if err != nil {
			return "", fmt.Errorf("reading input file failed: %v", err)
		}
		return string(input), nil
	}
}

// DownloadInput downloads the input for a specific Advent of Code puzzle from the website.
// It requires a session cookie stored in a file named "sessioncookie" in the current directory.
//
// Parameters:
//   - year: The year of the Advent of Code puzzle
//   - day: The day of the puzzle (1-25)
//
// Returns the downloaded puzzle input as a string and any error encountered.
// The function will trim any trailing newlines from the downloaded content.
func DownloadInput(year, day int) (string, error) {
	sessionCookie, err := os.ReadFile("./sessioncookie")
	if err != nil {
		return "", fmt.Errorf("reading sessioncookie failed: %v", err)
	}
	if sessionCookie[len(sessionCookie)-1] == '\n' {
		sessionCookie = sessionCookie[:len(sessionCookie)-1]
	}
	if sessionCookie[len(sessionCookie)-1] == '\r' {
		sessionCookie = sessionCookie[:len(sessionCookie)-1]
	}

	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		return "", fmt.Errorf("creating request failed: %v", err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = string(sessionCookie)
	req.AddCookie(cookie)

	req.Header.Add("User-Agent", "github.com/ralscha/aoc by me@rasc.ch")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if bodyBytes[len(bodyBytes)-1] == '\n' {
		bodyBytes = bodyBytes[:len(bodyBytes)-1]
	}
	return string(bodyBytes), nil
}
