package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		if err := json.Unmarshal(body, &dataSupport); err != nil {
			log.Fatal(err)
		}

		//fmt.Println("=================")
		//fmt.Println("Состояние системы Support:")
		//
		//for i := 0; i < len(dataSupport); i++ {
		//	fmt.Println(dataSupport[i])
		//}

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

	//b := make([][]int, len(dataSupport))
	//for i := 0; i < len(dataSupport); i++ {
	//	b[i] = make([]int, 2)
	//	for j := 0; j < 2; j++ {
	//		if j == 0 {
	//			if dataSupport[i].ActiveTickets < 9 {
	//				b[i][j] = 1
	//			} else if dataSupport[i].ActiveTickets >= 9 && dataSupport[i].ActiveTickets < 16 {
	//				b[i][j] = 2
	//			} else {
	//				b[i][j] = 3
	//			}
	//		} else {
	//			wait := 0
	//			wait = 60 / 18 * dataSupport[i].ActiveTickets
	//			b[i][j] = wait
	//		}
	//	}
	//}
	//fmt.Println(b)
	return dataSupportArray
}
