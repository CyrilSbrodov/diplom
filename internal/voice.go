package internal

import (
	data2 "diplom/data"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

var (
	voice     VoiceCallData
	dataVoice []VoiceCallData
)

func VoiceSystem() []VoiceCallData {
	fileName := "simulator/voice.data"
	file, err := os.Open(fileName)
	if err != nil {
		errors.New(fmt.Sprintf("не удалось получить данные. ошибка: %v", err))
	}
	defer file.Close()
	result, err := ioutil.ReadAll(file)
	if err != nil {
		errors.New(fmt.Sprintf("не удалось прочитать данные. ошибка: %v", err))
	}

	resultString := strings.Split(string(result), "\n")

	for i := 0; i < len(resultString); i++ {
		res := strings.Split(resultString[i], ";")
		if len(res) != 8 {
			continue
		} else {
			if _, exist := data2.Country[res[0]]; !exist {
				continue
			} else {
				voice.Country = res[0]
			}
			if _, exist := data2.VoiceProviders[res[3]]; !exist {
				continue
			} else {
				voice.Provider = res[3]
			}
			voice.Bandwidth = res[1]
			voice.ResponseTime = res[2]
			ConnectionStability, err := strconv.ParseFloat(res[4], 32)
			if err != nil {
				errors.New(fmt.Sprintf("ошибка конвертирования %v:", err))
			}
			voice.ConnectionStability = float32(ConnectionStability)
			voice.TTFB = stringToInt(res[5])
			voice.VoicePurity = stringToInt(res[6])
			voice.MedianOfCallsTime = stringToInt(res[7])
			dataVoice = append(dataVoice, voice)
		}
	}
	fmt.Println("=================")
	fmt.Println("Состояние системы Voice:")

	for i := 0; i < len(dataVoice); i++ {
		fmt.Println(dataVoice[i])
	}
	return dataVoice
}

func stringToInt(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		errors.New(fmt.Sprintf("%v не является числом, ошибка: %v:", s, err))
	}
	return number
}
