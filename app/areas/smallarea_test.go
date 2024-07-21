package areas

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/gmail/v1"
)

func TestPrepareSmallArea(t *testing.T) {
	// Define a sample Gmail message
	message := &gmail.Message{
		Id:           "12345",
		InternalDate: 1620000000000,
		LabelIds:     []string{"INBOX", "IMPORTANT"},
		SizeEstimate: 2048,
		Snippet:      "This is a snippet",
		ThreadId:     "67890",
		Payload: &gmail.MessagePart{
			Headers: []*gmail.MessagePartHeader{
				{Name: "Message-ID", Value: "<message123@example.com>"},
				{Name: "Date", Value: "Mon, 3 May 2021 10:00:00 +0000"},
				{Name: "From", Value: "sender@example.com"},
				{Name: "To", Value: "recipient@example.com"},
				{Name: "Subject", Value: "Test Email"},
			},
			Parts: []*gmail.MessagePart{
				{
					MimeType: "text/plain",
					Body: &gmail.MessagePartBody{
						Data: base64.URLEncoding.EncodeToString([]byte("Hello, this is a test email!")),
					},
				},
			},
		},
	}

	// Call the function
	result, err := PrepareSmallArea(message)

	// Check no error occurred
	require.NoError(t, err)

	// Verify the result
	assert.Equal(t, "12345", result.Id)
	assert.Equal(t, int64(1620000000000), result.InternalDate)
	assert.ElementsMatch(t, []string{"INBOX", "IMPORTANT"}, result.LabelIds)
	assert.Equal(t, int64(2048), result.SizeEstimate)
	assert.Equal(t, "This is a snippet", result.Snippet)
	assert.Equal(t, "67890", result.ThreadId)
	assert.Equal(t, "<message123@example.com>", result.MessageId)
	assert.Equal(t, "Mon, 3 May 2021 10:00:00 +0000", result.Date)
	assert.Equal(t, "sender@example.com", result.From)
	assert.Equal(t, "recipient@example.com", result.To)
	assert.Equal(t, "Test Email", result.Subject)
	assert.Equal(t, "Hello, this is a test email!", result.PlainText)
}

func TestTMessageSmallArea_String(t *testing.T) {
	message := TMessageSmallArea{
		Id:           "12345",
		InternalDate: 1620000000000,
		LabelIds:     []string{"INBOX", "IMPORTANT"},
		SizeEstimate: 2048,
		Snippet:      "This is a snippet",
		ThreadId:     "67890",
		MessageId:    "<message123@example.com>",
		Date:         "Mon, 3 May 2021 10:00:00 +0000",
		From:         "sender@example.com",
		To:           "recipient@example.com",
		Subject:      "Test Email",
		PlainText:    "Hello, this is a test email!",
	}

	expected := "ID: 12345\r\nInternal Date: 1620000000000\r\nLabel IDs: INBOX, IMPORTANT, \r\nSize Estimate: 2048\r\nSnippet: This is a snippet\r\nThread ID: 67890\r\n--- Headers ---\r\nMessage-ID: <message123@example.com>\r\nDate: Mon, 3 May 2021 10:00:00 +0000\r\nFrom: sender@example.com\r\nTo: recipient@example.com\r\nSubject: Test Email\r\n--- Plain Text ---\r\nHello, this is a test email!\r\n"
	assert.Equal(t, expected, message.String())
}

func TestTMessageSmallArea_ToJson(t *testing.T) {
	message := TMessageSmallArea{
		Id:           "12345",
		InternalDate: 1620000000000,
		LabelIds:     []string{"INBOX", "IMPORTANT"},
		SizeEstimate: 2048,
		Snippet:      "This is a snippet",
		ThreadId:     "67890",
		MessageId:    "<message123@example.com>",
		Date:         "Mon, 3 May 2021 10:00:00 +0000",
		From:         "sender@example.com",
		To:           "recipient@example.com",
		Subject:      "Test Email",
		PlainText:    "Hello, this is a test email!",
	}

	jsonData, err := message.ToJson()
	require.NoError(t, err)

	expected := `{"id":"12345","internalDate":"1620000000000","labelIds":["INBOX","IMPORTANT"],"sizeEstimate":2048,"snippet":"This is a snippet","threadId":"67890","messageId":"<message123@example.com>","date":"Mon, 3 May 2021 10:00:00 +0000","from":"sender@example.com","to":"recipient@example.com","subject":"Test Email","plainText":"Hello, this is a test email!"}`
	assert.JSONEq(t, expected, string(jsonData))
}
func TestTMessageSmallArea_ToTxt(t *testing.T) {
	message := TMessageSmallArea{
		Id:           "12345",
		InternalDate: 1620000000000,
		LabelIds:     []string{"INBOX", "IMPORTANT"},
		SizeEstimate: 2048,
		Snippet:      "This is a snippet",
		ThreadId:     "67890",
		MessageId:    "<message123@example.com>",
		Date:         "Mon, 3 May 2021 10:00:00 +0000",
		From:         "sender@example.com",
		To:           "recipient@example.com",
		Subject:      "Test Email",
		PlainText:    "Hello, this is a test email!",
	}

	txtData, err := message.ToTxt()
	require.NoError(t, err)

	expected := "ID: 12345\r\nInternal Date: 1620000000000\r\nLabel IDs: INBOX, IMPORTANT, \r\nSize Estimate: 2048\r\nSnippet: This is a snippet\r\nThread ID: 67890\r\n--- Headers ---\r\nMessage-ID: <message123@example.com>\r\nDate: Mon, 3 May 2021 10:00:00 +0000\r\nFrom: sender@example.com\r\nTo: recipient@example.com\r\nSubject: Test Email\r\n--- Plain Text ---\r\nHello, this is a test email!\r\n"
	assert.Equal(t, expected, string(txtData))
}
