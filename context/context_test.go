package context

import (
	"testing"
	"time"
	"net/http"
	"net/http/httptest"
	"context"
	// "errors"
)

type SpyStore struct {
	response string
	cancelled bool
	t *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}

func TestServer(t *testing.T) {

	data := "hello, world"

	t.Run("returns data from store", func(t *testing.T) {
		
		store := &SpyStore{response: data, t : t}
		svr := Server(store)
		
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder() // 返回一个初始化了的ResponseRecorder
		
		svr.ServeHTTP(response, request)
		
		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})

	
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t : t}
		svr := Server(store)
		
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5 * time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)
		
		response := httptest.NewRecorder() // 返回一个初始化了的ResponseRecorder
		
		svr.ServeHTTP(response, request)
		
		store.assertWasCancelled()
	})
}