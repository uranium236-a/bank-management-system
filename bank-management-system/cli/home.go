package cli

import (
	"bank-management-system/services"
	"bank-management-system/sessions"
	"fmt"
)

func homeAccount(service *services.BankService) {
	fmt.Println("\n-----------Account Home Screen-----------")

	fmt.Println("Account number: ", sessions.CurrentAccount.AccNo)
	fmt.Println("Account Holder: ", sessions.CurrentAccount.Name)

	fmt.Println("\n1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Check balance")
	fmt.Println("4. Transactions")
	fmt.Println("5. Transfer Money")
	fmt.Println("6. Change password")
	fmt.Println("7. Delete Account")
	fmt.Println("8. Exit")

	choice, _ := ReadInput(">> ")

	switch choice {
	case "1":
		fmt.Println("\n------------Deposit Amount-------------")

		depositStr, _ := ReadInput("Enter deposit money: ")

		err := service.DepositAmount(depositStr)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Amount deposited successfully")
		}

	case "2":
		fmt.Println("\n-------------Withdraw Amount--------------")

		withdrawStr, _ := ReadInput("Enter withdrawal money: ")

		err := service.WithdrawAmount(withdrawStr)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Amount Withdrawn successfully")
		}

	case "3":
		fmt.Println("\n---------Account Balance---------")
		fmt.Println("Account No: ", sessions.CurrentAccount.AccNo)
		fmt.Println("Name: ", sessions.CurrentAccount.Name)
		fmt.Println("Balance: ", sessions.CurrentAccount.Balance)

	case "4":
		fmt.Println("\n---------Transactions---------")

		for i := len(sessions.CurrentAccount.Transactions) - 1; i >= 0; i-- {
			fmt.Println(sessions.CurrentAccount.Transactions[i])
		}

	case "5":
		fmt.Println("\n------------Transfer Money-------------")
		amountstr, _ := ReadInput("Enter amount: ")
		accnostr, _ := ReadInput("Enter receiver Acc NO: ")

		err := service.TransferMoney(amountstr, accnostr)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Amount transferred successfully")
		}

	case "6":
		fmt.Println("\n------------Change Password--------------")
		oldpass, _ := ReadInput("Enter old password: ")

		if oldpass == sessions.CurrentAccount.Password {
			newpass, _ := ReadInput("Enter new password: ")
			sessions.CurrentAccount.Password = newpass
			fmt.Println("Password changed successfully")
		} else {
			fmt.Println("Invalid password, contact bank on this number: xxxxxx")
		}

	case "7":
		fmt.Println("\n------------Delete Account--------------")

		pass, _ := ReadInput("Enter password: ")

		if pass == sessions.CurrentAccount.Password {
			service.DeleteAccount()
			fmt.Println("Account Deleted successfully")
			sessions.CurrentAccount = nil
		} else {
			fmt.Println("Incorrect Password")
		}

	case "8":
		sessions.CurrentAccount = nil
	}
}

func homeAdmin(service *services.BankService) {
	fmt.Println("\n-----------Admin Home Screen-----------")

	fmt.Println("Admin ID: ", sessions.CurrentAdmin.AdminID)
	fmt.Println("Admin: ", sessions.CurrentAdmin.Name)

	fmt.Println("\n1. Change password")
	fmt.Println("2. Display Accounts")
	fmt.Println("3. Total Bank Balance")
	fmt.Println("4. Search Account")
	fmt.Println("5. Exit")

	choice, _ := ReadInput(">> ")

	switch choice {
	case "1":
		fmt.Println("\n------------Change Password--------------")
		oldpass, _ := ReadInput("Enter old password: ")

		if oldpass == sessions.CurrentAdmin.Password {
			newpass, _ := ReadInput("Enter new password: ")
			sessions.CurrentAdmin.Password = newpass
			fmt.Println("Password changed successfully")
		} else {
			fmt.Println("Invalid password")
		}

	case "2":
		fmt.Println("\n-------------All Accounts-------------")

		fmt.Println("Total Accounts: ", service.GetAccountsCount())

		for k, v := range service.GetAllAccounts() {
			fmt.Println(k, " - ", v.Name)
		}

	case "3":
		fmt.Println("\n------------Total Bank Balance------------")

		fmt.Println("\nTotal Money in Bank: ", service.TotalBankBalance())

	case "4":
		fmt.Println("\n-------------Search Account---------------")
		fmt.Println("\n1. By Acc No")
		fmt.Println("2. By Name")
		choice, _ := ReadInput(">> ")

		switch choice {
		case "1":
			accno, _ := ReadInput("Enter Acc No: ")
			account, err := service.GetAccountByAccNo(accno)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Acc No: ", account.AccNo)
				fmt.Println("Name: ", account.Name)
				fmt.Println("Email: ", account.Email)
			}

		case "2":
			name, _ := ReadInput("Enter Name: ")

			accounts, err := service.GetAccountByName(name)

			if err != nil {
				fmt.Println(err)
			} else {
				for _, account := range accounts {
					fmt.Println("Acc No: ", account.AccNo)
					fmt.Println("Name: ", account.Name)
					fmt.Println("Email: ", account.Email)
				}
			}
		}
	case "5":
		sessions.CurrentAdmin = nil
	}
}
