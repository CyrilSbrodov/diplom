package internal

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

var (
	billing BillingData
)

func BillingSystem() BillingData {
	fileName := "simulator/billing.data"
	file, err := os.Open(fileName)
	if err != nil {
		errors.New(fmt.Sprintf("не удалось открыть файл. ошибка: %v", err))
	}
	defer file.Close()
	result, err := ioutil.ReadAll(file)
	if err != nil {
		errors.New(fmt.Sprintf("не удалось прочитать файл. ошибка: %v", err))
	}

	var j float64
	var res uint8

	for i := len(result) - 1; i >= 0; i-- {
		if result[i] == 49 {
			res += uint8(math.Pow(2, j))
		} else {
			continue
		}
		j++
	}

	billing.CreateCustomer = result[0] == 49
	billing.Purchase = result[1] == 49
	billing.Payout = result[2] == 49
	billing.Recurring = result[3] == 49
	billing.FraudControl = result[4] == 49
	billing.CheckoutPage = result[5] == 49

	fmt.Println("=================")
	fmt.Println("Состояние системы Billing:")
	fmt.Println("billing в десятичном формате:")
	fmt.Println(res)
	fmt.Println("=================")
	fmt.Println("Состояние системы Billing:")
	fmt.Println("Create Customer:", billing.CreateCustomer)
	fmt.Println("Purchase:", billing.Purchase)
	fmt.Println("Payout:", billing.Payout)
	fmt.Println("Recurring:", billing.Recurring)
	fmt.Println("Fraud Control:", billing.FraudControl)
	fmt.Println("Checkout Page:", billing.CheckoutPage)

	return billing
}
