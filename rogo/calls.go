package rogo

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

// SendTranscription will make a request to the Twilio API
// to process the provided transcription file and call
// the provided number
func SendTranscription() {
	values := url.Values{}
	values.Set("To", Number)
	values.Set("From", fmt.Sprint(viper.Get("FROM_NUMBER")))
	values.Set("Twiml", getTranscription())
	body := *strings.NewReader(values.Encode())

	accountID := fmt.Sprint(viper.Get("ACCOUNT_ID"))
	authToken := fmt.Sprint(viper.Get("AUTH_TOKEN"))

	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, &body)
	checkFatal(err, "Failed to contact API")

	req.SetBasicAuth(accountID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	checkFatal(err, "Failed to contact API")
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println(string(colorGreen), "Sent call request")
	} else {
		printBody(&resp.Body)
		fatalError("Call request failed")
	}
}
