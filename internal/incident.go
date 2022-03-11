package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы: active и closed
}

var (
	incident     IncidentData
	dataIncident []IncidentData
)

func IncidentSystem() []IncidentData {
	resp, err := http.Get("http://127.0.0.1:8383/accendent")

	if err != nil {
		errors.New(fmt.Sprintf("не удалось получить данные. ошибка: %v", err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.New(fmt.Sprintf("не удалось прочитать данные. ошибка: %v", err))
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		if err := json.Unmarshal(body, &dataIncident); err != nil {
			errors.New(fmt.Sprintf("не удалось записать данные. ошибка: %v", err))
		}

		fmt.Println("=================")
		fmt.Println("Состояние системы Incident:")

		for i := 0; i < len(dataIncident); i++ {
			fmt.Println(dataIncident[i])
		}
	} else if resp.StatusCode == 500 {
		fmt.Println("Ошибка", dataIncident)
	}
	sort.Slice(dataIncident, func(i, j int) bool {
		return dataIncident[i].Status < dataIncident[j].Status
	})
	return dataIncident
}
