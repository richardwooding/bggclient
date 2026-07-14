package xml1

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestSearchCancelledBeforeRequest(t *testing.T) {
	var hits atomic.Int64
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
	}))
	defer srv.Close()

	api := NewAPI(Options{BaseURL: srv.URL, RequestInterval: time.Millisecond})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := api.SearchBoardgames(ctx, "Catan")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
	if n := hits.Load(); n != 0 {
		t.Errorf("server saw %d requests, want 0", n)
	}
}

func TestCollectionCancelledMidRetry(t *testing.T) {
	var hits atomic.Int64
	firstHit := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if hits.Add(1) == 1 {
			close(firstHit)
		}
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer srv.Close()

	api := NewAPI(Options{BaseURL: srv.URL, RequestInterval: 50 * time.Millisecond})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		<-firstHit
		cancel()
	}()

	start := time.Now()
	_, err := api.GetCollection(ctx, "someuser")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
	// Full retry budget would be 5 retries x 50ms; cancellation must cut
	// that short after the first response.
	if elapsed := time.Since(start); elapsed > time.Second {
		t.Errorf("took %v to return after cancellation", elapsed)
	}
	if n := hits.Load(); n > 2 {
		t.Errorf("server saw %d requests after cancellation, want <= 2", n)
	}
}
