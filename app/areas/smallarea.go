package areas

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"google.golang.org/api/gmail/v1"
)

// TMessageSmallArea defines a structure to store information about a Gmail message.
type TMessageSmallArea struct {
	// Id: The immutable ID of the message.
	Id string `json:"id,omitempty"`
	// InternalDate: The internal message creation timestamp (epoch ms), which
	// determines ordering in the inbox. For normal SMTP-received email, this
	// represents the time the message was originally accepted by Google, which is
	// more reliable than the `Date` header. However, for API-migrated mail, it can
	// be configured by client to be based on the `Date` header.
	InternalDate int64 `json:"internalDate,omitempty,string"`
	// LabelIds: List of IDs of labels applied to this message.
	LabelIds []string `json:"labelIds,omitempty"`
	// SizeEstimate: Estimated size in bytes of the message.
	SizeEstimate int64 `json:"sizeEstimate,omitempty"`
	// Snippet: A short part of the message text.
	Snippet string `json:"snippet,omitempty"`
	// ThreadId: The ID of the thread the message belongs to. To add a message or
	// draft to a thread, the following criteria must be met: 1. The requested
	// `threadId` must be specified on the `Message` or `Draft.Message` you supply
	// with your request. 2. The `References` and `In-Reply-To` headers must be set
	// in compliance with the RFC 2822 (https://tools.ietf.org/html/rfc2822)
	// standard. 3. The `Subject` headers must match.
	ThreadId string `json:"threadId,omitempty"`
	// Message-ID
	MessageId string `json:"messageId,omitempty"`
	// Date
	Date string `json:"date,omitempty"`
	// From
	From string `json:"from,omitempty"`
	// To
	To string `json:"to,omitempty"`
	// Subject
	Subject string `json:"subject,omitempty"`
	// PlainText:
	PlainText string `json:"plainText,omitempty"`
}

// PrepareAllArea takes a Gmail message and returns a TMessageSmallArea structure with the fields populated.
func PrepareSmallArea(m *gmail.Message) (TMessageSmallArea, error) {
	pm := new(TMessageSmallArea)
	var err error
	pm.Id = m.Id
	pm.InternalDate = m.InternalDate
	pm.LabelIds = m.LabelIds
	pm.SizeEstimate = m.SizeEstimate
	pm.Snippet = m.Snippet
	pm.ThreadId = m.ThreadId
	for _, h := range m.Payload.Headers {
		switch h.Name {
		case "Message-ID":
			pm.MessageId = h.Value
		case "Date":
			pm.Date = h.Value
		case "From":
			pm.From = h.Value
		case "To":
			pm.To = h.Value
		case "Subject":
			pm.Subject = h.Value
		}
	}
	bPlainText, err := base64.URLEncoding.DecodeString(getPlainTextBody(m.Payload))
	if err != nil {
		pm.PlainText = ""
	} else {
		pm.PlainText = string(bPlainText)
	}
	return *pm, nil
}

// String method returns a formatted string representation of TMessageSmallArea
func (Ma TMessageSmallArea) String() string {
	St := ""
	St = St + fmt.Sprintf("%s: %s\r\n", "ID", Ma.Id)
	St = St + fmt.Sprintf("%s: %v\r\n", "Internal Date", Ma.InternalDate)
	St = St + fmt.Sprintf("%s: ", "Label IDs")
	for _, label := range Ma.LabelIds {
		St = St + fmt.Sprintf("%s, ", label)
	}
	St = St + fmt.Sprintf("%s\r\n", "")
	St = St + fmt.Sprintf("%s: %v\r\n", "Size Estimate", Ma.SizeEstimate)
	St = St + fmt.Sprintf("%s: %s\r\n", "Snippet", Ma.Snippet)
	St = St + fmt.Sprintf("%s: %s\r\n", "Thread ID", Ma.ThreadId)
	St = St + fmt.Sprintf("%s\r\n", "--- Headers ---")
	St = St + fmt.Sprintf("%s: %s\r\n", "Message-ID", Ma.MessageId)
	St = St + fmt.Sprintf("%s: %s\r\n", "Date", Ma.Date)
	St = St + fmt.Sprintf("%s: %s\r\n", "From", Ma.From)
	St = St + fmt.Sprintf("%s: %s\r\n", "To", Ma.To)
	St = St + fmt.Sprintf("%s: %s\r\n", "Subject", Ma.Subject)
	St = St + fmt.Sprintf("%s\r\n%s\r\n", "--- Plain Text ---", Ma.PlainText)
	return St
}

// ToJson method converts the TMessageSmallArea structure to a JSON byte array.
func (Ma TMessageSmallArea) ToJson() ([]byte, error) {
	b, err := json.Marshal(Ma)
	return b, err
}

// ToTxt method converts the TMessageSmallArea structure to a plain text byte array.
func (Ma TMessageSmallArea) ToTxt() ([]byte, error) {
	b := []byte(fmt.Sprintf("%+v", Ma))
	return b, nil
}
