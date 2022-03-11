package internal

import (
	data2 "diplom/data"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

var (
	mms          MMSData
	dataMMS      []MMSData
	dataMMSArray [][]MMSData
)

func MMSSystem() []MMSData {
	resp, err := http.Get("http://127.0.0.1:8383/mms")

	if err != nil {
		errors.New(fmt.Sprintf("не удалось получить данные. ошибка:", err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.New(fmt.Sprintf("не удалось прочитать данные. ошибка:", err))
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		if err := json.Unmarshal(body, &dataMMS); err != nil {
			errors.New(fmt.Sprintf("не удалось прочитать данные. ошибка:", err))
		}
		for i := 0; i < len(dataMMS); i++ {
			if _, exist := data2.Country[dataMMS[i].Country]; !exist {
				dataMMS[i] = dataMMS[len(dataMMS)-1]
				dataMMS[len(dataMMS)-1] = MMSData{}
				dataMMS = dataMMS[:len(dataMMS)-1]
			}
			if _, exist := data2.Providers[dataMMS[i].Provider]; !exist {
				dataMMS[i] = dataMMS[len(dataMMS)-1]
				dataMMS[len(dataMMS)-1] = MMSData{}
				dataMMS = dataMMS[:len(dataMMS)-1]
			}
		}
		fmt.Println("=================")
		fmt.Println("Состояние системы MMS:")

		for i := 0; i < len(dataMMS); i++ {
			fmt.Println(dataMMS[i])
		}

	} else if resp.StatusCode == 500 {
		fmt.Println("Ошибка", dataMMS)
	}
	return dataMMS
}

func prepareMMS(mms []MMSData) [][]MMSData {

	copyMMS := make([]MMSData, len(mms))
	for i := 0; i < len(mms); i++ {
		mms[i].Country = data2.Country[mms[i].Country]
	}
	sort.Slice(mms, func(i, j int) bool {
		return mms[i].Provider < mms[j].Provider
	})
	dataMMSArray = append(dataMMSArray, mms)
	copy(copyMMS, mms)
	sort.Slice(copyMMS, func(i, j int) bool {
		return copyMMS[i].Country < copyMMS[j].Country
	})
	dataMMSArray = append(dataMMSArray, copyMMS)

	return dataMMSArray
}
