package user

import "sync"

type User struct {
	ID              string `json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Wallet          Wallet `json:"wallet"`
	QKDKey          string `json:"qkd_key"`
	MessageSequence MessageSequence
}
type MessageSequence struct {
	SendSequence    chan string
	ReceiveSequence chan string
	Mutex           sync.Mutex
}
type Wallet struct {
	ID         string `json:"id"`
	UId        string `json:"u_id"`
	Amount     int    `json:"amount"`
	Address    string `json:"address"`
	WalletType int    `json:"wallet_type"` //0: 本地钱包 1: 远程钱包
}

func (u *User) GetWallet() Wallet {
	return u.Wallet
}

func (u *User) GetUsernameAndPassword() map[string]string {
	return map[string]string{u.Username: u.Password}
}
func (u *User) GetWalletAmount() int {
	return u.Wallet.Amount
}
func (u *User) GetUId() string {
	return u.Wallet.UId
}
