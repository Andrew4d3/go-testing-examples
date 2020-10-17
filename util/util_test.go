package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedBankConn struct {
	mock.Mock
}

func (m *MockedBankConn) GetBalanceByID(id int) (float64, error) {
	args := m.Called(id)
	return (args.Get(0)).(float64), args.Error(1)
}

func Test_SumAccountBalances(t *testing.T) {
	t.Run("Should return the correct sum of both accounts", func(t *testing.T) {
		// Define our mocked connection object
		bankConn := new(MockedBankConn)
		// Setup expectation
		bankConn.On("GetBalanceByID", 1).Return(float64(1000), nil)
		bankConn.On("GetBalanceByID", 2).Return(float64(2000), nil)
		// Call the targeted function
		total, err := SumAccountBalances(1, 2, bankConn)

		assert.NoError(t, err)
		assert.Equal(t, total, float64(3000))

		bankConn.AssertExpectations(t)
	})

	t.Run("Should return error if the first account gets Error", func(t *testing.T) {
		// Define our mocked connection object
		bankConn := new(MockedBankConn)
		// Setup expectation
		bankConn.On("GetBalanceByID", 1).Return(float64(0), errors.New("Boom A"))
		// Call the targeted function
		_, err := SumAccountBalances(1, 2, bankConn)

		assert.Errorf(t, err, "Boom A")

		bankConn.AssertExpectations(t)
	})

	t.Run("Should return error if the first account gets Error", func(t *testing.T) {
		// Define our mocked connection object
		bankConn := new(MockedBankConn)
		// Setup expectation
		bankConn.On("GetBalanceByID", 1).Return(float64(1000), nil)
		bankConn.On("GetBalanceByID", 2).Return(float64(0), errors.New("Boom B"))
		// Call the targeted function
		_, err := SumAccountBalances(1, 2, bankConn)

		assert.Errorf(t, err, "Boom B")

		bankConn.AssertExpectations(t)
	})
}

// Traditional way
func Test_IsEven(t *testing.T) {

	tests := []struct {
		name   string
		number int
		want   bool
	}{
		{"Even number", 4, true},
		{"Odd number", 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.number); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Alternative way
func Test_IsEven_Alternative(t *testing.T) {
	t.Run("Should return true if an even number is entered", func(t *testing.T) {
		if got := IsEven(4); got != true {
			t.Errorf("Got = %v, expected %v", got, true)
		}
	})

	t.Run("Should return false if an odd number is entered", func(t *testing.T) {
		if got := IsEven(3); got != false {
			t.Errorf("Got = %v, expected %v", got, false)
		}
	})
}

func Test_IsEven_usingTestify(t *testing.T) {
	t.Run("Should return true if an even number is entered", func(t *testing.T) {
		assert.Equal(t, IsEven(4), true)
	})

	t.Run("Should return false if an odd number is entered", func(t *testing.T) {
		assert.Equal(t, IsEven(3), false)
	})
}
