package mockconfig

import (
	"github.com/josephsalimin/simple-ctftime-bot/internal/config"
	"github.com/stretchr/testify/mock"
)

type MockReader struct {
	mock.Mock
}

func (r *MockReader) Read() error {
	args := r.Called()

	return args.Error(0)
}

func (r *MockReader) Decode() (*config.Config, error) {
	args := r.Called()

	if c, ok := args.Get(0).(*config.Config); ok {
		return c, args.Error(1)
	}

	return nil, args.Error(1)
}
