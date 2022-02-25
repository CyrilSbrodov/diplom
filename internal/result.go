package internal

import (
	"encoding/json"
	"fmt"
)

type Results interface {
	GetResultData(sms []SMSData, mms []MMSData, email []EmailData, voiceCall []VoiceCallData, billing BillingData,
		support []SupportData, incident []IncidentData) ResultSetT
}

type ResultT struct {
	Status bool       `json:"status"` // true, если все этапы сбора данных прошли успешно, false во всех остальных случаях
	Data   ResultSetT `json:"data"`   // заполнен, если все этапы сбора данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // пустая строка если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки (детали ниже)
}
type ResultSetT struct {
	SMS       [][]SMSData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   BillingData              `json:"billing"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}

var (
	ResultS ResultSetT
	Result  ResultT
)

func GetResultData(sms []SMSData, mms []MMSData, email []EmailData, voiceCall []VoiceCallData, billing BillingData,
	support []SupportData, incident []IncidentData) ResultSetT {

	var supportData []int
	dataMMSArray = prepareMMS(mms)
	dataSMSArray = prepareSMS(sms)
	dataEmailMap = prepareEmail(email)
	supportData = prepareSupport(support)
	dataEmailMap = prepareEmail(email)

	ResultS.SMS = dataSMSArray
	ResultS.MMS = dataMMSArray
	ResultS.VoiceCall = voiceCall
	ResultS.Billing = billing
	ResultS.Incidents = incident
	ResultS.Support = supportData
	ResultS.Email = dataEmailMap

	return ResultS
}

func Name() {
	json1, _ := json.Marshal(Result)
	fmt.Println(string(json1))
	json2, _ := json.Marshal(ResultS)
	fmt.Println(string(json2))
}
