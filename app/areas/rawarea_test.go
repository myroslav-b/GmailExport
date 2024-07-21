package areas

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/gmail/v1"
)

func TestPrepareRawArea(t *testing.T) {
	// Define a sample Gmail message
	message := &gmail.Message{
		Id:           "12345",
		InternalDate: 1620000000000,
		LabelIds:     []string{"INBOX", "IMPORTANT"},
		SizeEstimate: 2048,
		Snippet:      "This is a snippet",
		ThreadId:     "67890",
		Raw:          base64.URLEncoding.EncodeToString([]byte("Raw email content")),
	}

	// Call the function
	result, err := PrepareRawArea(message)

	// Check no error occurred
	require.NoError(t, err)

	// Verify the result
	assert.Equal(t, "12345", result.Id)
	assert.Equal(t, int64(1620000000000), result.InternalDate)
	assert.ElementsMatch(t, []string{"INBOX", "IMPORTANT"}, result.LabelIds)
	assert.Equal(t, int64(2048), result.SizeEstimate)
	assert.Equal(t, "This is a snippet", result.Snippet)
	assert.Equal(t, "67890", result.ThreadId)
	assert.Equal(t, "Raw email content", result.Raw)
}

func TestTMessageRawArea_String(t *testing.T) {
	message := TMessageRawArea{
		Id:           "12345",
		InternalDate: 1620000000000,
		LabelIds:     []string{"INBOX", "IMPORTANT"},
		SizeEstimate: 2048,
		Snippet:      "This is a snippet",
		ThreadId:     "67890",
		Raw:          "Raw email content",
	}

	expected := "ID: 12345\r\nInternal Date: 1620000000000\r\nLabel IDs: INBOX, IMPORTANT, \r\nSize Estimate: 2048\r\nSnippet: This is a snippet\r\nThread ID: 67890\r\n--- Raw Body ---:\r\nRaw email content\r\n"
	assert.Equal(t, expected, message.String())
}

func TestTMessageRawArea_ToJson(t *testing.T) {
	message := TMessageRawArea{
		Id:           "12345",
		InternalDate: 1620000000000,
		LabelIds:     []string{"INBOX", "IMPORTANT"},
		SizeEstimate: 2048,
		Snippet:      "This is a snippet",
		ThreadId:     "67890",
		Raw:          "Raw email content",
	}

	jsonData, err := message.ToJson()
	require.NoError(t, err)

	expected := `{"id":"12345","internalDate":"1620000000000","labelIds":["INBOX","IMPORTANT"],"sizeEstimate":2048,"snippet":"This is a snippet","threadId":"67890","raw":"Raw email content"}`
	assert.JSONEq(t, expected, string(jsonData))
}

func TestTMessageRawArea_ToTxt(t *testing.T) {
	message := TMessageRawArea{
		Id:           "12345",
		InternalDate: 1620000000000,
		LabelIds:     []string{"INBOX", "IMPORTANT"},
		SizeEstimate: 2048,
		Snippet:      "This is a snippet",
		ThreadId:     "67890",
		Raw:          "Raw email content",
	}

	txtData, err := message.ToTxt()
	require.NoError(t, err)

	expected := "ID: 12345\r\nInternal Date: 1620000000000\r\nLabel IDs: INBOX, IMPORTANT, \r\nSize Estimate: 2048\r\nSnippet: This is a snippet\r\nThread ID: 67890\r\n--- Raw Body ---:\r\nRaw email content\r\n"
	assert.Equal(t, expected, string(txtData))
}
