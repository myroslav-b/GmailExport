package main

import (
	"errors"
	"gmailexport/app/areas"

	"google.golang.org/api/gmail/v1"
)

// iAreaMolder interface defines methods for converting message data to different formats
type iAreaMolder interface {
	ToJson() ([]byte, error)
	ToTxt() ([]byte, error)
}

// performance processes a list of messages according to the given statement
// and returns the formatted output as a slice of byte slices
func performance(listMessages *tListMessages, statement tStatement) ([][]byte, error) {
	outBlocks := make([][]byte, 0)
	for _, message := range listMessages.messages {
		preparedMessage, err := prepareMessage(message, statement.Area)
		if err != nil {
			return outBlocks, err
		}
		block, err := toFormat(preparedMessage, statement.Format)
		if err != nil {
			return outBlocks, err
		}
		outBlocks = append(outBlocks, block)
	}
	return outBlocks, nil
}

// prepareMessage prepares a Gmail message according to the specified area
func prepareMessage(message *gmail.Message, area string) (iAreaMolder, error) {
	var preparedMessage iAreaMolder
	//var preparedMessage tMessageAllArea
	var err error
	switch area {
	case "small":
		preparedMessage, err = areas.PrepareSmallArea(message)
		if err != nil {
			return nil, err
		}
		return preparedMessage, nil
	case "easy":
		preparedMessage, err = areas.PrepareEasyArea(message)
		if err != nil {
			return nil, err
		}
		return preparedMessage, nil
	case "all":
		preparedMessage, err = areas.PrepareAllArea(message)
		if err != nil {
			return nil, err
		}
		return preparedMessage, nil
	case "raw":
		preparedMessage, err = areas.PrepareRawArea(message)
		if err != nil {
			return nil, err
		}
		return preparedMessage, nil
	default:
		return nil, errors.New("undefined parameter Area")
	}
}

// toFormat converts a prepared message to the specified format (JSON or TXT)
func toFormat(prepMessages iAreaMolder, format string) ([]byte, error) {
	switch format {
	case "json":
		bytes, err := prepMessages.ToJson()
		if err != nil {
			return nil, err
		}
		return bytes, nil
	case "txt":
		bytes, err := prepMessages.ToTxt()
		if err != nil {
			return nil, err
		}
		return bytes, nil
	default:
		return nil, errors.New("undefined parameter Format")
	}
}
