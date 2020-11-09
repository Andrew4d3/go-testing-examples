package util

import (
	"math/rand"

	"errors"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"

	"github.com/Andrew4d3/go-testing-examples/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_GetBalanceByID(t *testing.T) {
	defer monkey.UnpatchAll()

	t.Run("Should return the correct Balance amount", func(t *testing.T) {
		bankConn := new(bankConnection)
		monkey.PatchInstanceMethod(reflect.TypeOf(&rand.Rand{}), "Float64", func(_ *rand.Rand) float64 {
			return 0.5
		})

		balance, err := bankConn.GetBalanceByID(1)

		assert.NoError(t, err)
		assert.Equal(t, float64(500), balance)
	})
}

func Test_GetCurrentISOTime(t *testing.T) {
	defer monkey.UnpatchAll()

	t.Run("Should return the correct ISO time", func(t *testing.T) {
		monkey.Patch(time.Now, func() time.Time {
			return time.Date(2020, 10, 25, 0, 0, 0, 0, time.UTC)
		})

		assert.Equal(t, "2020-10-25T00:00:00Z", GetCurrentISOTime())
	})
}

func Test_ExtractContextValue(t *testing.T) {
	t.Run("Should return the expected value from the context", func(t *testing.T) {
		ctx := new(mocks.MockedContext)
		ctx.On("Value", "foo").Return("bar")

		value := ExtractContextValue(ctx, "foo")

		assert.Equal(t, value, "bar")
		ctx.AssertExpectations(t)
	})
}

func Test_SumAccountBalances(t *testing.T) {
	t.Run("Should return the correct sum of both accounts", func(t *testing.T) {
		// Using the generated mocks
		bankConn := new(mocks.BankConnection)
		bankConn.On("GetBalanceByID", 1).Return(float64(1000), nil)
		bankConn.On("GetBalanceByID", 2).Return(float64(2000), nil)

		total, err := SumAccountBalances(1, 2, bankConn)

		assert.NoError(t, err)
		assert.Equal(t, total, float64(3000))

		bankConn.AssertExpectations(t)
	})

	t.Run("Should return error if the first account gets Error", func(t *testing.T) {
		// Using the generated mocks
		bankConn := new(mocks.BankConnection)
		bankConn.On("GetBalanceByID", 1).Return(float64(0), errors.New("Boom A"))

		_, err := SumAccountBalances(1, 2, bankConn)

		assert.Errorf(t, err, "Boom A")
		bankConn.AssertExpectations(t)
	})

	t.Run("Should return error if the first account gets Error", func(t *testing.T) {
		// Using the generated mocks
		bankConn := new(mocks.BankConnection)
		bankConn.On("GetBalanceByID", 1).Return(float64(1000), nil)
		bankConn.On("GetBalanceByID", 2).Return(float64(0), errors.New("Boom B"))

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
