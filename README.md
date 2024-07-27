# Gmail Export Tool

## Disclaimer

The project was created for my own needs and as an exercise in using the Google API. I apologize if the quality of the code seems insultingly low to you. But it works ;-)

## Overview

The Gmail Export Tool is a Go application designed to search, retrieve, and export Gmail messages based on various criteria. It offers flexible filtering options and multiple output formats, making it solution for users who need to extract and analyze their Gmail data.

## Features

- **Advanced Search**: Filter messages by ID, label, sender, recipient, and subject.
- **Multiple Output Formats**: Export data in JSON or TXT format.
- **Customizable Output Areas**: Choose between different levels of message detail (raw, small, easy, all).
- **Flexible Output Options**: Write to stdout or files, with the option to split results into multiple files.
- **OAuth 2.0 Authentication**: Secure access to Gmail API using Google's OAuth 2.0 protocol.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go 1.15 or higher
- Access to a Google Cloud Project with the Gmail API enabled
- OAuth 2.0 credentials (client ID and client secret) for your Google Cloud Project

Quickstarts explain how to set up and run an app that calls a Google Workspace API:
https://developers.google.com/gmail/api/quickstart/go


## Installation

1. Clone the repository:
   ```
   git clone https://github.com/myroslav-b/gmailexport.git
   cd gmailexport
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Place your `credentials.json` file (containing OAuth 2.0 client ID and client secret) in the project root directory.

## Usage

Build the tool using the following command:

```
go build -o gmailexport ./app
```

Run the tool:

```
./gmailexport [options]
```

### Options

All startup options can be found by running the command:

```
./gmailexport -h
```
#### Selection Conditions:
- `-m, --message`: Message ID
- `-l, --label`: Label
- `-f, --from`: Sender's email address
- `-t, --to`: Recipient's email address
- `-s, --subject`: Email subject

#### Presentation of Results:
- `-O, --output`: Output path
  - Use "stdout" if missing, else output to file
  - Specify a file path for file output
  - Use "gmail" if option occurs without an argument
- `-S, --split`: Split output into multiple files
- `-F, --format`: Output format (choices: "json", "txt", default: "json")
- `-A, --area`: Fullness of the output (choices: "raw", "all", "small", "easy", default: "all")

### Examples

1. Search for emails from a specific sender and export as JSON:
   ```
   ./gmaiexport --from example@email.com --format json --output results.json
   ```

2. Export all emails with a specific label as TXT, split into multiple files:
   ```
   ./gmaiexport --label important --format txt --output exported_emails --split
   ```

3. Search for emails with a specific subject and display raw output to console:
   ```
   ./gmaiexport --subject "Meeting Notes" --area raw
   ```

