package main

import (
	"errors"
	"gmailexport/app/areas"

	"google.golang.org/api/gmail/v1"
)

type iAreaMolder interface {
	//prepare(m gmail.Message) (tMessageAllArea, error)
	ToJson() ([]byte, error)
	ToTxt() ([]byte, error)
}

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

func prepareMessage(message *gmail.Message, area string) (iAreaMolder, error) {
	var preparedMessage iAreaMolder
	//var preparedMessage tMessageAllArea
	var err error
	switch area {
	case "smart":
		return nil, nil
	case "part":
		return nil, nil
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

/*
func performance(listMessages *tListMessages, statement tStatement) error {
	//fmt.Printf("%v\n", listMessages)

	err := output([]byte(listMessages), statement.Output)
	if err != nil {
		return err
	}

	return nil
}

func output(b []byte, st string) error {

	return nil
}
*/
