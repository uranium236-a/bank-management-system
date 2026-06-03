package cli

import (
	"bank-management-system/services"
	"bank-management-system/sessions"
	"fmt"
	"strconv"
)

func Start(service *services.BankService) {

	for {

		if sessions.CurrentAccount == nil && sessions.CurrentAdmin == nil {
			fmt.Println("\n---------------", service.BankName(), "---------------")

			fmt.Println("\n1. Signup")
			fmt.Println("2. Login")
			fmt.Println("3. Exit")

			choice, _ := ReadInput(">> ")

			if choice == "3" {
				break
			}

			switch choice {
			case "1":
				signup(service)
			case "2":
				login(service)
			}
		} else if sessions.CurrentAccount != nil {
			homeAccount(service)
		} else if sessions.CurrentAdmin != nil {
			homeAdmin(service)
		}

	}

}

func signup(service *services.BankService) {
	fmt.Println("\n--------------Signup--------------")

	fmt.Println("1. New Account")
	fmt.Println("2. New Admin")
	choice, _ := ReadInput(">> ")

	name, _ := ReadInput("Enter name: ")
	email, _ := ReadInput("Enter email: ")
	password, _ := ReadInput("Enter password: ")

	switch choice {
	case "1":
		depositStr, _ := ReadInput("Enter initial deposit(>= 100): ")
		initialDeposit, err := strconv.ParseFloat(depositStr, 32)
		if err != nil {
			fmt.Println("Invalid deposit")
			return
		}
		account, err := service.CreateAccount(name, email, password, float32(initialDeposit))

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Account created successfully: ", account.Name)

		sessions.CurrentAccount = account

		homeAccount(service)

	case "2":
		accessKey, _ := ReadInput("Enter access key: ")

		admin, err := service.CreateAdmin(name, email, password, accessKey)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Admin created succesfully: ", admin.Name)

		sessions.CurrentAdmin = admin
		homeAdmin(service)
	}

}

func login(service *services.BankService) {
	fmt.Println("---------------Login-----------------")

	fmt.Println("1. Account")
	fmt.Println("2. Admin")
	choice, _ := ReadInput(">> ")

	email, _ := ReadInput("Enter email: ")
	password, _ := ReadInput("Enter password: ")

	switch choice {
	case "1":
		account, err := service.GetAccountByEmail(email)

		if err != nil {
			fmt.Println(err)
			return
		}

		if password != account.Password {
			fmt.Println("Incorrect password")
			return
		}

		sessions.CurrentAccount = account
		homeAccount(service)

	case "2":
		admin, err := service.GetAdmin(email)

		if err != nil {
			fmt.Println(err)
			return
		}

		if password != admin.Password {
			fmt.Println("Incorrect password")
			return
		}

		sessions.CurrentAdmin = admin
		homeAdmin(service)
	}
}
