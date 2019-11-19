package check

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultCheck(t *testing.T) {
	t.Run("Check", func(t *testing.T) {
		d := &DefaultCheck{}
		require.Implements(t, new(Checker), d)

		require.NoError(t, d.Check(context.TODO()))
	})
}
