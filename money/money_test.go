package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertError(t testing.TB, actual, expected error) {
	t.Helper()
	if actual == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if actual.Error() != expected.Error() {
		t.Errorf("got %q, want %q", actual, expected)
	}
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	actual := wallet.Balance()
	assert.Equalf(t, expected, actual, "got %s but expected %s", actual.String(), expected.String())
}

func TestWallet_Deposit(t *testing.T) {
	moneyTests := []struct {
		name      string
		wallet    Wallet
		toDeposit Bitcoin
		expected  Bitcoin
	}{
		{"Deposit 10 coins", Wallet{}, 10, 10},
		{"Deposit 1000 coins", Wallet{}, 1000, 1000},
	}

	for _, tt := range moneyTests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wallet.Deposit(tt.toDeposit)
			assertBalance(t, tt.wallet, tt.expected)
		})
	}
}

func TestWallet_Withdraw(t *testing.T) {
	moneyTests := []struct {
		name     string
		wallet   Wallet
		balance  Bitcoin
		withdraw Bitcoin
		expected Bitcoin
		wantErr  bool
	}{
		{"Withdraw 10 coins", Wallet{}, 10, 10, 0, false},
		{"Withdraw 1000 coins", Wallet{}, 1010, 1000, 10, false},
		{"Withdraw insufficient funds", Wallet{}, 500, 1000, 0, true},
	}

	for _, tt := range moneyTests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wallet.Deposit(tt.balance)
			err := tt.wallet.Withdraw(tt.withdraw)
			if tt.wantErr {
				assertError(t, err, ErrInsufficientFunds)
				return
			}
			assertBalance(t, tt.wallet, tt.expected)
		})
	}
}
