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
