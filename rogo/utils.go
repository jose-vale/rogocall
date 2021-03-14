package rogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"regexp"
	"strings"
)

const (
	// Example is used to show how to use the app in errors
	Example = "Usage: rogo [outbound number] [transcription xml file]"

	colorReset = "\033[0m"
	colorGreen = "\033[32m"
	colorGray  = "\033[37m"
	colorRed   = "\033[31m"
)

// ErrorWithExample prints the provided error message
// with instructions on how to use the app
func ErrorWithExample(err string) {
	msg := fmt.Sprintf("%s %s(%s)", err, string(colorGray), Example)
	exitWithError(msg)
}

// IsValidNumber checks if the provided number starts
// with a + and is only digits
func IsValidNumber() bool {
	regex := regexp.MustCompile(`^[+][0-9]+$`)
	return regex.MatchString(Number)
}

// HasCorrectExtension checks if the provided file
// has the .xml extension
func HasCorrectExtension() bool {
	lowercase := strings.ToLower(Transcription)
	return bytes.HasSuffix([]byte(lowercase), []byte(".xml"))
}

func exitWithError(msg string) {
	fmt.Println(string(colorRed), msg, string(colorReset))
	os.Exit(1)
}

func checkFatal(err error, msg string) {
	if err != nil {
		exitWithError(msg)
	}
}

func fatalError(msg string) {
	exitWithError(msg)
}

func getHomeFolder() string {
	usr, err := user.Current()
	checkFatal(err, "Error loading system information")

	return usr.HomeDir
}

func getTranscription() string {
	xmlFile, err := os.Open(Transcription)
	checkFatal(err, "Error loading transcription file")
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	checkFatal(err, "Error reading transcription file")

	return string(byteValue)
}

func printBody(body *io.ReadCloser) {
	var data map[string]interface{}

	bytes, err := ioutil.ReadAll(*body)
	if err == nil {
		if err := json.Unmarshal(bytes, &data); err == nil {
			prettyJSON, err := json.MarshalIndent(data, "", "  ")
			if err == nil {
				fmt.Printf("%s\n", prettyJSON)
			}
		}
	}
}
