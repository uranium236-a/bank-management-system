package main

import (
	"bank-management-system/cli"
	"bank-management-system/models"
	"bank-management-system/services"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const filepath = "storage/bank.json"

func main() {

	var bank *models.Bank
	var err error

	bank, err = loadBankFromFile(filepath)

	if err != nil {
		log.Println("No saved data found, creating new Bank....")
		bank = models.NewBank("Justice Bank")
	} else {
		log.Println("Bank loaded successfully")
	}

	service := services.NewBankService(bank)

	cli.Start(service)

	if err := saveBankToFile(bank, filepath); err != nil {
		log.Println(err)
	} else {
		log.Println("Bank saved to file successfully")
	}
}

func saveBankToFile(bank *models.Bank, filepath string) error {

	data, err := json.MarshalIndent(bank, "", "   ")

	if err != nil {
		return fmt.Errorf("saving bank data failed: %v", err)
	}

	return os.WriteFile(filepath, data, 0666)
}

func loadBankFromFile(filepath string) (*models.Bank, error) {
	data, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	var bank models.Bank
	err = json.Unmarshal(data, &bank)

	if err != nil {
		return nil, err
	}

	return &bank, nil
}
