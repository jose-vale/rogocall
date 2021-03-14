# RoGOcall
[![CodeQL](https://github.com/jose-vale/rogocall/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/jose-vale/rogocall/actions/workflows/codeql-analysis.yml)

The app will allow the user to place a call from the terminal using Twilio, and leverages the TwiML API to convert the transcriptions to speech.

## Instalation

Download the latest release [here](https://github.com/jose-vale/rogocall/releases/latest).  
For macOS users, either add the executable to the shell config, or place the file in `/usr/local/bin`.  
For Windows users add the rogo folder to the PATH environment variable.

## Setup

On the first run the app will ask the user to input the following values:

- Account ID: The twilio project account number
- Auth Token: The twilio project authenticaton token
- From Number: The number to be used for the outbound calls

The setup process will create a `.rogo` file in the user's home folder. Deleting it will force the setup to run on the next execution.

## Usage

To perform an outbound call simply run the command with the necessary arguments from the command line:

```bash
rogo [outbound number] [transcriptions xml file]
```

Run `rogo -h` to show the help.

## Comments

rogo will verify if the number is in a correct format, as well as the extension of the provided transcription:

- Phone number starts with `+` and is only digits, ex.: +1555555
- Provided transcription is a `.xml` file, ex.: transcription.xml

For information regarding the transcription options, check [TwiML](https://www.twilio.com/docs/voice/twiml).
