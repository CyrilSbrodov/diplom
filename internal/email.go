package internal

import (
	data2 "diplom/data"
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
	fileName := "simulator/skillbox-diploma/email.data"
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	result, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
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

	data := make(map[string][][]EmailData)

	sort.Slice(email, func(i, j int) bool {
		return email[i].Country < email[j].Country
	})

	for i := 0; i < len(email); i++ {
		c := make([][]EmailData, 2)
		for j := 0; j < len(email); j++ {
			if email[i].Country == email[j].Country {
				c[0] = append(c[0], email[j])
				c[1] = append(c[1], email[j])
				data[email[i].Country] = c
			}
		}
	}

	for i := 0; i < len(email); i++ {
		if email[i].DeliveryTime > data[email[i].Country][0][0].DeliveryTime {
			data[email[i].Country][0][2] = data[email[i].Country][0][1]
			data[email[i].Country][0][1] = data[email[i].Country][0][0]
			data[email[i].Country][0][0] = email[i]
		} else if email[i].DeliveryTime > data[email[i].Country][0][1].DeliveryTime && email[i].DeliveryTime < data[email[i].Country][0][0].DeliveryTime {
			data[email[i].Country][0][2] = data[email[i].Country][0][1]
			data[email[i].Country][0][1] = email[i]
		} else if email[i].DeliveryTime > data[email[i].Country][0][2].DeliveryTime && email[i].DeliveryTime < data[email[i].Country][0][1].DeliveryTime {
			data[email[i].Country][0][2] = email[i]
		}
	}

	for i := 0; i < len(email); i++ {
		if email[i].DeliveryTime < data[email[i].Country][1][0].DeliveryTime {
			data[email[i].Country][1][2] = data[email[i].Country][1][1]
			data[email[i].Country][1][1] = data[email[i].Country][1][0]
			data[email[i].Country][1][0] = email[i]
		} else if email[i].DeliveryTime < data[email[i].Country][1][1].DeliveryTime && email[i].DeliveryTime > data[email[i].Country][1][0].DeliveryTime {
			data[email[i].Country][1][2] = data[email[i].Country][1][1]
			data[email[i].Country][1][1] = email[i]
		} else if email[i].DeliveryTime < data[email[i].Country][1][2].DeliveryTime && email[i].DeliveryTime > data[email[i].Country][1][1].DeliveryTime {
			data[email[i].Country][1][2] = email[i]
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < len(email); j++ {
			if data[email[j].Country][0][i] == data[email[j].Country][0][i] {
				data[email[j].Country][0] = data[email[j].Country][0][0:3]
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < len(email); j++ {
			if data[email[j].Country][1][i] == data[email[j].Country][1][i] {
				data[email[j].Country][1] = data[email[j].Country][1][0:3]
			}
		}
	}

	return data
}
