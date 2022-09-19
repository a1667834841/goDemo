package pointer

import (
	"errors"
	"fmt"
	"testing"
)

// Go 允许从现有的类型创建新的类型。
type Bitcoin int

func (b *Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (b *Bitcoin) Int() int {
	return int(*b)
}

type Wallet struct {
	balance Bitcoin
}

// 转string 工具接口
type Stringer interface {
	String() string
}

type Integer interface {
	Int() int
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

//  func (wallet Wallet) 修改为 func (wallet *Wallet) 你可以将其解读为「指向 wallet 的指针」。
func (wallet *Wallet) Deposit(monry Bitcoin) {
	fmt.Println("deposit address of balance in test is", &wallet.balance)
	wallet.balance += monry
}

// 取钱
func (wallet *Wallet) Withdraw(money Bitcoin) error {
	if money > wallet.balance {
		return InsufficientFundsError
	}
	wallet.balance -= money
	return nil
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		// 存款10
		wallet.Deposit(10)

		want := Bitcoin(10)

		fmt.Println("mian address of balance in test is", &wallet.balance)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(10)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		got := wallet.Withdraw(200)

		assertBalance(t, wallet, Bitcoin(20))

		assertError(t, got, InsufficientFundsError)

	})

}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
