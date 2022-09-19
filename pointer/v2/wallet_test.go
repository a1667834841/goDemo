package pointer

import (
	"fmt"
	"testing"
)

// deposit address of balance in test is 0xc00009e278
// mian address of balance in test is 0xc00009e278
// 打印的两者内存地址相同，所以 got == want

type Wallet struct {
	balance int
}

// 当调用 func (wallet Wallet) Deposit(amount int) 时，wallet 是来自我们调用方法的副本。
// 运行结果
// address of balance in Deposit is 0xc420012268
// address of balance in test is 0xc420012260

//  func (wallet Wallet) 修改为 func (wallet *Wallet) 你可以将其解读为「指向 wallet 的指针」。
func (wallet *Wallet) Deposit(monry int) {
	fmt.Println("deposit address of balance in test is", &wallet.balance)
	wallet.balance += monry
}

func (wallet *Wallet) Balance() int {
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
