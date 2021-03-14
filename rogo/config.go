package rogo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

// Number is the recipient of the outbound call
var Number string

// Transcription is the path of the file to be used
var Transcription string

// URL is the API endpoint with the account id
var URL string

const (
	configFile   = ".rogo"
	configString = "ACCOUNT_ID=%s\nAUTH_TOKEN=%s\nFROM_NUMBER=%s"
	endpoint     = "https://api.twilio.com/2010-04-01/Accounts/%s/Calls.json"
)

// LoadConfig either loads the file from the home folder
// or shows setup to generate a new one
func LoadConfig() {
	path := fmt.Sprintf("%s/%s", getHomeFolder(), configFile)

	if _, err := os.Stat(path); err == nil {
		parseConfig(path)
	} else if os.IsNotExist(err) {
		generateConfig(path)
		os.Exit(0)
	} else {
		fatalError("Unexpected error loading configuration")
	}
}

func parseConfig(path string) {
	data, err := ioutil.ReadFile(path)
	checkFatal(err, "Error loading config file")

	viper.SetConfigType("env")
	if err := viper.ReadConfig(bytes.NewBuffer(data)); err != nil {
		checkFatal(err, "Error reading config file")
	}

	URL = fmt.Sprintf(endpoint, viper.Get("ACCOUNT_ID"))
}

func generateConfig(path string) {
	var (
		accountID  string
		authToken  string
		fromNumber string
	)

	fmt.Print("Enter account id: ")
	fmt.Scan(&accountID)

	fmt.Print("Enter auth token: ")
	fmt.Scan(&authToken)

	fmt.Print("Enter from number: ")
	fmt.Scan(&fromNumber)

	viper.SetConfigType("env")
	viper.Set("ACCOUNT_ID", accountID)
	viper.Set("AUTH_TOKEN", authToken)
	viper.Set("FROM_NUMBER", fromNumber)

	config := fmt.Sprintf(configString, accountID, authToken, fromNumber)
	err := ioutil.WriteFile(path, []byte(config), 0644)
	checkFatal(err, "Error creating configuration file")
}
