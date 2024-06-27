package main

import (
	"time"

	"google.golang.org/api/gmail/v1"
)

// tListMessages represents a collection of Gmail messages and the estimated total number of results.
type tListMessages struct {
	// Messages: List of messages.
	messages []*gmail.Message
	// ResultSizeEstimate: Estimated total number of results.
	resultSizeEstimate int64
}

// newListMessages initializes a new instance of tListMessages with default values.
func newListMessages() *tListMessages {
	return &tListMessages{
		messages:           make([]*gmail.Message, 0),
		resultSizeEstimate: 0,
	}
}

// addList adds a list of messages and updates the result size estimate.
// listMsg: List of messages to add.
// size: The number to add to the result size estimate.
func (listMessages *tListMessages) addList(listMsg []*gmail.Message, size int64) {
	listMessages.messages = append(listMessages.messages, listMsg...)
	listMessages.resultSizeEstimate += size
}

// search retrieves messages from a user's Gmail account based on the provided filter.
// srv: The Gmail service instance used to make API calls.
// user: The email address (or me) of the user whose messages should be retrieved.
// filter: The filter criteria used to search for messages.
// Returns a tListMessages containing the retrieved messages and an error, if any.
func search(srv *gmail.Service, user string, filter tFilter) (*tListMessages, error) {
	listMessages := newListMessages()
	pageToken := ""
	startFlag := true

	for startFlag || pageToken != "" {
		// Retrieve a page of messages based on the filter and current page token.
		listMessagesResp, err := srv.Users.Messages.List(user).Q(filter.query()).PageToken(pageToken).Do()
		if err != nil {
			return nil, err
		}

		// Add the retrieved messages to the list and update the result size estimate.
		listMessages.addList(listMessagesResp.Messages, listMessagesResp.ResultSizeEstimate)

		// Update the page token for the next iteration.
		pageToken = listMessagesResp.NextPageToken
		startFlag = false
		// Introduce a short delay to avoid hitting rate limits.
		// There is probably a short delay between when a next_page_token is issued and when it will become valid.
		// https://stackoverflow.com/questions/41994292/trouble-looping-through-google-places-api-with-pagetoken-in-go
		time.Sleep(100 * time.Millisecond)
	}

	for i, m := range listMessages.messages {
		// "full" (default) - Returns the full email message data with body content
		// Parsed in the `payload` field; the `raw` field is not used
		message, err := srv.Users.Messages.Get(user, m.Id).Format("full").Do()
		if err != nil {
			return nil, err
		}
		// "raw" - Returns the full email message data with body content in the `raw`
		// field as a base64url encoded string; the `payload` field is not used
		message1, err := srv.Users.Messages.Get(user, m.Id).Format("raw").Do()
		message.Raw = message1.Raw
		if err != nil {
			return nil, err
		}
		listMessages.messages[i] = message
	}

	return listMessages, nil
}
