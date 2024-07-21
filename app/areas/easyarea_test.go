package areas

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/gmail/v1"
)

func TestPrepareEasyArea(t *testing.T) {
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
				{Name: "From", Value: "test@example.com"},
				{Name: "To", Value: "recipient@example.com"},
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
	result, err := PrepareEasyArea(message)

	// Check no error occurred
	require.NoError(t, err)

	// Verify the result
	assert.Equal(t, "12345", result.Id)
	assert.Equal(t, int64(1620000000000), result.InternalDate)
	assert.ElementsMatch(t, []string{"INBOX", "IMPORTANT"}, result.LabelIds)
	assert.Equal(t, int64(2048), result.SizeEstimate)
	assert.Equal(t, "This is a snippet", result.Snippet)
	assert.Equal(t, "67890", result.ThreadId)
	assert.Equal(t, "test@example.com", result.Headers[0].Value)
	assert.Equal(t, "recipient@example.com", result.Headers[1].Value)
	assert.Equal(t, "Hello, this is a test email!", result.PlainText)
}

func TestTMessageEasyArea_String(t *testing.T) {
	message := TMessageEasyArea{
		Id:           "12345",
		InternalDate: 1620000000000,
		LabelIds:     []string{"INBOX", "IMPORTANT"},
		SizeEstimate: 2048,
		Snippet:      "This is a snippet",
		ThreadId:     "67890",
		Headers: []struct {
			Name  string `json:"name,omitempty"`
			Value string `json:"value,omitempty"`
		}{
			{Name: "From", Value: "test@example.com"},
			{Name: "To", Value: "recipient@example.com"},
		},
		PlainText: "Hello, this is a test email!",
	}

	expected := "ID: 12345\r\nInternal Date: 1620000000000\r\nLabel IDs: INBOX, IMPORTANT, \r\nSize Estimate: 2048\r\nSnippet: This is a snippet\r\nThread ID: 67890\r\n--- Headers ---\r\nFrom: test@example.com\r\nTo: recipient@example.com\r\n--- Plain Text ---\r\nHello, this is a test email!\r\n"
	assert.Equal(t, expected, message.String())
}

func TestTMessageEasyArea_ToJson(t *testing.T) {
	message := TMessageEasyArea{
		Id:           "12345",
		InternalDate: 1620000000000,
		LabelIds:     []string{"INBOX", "IMPORTANT"},
		SizeEstimate: 2048,
		Snippet:      "This is a snippet",
		ThreadId:     "67890",
		Headers: []struct {
			Name  string `json:"name,omitempty"`
			Value string `json:"value,omitempty"`
		}{
			{Name: "From", Value: "test@example.com"},
			{Name: "To", Value: "recipient@example.com"},
		},
		PlainText: "Hello, this is a test email!",
	}

	jsonData, err := message.ToJson()
	require.NoError(t, err)

	expected := `{"id":"12345","internalDate":"1620000000000","labelIds":["INBOX","IMPORTANT"],"sizeEstimate":2048,"snippet":"This is a snippet","threadId":"67890","headers":[{"name":"From","value":"test@example.com"},{"name":"To","value":"recipient@example.com"}],"plainText":"Hello, this is a test email!"}`
	assert.JSONEq(t, expected, string(jsonData))
}

func TestTMessageEasyArea_ToTxt(t *testing.T) {
	message := TMessageEasyArea{
		Id:           "12345",
		InternalDate: 1620000000000,
		LabelIds:     []string{"INBOX", "IMPORTANT"},
		SizeEstimate: 2048,
		Snippet:      "This is a snippet",
		ThreadId:     "67890",
		Headers: []struct {
			Name  string `json:"name,omitempty"`
			Value string `json:"value,omitempty"`
		}{
			{Name: "From", Value: "test@example.com"},
			{Name: "To", Value: "recipient@example.com"},
		},
		PlainText: "Hello, this is a test email!",
	}

	txtData, err := message.ToTxt()
	require.NoError(t, err)

	expected := "ID: 12345\r\nInternal Date: 1620000000000\r\nLabel IDs: INBOX, IMPORTANT, \r\nSize Estimate: 2048\r\nSnippet: This is a snippet\r\nThread ID: 67890\r\n--- Headers ---\r\nFrom: test@example.com\r\nTo: recipient@example.com\r\n--- Plain Text ---\r\nHello, this is a test email!\r\n"
	assert.Equal(t, expected, string(txtData))
}
