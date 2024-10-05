package account

import "sync"

type Account struct {
	balance int64
	mu      sync.Mutex
	open    bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{
		balance: amount,
		open:    true,
	}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.open == false {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.open == false {
		return 0, false
	}

	if a.balance < -amount {
		return 0, false
	}

	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.open {
		return 0, false
	}

	a.open = false
	payout := a.balance
	a.balance = 0
	return payout, true
}
