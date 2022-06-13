package pointer

import (
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

//  func (wallet Wallet) 修改为 func (wallet *Wallet) 你可以将其解读为「指向 wallet 的指针」。
func (wallet *Wallet) Deposit(monry Bitcoin) {
	fmt.Println("deposit address of balance in test is", &wallet.balance)
	wallet.balance += monry
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	// 存款10
	wallet.Deposit(10)

	got := wallet.Balance()
	want := Bitcoin(10)

	fmt.Println("mian address of balance in test is", &wallet.balance)

	if got != want { // 比较内存地址
		t.Errorf("got %v, want %v", got, want)
	}
}
