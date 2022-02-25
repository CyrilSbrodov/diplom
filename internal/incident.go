package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		if err := json.Unmarshal(body, &dataIncident); err != nil {
			log.Fatal(err)
		}

		//fmt.Println("=================")
		//fmt.Println("Состояние системы Incident:")
		//
		//for i := 0; i < len(dataIncident); i++ {
		//	fmt.Println(dataIncident[i])
		//}
	} else if resp.StatusCode == 500 {
		fmt.Println("Ошибка", dataIncident)
	}
	sort.Slice(dataIncident, func(i, j int) bool {
		return dataIncident[i].Status < dataIncident[j].Status
	})
	return dataIncident
}
