package pointer

import (
	"fmt"
	"testing"
)

//deposit address of balance in test is 0xc0000162d0
//mian address of balance in test is 0xc0000162c8
// 打印申明的Wallet结构体和调用的Wallet结构体，发现两者内存地址不一致，调用的Wallet是申明的Wallet的副本
// 所以 got != want 这时候就需要使用到指针，将调用者Wallet变量，指向申明的Wallet的内存地址

type Wallet struct {
	balance int
}

// 当调用 func (wallet Wallet) Deposit(amount int) 时，wallet 是来自我们调用方法的副本。
func (wallet Wallet) Deposit(monry int) {
	fmt.Println("deposit address of balance in test is", &wallet.balance)
	wallet.balance += monry
}

func (wallet Wallet) Balance() int {
	return wallet.balance
}

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	// 存款10
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	fmt.Println("mian address of balance in test is", &wallet.balance)

	if got != want { // 比较内存地址
		t.Errorf("got %v, want %v", got, want)
	}
}
