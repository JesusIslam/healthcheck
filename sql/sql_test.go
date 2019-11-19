package sql

import (
	"context"
	"sync"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/JesusIslam/healthcheck/check"
	"github.com/stretchr/testify/require"
)

func TestSQL(t *testing.T) {
	var wg sync.WaitGroup
	var s *SQLCheck
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sql mock: %v", err)
	}
	defer db.Close()

	wg.Add(1)
	t.Run("New", func(t *testing.T) {
		defer wg.Done()

		s = New(db)
		require.Implements(t, new(check.Checker), s)
	})
	wg.Wait()

	t.Run("Check", func(t *testing.T) {
		err = s.Check(context.TODO())

		require.NoError(t, mock.ExpectationsWereMet())
		require.NoError(t, err)
	})
}
