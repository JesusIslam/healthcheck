package healthcheck

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type dummyCheck struct {
	err bool
}

func (d *dummyCheck) Check(ctx context.Context) (err error) {
	if d.err {
		err = errors.New("dummy")
	}
	return
}

func TestHealthcheck(t *testing.T) {
	var h *Healthcheck
	var wg sync.WaitGroup

	wg.Add(1)
	t.Run("New default", func(t *testing.T) {
		defer wg.Done()

		h = New(nil, "", 0)
		require.NotNil(t, h)
	})
	wg.Wait()

	wg.Add(1)
	t.Run("New complete", func(t *testing.T) {
		defer wg.Done()

		h = New(&dummyCheck{}, "something", time.Second)
		require.NotNil(t, h)
	})
	wg.Wait()

	gin.SetMode(gin.TestMode)
	t.Run("Handle no header", func(t *testing.T) {
		rw := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rw)
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Fatalf("Failed to create new request: %v", err)
		}

		c.Request = req
		h.Handle(c)

		resp := rw.Result()
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response data: %v", err)
		}

		require.Equal(t, "ERROR: "+InvalidTokenMessage, string(data))
	})

	t.Run("Handle no error", func(t *testing.T) {
		rw := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rw)
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Fatalf("Failed to create new request: %v", err)
		}
		req.Header.Add("Authorization", "something")

		c.Request = req
		h.Handle(c)

		resp := rw.Result()
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response data: %v", err)
		}

		require.Equal(t, "OK", string(data))
	})

	t.Run("Handle with error", func(t *testing.T) {
		h.Checker = &dummyCheck{err: true}

		rw := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rw)
		req, err := http.NewRequest(http.MethodGet, "", nil)
		if err != nil {
			t.Fatalf("Failed to create new request: %v", err)
		}
		req.Header.Add("Authorization", "something")

		c.Request = req
		h.Handle(c)

		resp := rw.Result()
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response data: %v", err)
		}

		require.Equal(t, "ERROR: dummy", string(data))
	})
}
