package main

import (
	"sync"
	"testing"
	"time"

	"testing/synctest"

	"github.com/stretchr/testify/assert"
)

func TestTodoMarkCompleted(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		todo := &Todo{Title: "Test Todo", Completed: false}
		todo.MarkComplete()
		time.Sleep(5 * time.Second)
		now := time.Now()
		assert.True(t, todo.Completed)
		assert.NotNil(t, todo.CompletedAt)
		assert.Equal(t, now, *todo.CompletedAt)
	})
}

func TestSendRequest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		time.Sleep(26 * 365 * 24 * time.Hour) // sleep 26 year
		synctest.Wait()
		resp, err := SendRequestWithContext(t.Context(), "https://example.com", "GET", nil)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
}

func TestUseBubbleChannelOutsideBubble(t *testing.T) {
	var wg *sync.WaitGroup
	synctest.Test(t, func(t *testing.T) {
		var inBubbleWg sync.WaitGroup
		inBubbleWg.Add(1)
		wg = &inBubbleWg
	})
	synctest.Test(t, func(t *testing.T) {
		var inBubbleWg sync.WaitGroup
		inBubbleWg.Add(1)
		wg = &inBubbleWg
	})
	wg.Add(1)
	t.Log("Both bubbles are done")
}
