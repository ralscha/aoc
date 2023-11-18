package download

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func ReadInput(fileName string, year, day int) (string, error) {
	_, err := os.Stat(fileName)
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

	req.Header.Add("User-Agent", "github.com/ralscha/go-playground by me@rasc.ch")

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
	return string(bodyBytes), nil
}
