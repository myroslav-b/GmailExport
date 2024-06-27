package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/gmail/v1"
)

// Test newListMessages function
func TestNewListMessages(t *testing.T) {
	lm := newListMessages()
	assert.NotNil(t, lm)
	assert.Empty(t, lm.messages)
	assert.Equal(t, int64(0), lm.resultSizeEstimate)
}

// Test addList function
func TestAddList(t *testing.T) {
	lm := newListMessages()
	msgs := []*gmail.Message{{Id: "123"}, {Id: "456"}}
	lm.addList(msgs, 2)
	assert.Len(t, lm.messages, 2)
	assert.Equal(t, "123", lm.messages[0].Id)
	assert.Equal(t, "456", lm.messages[1].Id)
	assert.Equal(t, int64(2), lm.resultSizeEstimate)
}

// To test the search function, you need to create mocks for the Gmail API service and the filter.
// This can be done using a mocking library such as github.com/golang/mock.
/*
func TestSearch(t *testing.T) {
	// Mock the Gmail service and user for testing purposes
	srv := &gmail.Service{}
	user := "me"
	filter := tFilter{Subject: "test"}

	// Mock the response from the Gmail API
	mockMessages := []*gmail.Message{{Id: "123"}, {Id: "456"}}
	mockListMessages := &tListMessages{messages: mockMessages, resultSizeEstimate: 2}

	// Mock the search function to return the mock response
	search := func(srv *gmail.Service, user string, filter tFilter) (*tListMessages, error) {
		return mockListMessages, nil
	}

	// Perform the test
	lm, err := search(srv, user, filter)
	assert.NoError(t, err)
	assert.Equal(t, mockListMessages, lm)
}
*/
