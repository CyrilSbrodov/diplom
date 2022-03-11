package internal

import (
	data2 "diplom/data"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

var (
	email        EmailData
	dataEmail    []EmailData
	dataEmailMap map[string][][]EmailData
)

func EmailSystem() []EmailData {
	fileName := "simulator/email.data"
	file, err := os.Open(fileName)
	if err != nil {
		errors.New(fmt.Sprintf("не удалось открыть файл. ошибка: %v", err))
	}
	defer file.Close()
	result, err := ioutil.ReadAll(file)
	if err != nil {
		errors.New(fmt.Sprintf("не удалось прочитать файл. ошибка: %v", err))
	}
	resultString := strings.Split(string(result), "\n")

	for i := 0; i < len(resultString); i++ {
		res := strings.Split(resultString[i], ";")
		if len(res) != 3 {
			continue
		} else {
			if _, exist := data2.Country[res[0]]; !exist {
				continue
			} else {
				email.Country = res[0]
			}
			if _, exist := data2.EmailProviders[res[1]]; !exist {
				continue
			} else {
				email.Provider = res[1]
			}
			email.DeliveryTime = stringToInt(res[2])
			dataEmail = append(dataEmail, email)
		}
	}
	fmt.Println("=================")
	fmt.Println("Состояние системы Email:")

	for i := 0; i < len(dataEmail); i++ {
		fmt.Println(dataEmail[i])
	}
	return dataEmail
}

func prepareEmail(email []EmailData) map[string][][]EmailData {
	result := make(map[string][][]EmailData, 0)

	countries := make(map[string]int)
	for _, elem := range email {
		countries[elem.Country]++
	}

	for countryCode, _ := range countries {
		var emailDataItem [][]EmailData
		emailDataItem = append(emailDataItem, Get3Min(email, countryCode))
		emailDataItem = append(emailDataItem, Get3Max(email, countryCode))

		result[countryCode] = emailDataItem
	}

	return result
}

func Get3Min(data []EmailData, code string) []EmailData {
	result := make([]EmailData, 0)
	for _, elem := range data {
		if elem.Country == code {
			result = append(result, elem)
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].DeliveryTime < result[j].DeliveryTime
	})

	if len(result) < 3 {
		return result
	}

	return result[:3]
}

func Get3Max(data []EmailData, code string) []EmailData {
	result := make([]EmailData, 0)
	for _, elem := range data {
		if elem.Country == code {
			result = append(result, elem)
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].DeliveryTime > result[j].DeliveryTime
	})

	if len(result) < 3 {
		return result
	}

	return result[:3]
}
