package healthcheck

import (
	"context"
	"net/http"
	"time"

	"github.com/JesusIslam/healthcheck/check"
	"github.com/gin-gonic/gin"
)

const (
	DefaultAuthToken = "healthcheck"
	DefaultTimeout   = time.Second * 30

	AuthHeaderKey = "Authorization"

	SuccessFormat = "OK"
	ErrorFormat   = "ERROR: %v"

	InvalidTokenMessage = "invalid token"
)

type Healthcheck struct {
	Checker   check.Checker
	AuthToken string
	Timeout   time.Duration
}

func New(checker check.Checker, authToken string, timeout time.Duration) *Healthcheck {
	h := &Healthcheck{}
	h.Checker = &check.DefaultCheck{}
	h.AuthToken = DefaultAuthToken
	h.Timeout = DefaultTimeout

	if checker != nil {
		h.Checker = checker
	}
	if authToken != "" {
		h.AuthToken = authToken
	}
	if timeout > 0 {
		h.Timeout = timeout
	}

	return h
}

func (h *Healthcheck) Handle(c *gin.Context) {
	authToken := c.GetHeader(AuthHeaderKey)
	if authToken != h.AuthToken {
		c.String(http.StatusUnauthorized, ErrorFormat, InvalidTokenMessage)
		return
	}

	ctx, cancel := context.WithTimeout(context.TODO(), h.Timeout)
	defer cancel()

	err := h.Checker.Check(ctx)
	if err != nil {
		c.String(http.StatusServiceUnavailable, ErrorFormat, err)
		return
	}

	c.String(http.StatusOK, SuccessFormat)
}
