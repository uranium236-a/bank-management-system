# Bank Management System

A simple command-line Bank Management System built with Go. The application allows customers to create accounts, manage funds, track transactions, and transfer money. It also provides administrative functionality for monitoring bank accounts and balances.

## Features

### Customer Features

* Create a new bank account
* Login using email and password
* Deposit money
* Withdraw money
* Check account balance
* View transaction history
* Transfer money to another account
* Change account password
* Delete account

### Admin Features

* Create admin accounts using an access key
* Login as administrator
* View all customer accounts
* Search accounts
* View total bank balance
* Change admin password

### Data Persistence

* Bank data is automatically saved to a JSON file (`storage/bank.json`)
* Existing data is loaded when the application starts
* Account and transaction information persists between sessions

---

## Project Structure

```text
bank-management-system/
│
├── cli/                    # Command-line interface and menus
│   ├── auth.go
│   ├── home.go
│   └── reader.go
│
├── models/                 # Core data models
│   ├── account.go
│   ├── admin.go
│   └── bank.go
│
├── services/               # Business logic
│   └── bank_service.go
│
├── sessions/               # Session management
│   └── session.go
│
├── storage/
│   └── bank.json           # Persistent storage
│
├── main.go
├── go.mod
└── README.md
```

---

## Technologies Used

* Go (Golang)
* JSON file storage
* CLI (Command Line Interface)

---

## Installation

### Prerequisites

Make sure Go is installed:

```bash
go version
```

### Clone the Repository

```bash
git clone <repository-url>
cd bank-management-system
```

### Run the Application

```bash
go run .
```

Or:

```bash
go run main.go
```

---

## Usage

### Main Menu

```text
1. Signup
2. Login
3. Exit
```

### Create a Customer Account

1. Select **Signup**
2. Choose **New Account**
3. Enter:

   * Name
   * Email
   * Password
   * Initial Deposit (minimum ₹100)

### Create an Admin Account

1. Select **Signup**
2. Choose **New Admin**
3. Enter admin details
4. Provide the access key:

```text
ADMIN123
```

### Login

Login using your registered email and password.

---

## Account Operations

After logging in as a customer:

```text
1. Deposit
2. Withdraw
3. Check Balance
4. Transactions
5. Transfer Money
6. Change Password
7. Delete Account
8. Exit
```

### Deposits

Add money to your account balance.

### Withdrawals

Withdraw available funds from your account.

### Transfers

Transfer money directly to another account using the recipient's account number.

### Transactions

All deposits, withdrawals, and transfers are stored with timestamps.

---

## Admin Operations

After logging in as an administrator:

```text
1. Change Password
2. Display Accounts
3. Total Bank Balance
4. Search Account
5. Exit
```

Administrators can monitor customer accounts and overall bank balance.

---

## Data Storage

The application stores all bank information in:

```text
storage/bank.json
```

Data is automatically:

* Loaded on startup
* Saved on exit

This provides simple persistence without requiring a database.

---

## Sample Workflow

```text
Signup -> Create Account -> Login
      -> Deposit Money
      -> Transfer Money
      -> Check Transactions
      -> Exit

Application saves data to bank.json
```

---

## Future Improvements

* Password hashing and authentication security
* Database integration (MySQL/PostgreSQL)
* Account locking and security policies
* Interest calculation
* Loan management
* Role-based access control
* REST API support
* Unit and integration tests
* Transaction rollback support

---

## License

This project is open-source and available for educational and learning purposes.
