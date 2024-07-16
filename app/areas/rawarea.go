package areas

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"google.golang.org/api/gmail/v1"
)

// TMessageRawArea defines a structure to store information about a Gmail message.
type TMessageRawArea struct {
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
	// Raw: The entire email message in an RFC 2822 formatted.
	Raw string `json:"raw,omitempty"`
}

// PrepareAllArea takes a Gmail message and returns a TMessageRawArea structure with the fields populated.
func PrepareRawArea(m *gmail.Message) (TMessageRawArea, error) {
	pm := new(TMessageRawArea)
	pm.Id = m.Id
	pm.InternalDate = m.InternalDate
	pm.LabelIds = m.LabelIds
	pm.SizeEstimate = m.SizeEstimate
	pm.Snippet = m.Snippet
	pm.ThreadId = m.ThreadId

	raw, err := base64.URLEncoding.DecodeString(m.Raw)
	if err != nil {
		return *pm, err
	}
	pm.Raw = string(raw)
	return *pm, nil
}

// String method returns a formatted string representation of TMessageRawArea
func (Ma TMessageRawArea) String() string {
	St := ""
	St = St + fmt.Sprintf("%s: %s\r\n", "ID", Ma.Id)
	St = St + fmt.Sprintf("%s: %v\r\n", "Internal Date", Ma.InternalDate)
	St = St + fmt.Sprintf("%s: ", "Label IDs")
	for _, label := range Ma.LabelIds {
		St = St + fmt.Sprintf("%s, ", label)
	}
	St = St + fmt.Sprintln()
	St = St + fmt.Sprintf("%s: %v\r\n", "Size Estimate", Ma.SizeEstimate)
	St = St + fmt.Sprintf("%s: %s\r\n", "Snippet", Ma.Snippet)
	St = St + fmt.Sprintf("%s: %s\r\n", "Thread ID", Ma.ThreadId)
	St = St + fmt.Sprintf("%s:\r\n", "--- Raw Body ---")
	St = St + fmt.Sprintf("%s\r\n", Ma.Raw)
	return St
}

// ToJson method converts the TMessageRawArea structure to a JSON byte array.
func (Ma TMessageRawArea) ToJson() ([]byte, error) {
	b, err := json.Marshal(Ma)
	return b, err
}

// ToTxt method converts the TMessageRawArea structure to a plain text byte array.
func (Ma TMessageRawArea) ToTxt() ([]byte, error) {
	b := []byte(Ma.String())
	return b, nil
}
