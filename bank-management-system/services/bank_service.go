package services

import (
	"bank-management-system/models"
	"bank-management-system/sessions"
	"errors"
	"strconv"
	"time"
)

type BankService struct {
	bank *models.Bank
}

func NewBankService(bank *models.Bank) *BankService {
	return &BankService{
		bank: bank,
	}
}

func (s *BankService) BankName() string {
	return s.bank.Name
}

func (s *BankService) CreateAccount(name, email, password string, initialDeposit float32) (*models.Account, error) {

	if s.bank.AccountExists(email) {
		return nil, errors.New("Account already exists")
	}

	if initialDeposit < 100 {
		return nil, errors.New("Minimum deposit should be 100")
	}

	return s.bank.AddAccount(name, email, password, initialDeposit), nil
}

func (s *BankService) CreateAdmin(name, email, password, accessKey string) (*models.Admin, error) {
	if s.bank.AdminExists(email) {
		return nil, errors.New("Admin already exists")
	}

	if accessKey != sessions.AccessKey {
		return nil, errors.New("Invalid Access Key")
	}

	return s.bank.AddAdmin(name, email, password), nil
}

func (s *BankService) AssignAccount(id int) {
	sessions.CurrentAccount = s.bank.Accounts[id]
}

func (s *BankService) AssignAdmin(id int) {
	sessions.CurrentAdmin = s.bank.Admins[id]
}

func (s *BankService) GetAccountByEmail(email string) (*models.Account, error) {
	for _, v := range s.bank.Accounts {
		if v.Email == email {
			return v, nil
		}
	}

	return nil, errors.New("Account doesn't exists")
}

func (s *BankService) GetAccountByAccNo(accnostr string) (*models.Account, error) {
	accno, err := strconv.Atoi(accnostr)

	if err != nil {
		return nil, errors.New("Invalid Account No.")
	}

	if _, exist := s.bank.Accounts[accno]; exist {
		return s.bank.Accounts[accno], nil
	} else {
		return nil, errors.New("Account doesn't exist")
	}
}

func (s *BankService) GetAccountByName(name string) ([]*models.Account, error) {

	accounts := make([]*models.Account, 0)

	for _, v := range s.bank.Accounts {
		if v.Name == name {
			accounts = append(accounts, v)
		}
	}

	if len(accounts) == 0 {
		return nil, errors.New("Account not found")
	}

	return accounts, nil

}

func (s *BankService) GetAdmin(email string) (*models.Admin, error) {
	for _, v := range s.bank.Admins {
		if v.Email == email {
			return v, nil
		}
	}

	return nil, errors.New("Admin doesn't exists")
}

func (s *BankService) DepositAmount(amountstr string) error {

	amount64, err := strconv.ParseFloat(amountstr, 32)

	if err != nil {
		return errors.New("Invalid amount format")
	}

	amount := float32(amount64)

	if amount < 0 {
		return errors.New("Amount must be greater than 0")
	}

	sessions.CurrentAccount.Balance += amount
	s.bank.Balance += amount

	transaction := dateTime() + " Deposited: " + amountstr
	sessions.CurrentAccount.Transactions = append(sessions.CurrentAccount.Transactions, transaction)

	return nil
}

func (s *BankService) WithdrawAmount(amountstr string) error {
	amount64, err := strconv.ParseFloat(amountstr, 32)

	if err != nil {
		return errors.New("Invalid amount format")
	}

	amount := float32(amount64)

	if amount > sessions.CurrentAccount.Balance {
		return errors.New("Amount exceeds account balance")
	}

	sessions.CurrentAccount.Balance -= amount
	s.bank.Balance -= amount

	transaction := dateTime() + " Withdrawn: " + amountstr

	sessions.CurrentAccount.Transactions = append(sessions.CurrentAccount.Transactions, transaction)

	return nil
}

func dateTime() string {
	now := time.Now()

	now = now.Truncate(time.Minute)

	dateTimeStr := now.Format("2006-01-02 15:04")

	return "[" + dateTimeStr + "]"
}

func (s *BankService) TransferMoney(amountstr, accnostr string) error {

	amount64, err := strconv.ParseFloat(amountstr, 32)

	if err != nil {
		return errors.New("Invalid amount format")
	}

	amount := float32(amount64)

	accno, err := strconv.Atoi(accnostr)

	if err != nil {
		return errors.New("Invalid account number")
	}

	if _, exists := s.bank.Accounts[accno]; !exists {
		return errors.New("Receiver account does not exist")
	}

	if sessions.CurrentAccount.AccNo == accno {
		return errors.New("Can't transfer money to oneself")
	}

	if amount <= 0 {
		return errors.New("Amount should be greater than 0")
	}

	if amount > sessions.CurrentAccount.Balance {
		return errors.New("Insufficient Balance")
	}

	s.bank.Accounts[accno].Balance += amount

	sessions.CurrentAccount.Balance -= amount

	senderAcc := strconv.FormatInt(int64(sessions.CurrentAccount.AccNo), 10)

	transaction := dateTime() + " Amount received: " + amountstr + " from Acc No: " + senderAcc
	s.bank.Accounts[accno].Transactions = append(s.bank.Accounts[accno].Transactions, transaction)

	transaction = dateTime() + " Amount transferred: " + amountstr + " to Acc No: " + accnostr
	sessions.CurrentAccount.Transactions = append(sessions.CurrentAccount.Transactions, transaction)

	return nil
}

func (s *BankService) DeleteAccount() {
	delete(s.bank.Accounts, sessions.CurrentAccount.AccNo)
}

func (s *BankService) GetAllAccounts() map[int]*models.Account {

	return s.bank.Accounts
}

func (s *BankService) GetAccountsCount() int {
	return len(s.bank.Accounts)
}

func (s *BankService) TotalBankBalance() float32 {
	return s.bank.Balance
}
