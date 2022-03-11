package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

var (
	support     SupportData
	dataSupport []SupportData
)

func SupportSystem() []SupportData {
	resp, err := http.Get("http://127.0.0.1:8383/support")

	if err != nil {
		errors.New(fmt.Sprintf("не удалось получить данные. ошибка: %v", err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.New(fmt.Sprintf("не удалось прочитать данные. ошибка: %v", err))
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		if err := json.Unmarshal(body, &dataSupport); err != nil {
			errors.New(fmt.Sprintf("не удалось прочитать данные. ошибка: %v", err))
		}

		fmt.Println("=================")
		fmt.Println("Состояние системы Support:")

		for i := 0; i < len(dataSupport); i++ {
			fmt.Println(dataSupport[i])
		}

	} else if resp.StatusCode == 500 {
		fmt.Println("Ошибка", dataSupport)
	}
	return dataSupport
}

func prepareSupport(dataSupport []SupportData) []int {
	dataSupportArray := make([]int, 0)
	supportLoad := 0

	for i := 0; i < len(dataSupport); i++ {
		supportLoad += dataSupport[i].ActiveTickets
	}
	supportTime := 60 / 18 * supportLoad

	if supportLoad < 9 {
		dataSupportArray = append(dataSupportArray, 1, supportTime)
	} else if supportLoad >= 9 && supportLoad < 16 {
		dataSupportArray = append(dataSupportArray, 2, supportTime)
	} else {
		dataSupportArray = append(dataSupportArray, 3, supportTime)
	}

	return dataSupportArray
}
