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

var (
	sms          SMSData
	dataSMS      []SMSData
	dataSMSArray [][]SMSData
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func SMSSystem() []SMSData {
	fileName := "simulator/sms.data"
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
		if len(res) != 4 {
			continue
		} else {
			if _, exist := data2.Country[res[0]]; !exist {
				continue
			} else {
				sms.Country = res[0]
			}
			if _, exist := data2.Providers[res[3]]; !exist {
				continue
			} else {
				sms.Provider = res[3]
			}
			sms.Bandwidth = res[1]
			sms.ResponseTime = res[2]
			dataSMS = append(dataSMS, sms)
		}
	}
	fmt.Println("=================")
	fmt.Println("Состояние системы SMS:")

	for i := 0; i < len(dataSMS); i++ {
		fmt.Println(dataSMS[i])
	}

	return dataSMS
}

func prepareSMS(sms []SMSData) [][]SMSData {

	copySMS := make([]SMSData, len(sms))
	for i := 0; i < len(sms); i++ {
		sms[i].Country = data2.Country[sms[i].Country]
	}
	sort.Slice(sms, func(i, j int) bool {
		return sms[i].Provider < sms[j].Provider
	})
	dataSMSArray = append(dataSMSArray, sms)
	copy(copySMS, sms)
	sort.Slice(copySMS, func(i, j int) bool {
		return copySMS[i].Country < copySMS[j].Country
	})
	dataSMSArray = append(dataSMSArray, copySMS)

	return dataSMSArray
}
