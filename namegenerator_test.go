package codenamegenerator

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockReaderCloser struct {
	mock.Mock
}

func (m *mockReaderCloser) Read(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *mockReaderCloser) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestNameGenerate(t *testing.T) {
	b := []byte("test\nblock\ntoronto")
	reader := new(mockReaderCloser)
	reader.On("Read", b).Return(18, nil)
	reader.On("Close").Return(nil)

	n := NameGenerate(reader)
	if n != "test-block-toronto" {
		t.Errorf("Expected %b but got %s", b, n)
	}

	reader.AssertExpectations(t)
}
