package linkmasker

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) produce() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

type MockPresenter struct {
	mock.Mock
}

func (m *MockPresenter) present(messages []string) error {
	args := m.Called(messages)
	return args.Error(0)
}

func TestHideLinks(t *testing.T) {
	service := &LinkMasker{}

	tests := []struct {
		name     string
		input  string
		expected string
	}{
		{
			name:     "No links",
			input:    "",
			expected: "",
		},
		{
			name:     "Single link",
			input:    "Here's my spammy page: http://hehefouls.netHAHAHA see you.",
			expected: "Here's my spammy page: http://******************* see you.",
		},
		{
			name:     "Multiple links",
			input:    "http://hehefouls.netHAHAHA http://hehefouls.netHAHAHA",
			expected: "http://******************* http://*******************",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.hideLinks(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name             string
		produceResult    []string
		produceError     error
		presentMessages  []string
		presentError     error
		expectedError    string
	}{
		{
			name:            "Success",
			produceResult:   []string{"This is a link http://example.com"},
			produceError:    nil,
			presentMessages: []string{"This is a link http://***********"},
			presentError:    nil,
			expectedError:   "",
		},
		{
			name:          "Producer Error",
			produceResult: nil,
			produceError:  errors.New("failed to produce"),
			expectedError: "failed to produce",
		},
		{
			name:            "Presenter Error",
			produceResult:   []string{"This is a link http://example.com"},
			produceError:    nil,
			presentMessages: []string{"This is a link http://***********"},
			presentError:    errors.New("failed to present"),
			expectedError:   "failed to present",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockProducer := new(MockProducer)
			mockPresenter := new(MockPresenter)

			service := NewService(mockProducer, mockPresenter)

			mockProducer.On("produce").Return(test.produceResult, test.produceError)
			if test.produceError == nil {
				mockPresenter.On("present", test.presentMessages).Return(test.presentError)
			}

			err := service.Run()

			if test.expectedError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.expectedError)
			}

			mockProducer.AssertExpectations(t)
			mockPresenter.AssertExpectations(t)
		})
	}
}