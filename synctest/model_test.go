package main

import (
	"fmt"
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
		fmt.Println("now:", now)
		assert.True(t, todo.Completed)
		assert.NotNil(t, todo.CompletedAt)
		assert.Equal(t, now, todo.CompletedAt.Add(5*time.Second))
	})

	synctest.Test(t, func(t *testing.T) {
		var wg sync.WaitGroup
		for i := 1; i <= 100; i++ {
			wg.Go(func() {
				time.Sleep(time.Duration(i) * time.Second)
				t.Log("Inner bubble done")
			})
		}
		synctest.Wait()
		wg.Wait()
		t.Log("Outer bubble done")
	})

}

func TestSendRequest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
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
