package models

type Bank struct {
	Name          string
	Accounts      map[int]*Account
	Admins        map[int]*Admin
	Balance       float32
	NextAccountID int
	NextAdminID   int
}

func NewBank(name string) *Bank {
	return &Bank{
		Name:          name,
		Accounts:      make(map[int]*Account),
		Admins:        make(map[int]*Admin),
		Balance:       100000,
		NextAccountID: 0,
		NextAdminID:   0,
	}
}

func (b *Bank) AddAccount(name, email, password string, initialDeposit float32) *Account {

	id := b.NextAccountID

	b.Accounts[id] = &Account{
		AccNo:        id,
		Name:         name,
		Email:        email,
		Password:     password,
		Balance:      initialDeposit,
		Transactions: []string{},
	}

	b.NextAccountID++

	return b.Accounts[id]
}

func (b *Bank) AddAdmin(name, email, password string) *Admin {
	id := b.NextAdminID

	b.Admins[id] = &Admin{
		AdminID:  id,
		Name:     name,
		Email:    email,
		Password: password,
	}

	b.NextAdminID++

	return b.Admins[id]
}

func (b *Bank) AccountExists(email string) bool {
	for _, v := range b.Accounts {
		if v.Email == email {
			return true
		}
	}

	return false
}

func (b *Bank) AdminExists(email string) bool {
	for _, v := range b.Admins {
		if v.Email == email {
			return true
		}
	}
	return false
}
