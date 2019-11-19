package check

import (
	"context"
)

type Checker interface {
	Check(context.Context) error
}

type DefaultCheck struct{}

func (d *DefaultCheck) Check(ctx context.Context) error {
	return nil
}
