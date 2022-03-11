package internal

import (
	"fmt"
)

type Results interface {
	GetResultData(sms []SMSData, mms []MMSData, email []EmailData, voiceCall []VoiceCallData, billing BillingData,
		support []SupportData, incident []IncidentData) ResultT
}

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
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
	support []SupportData, incident []IncidentData) ResultT {

	var supportData []int
	dataMMSArray = prepareMMS(mms)
	dataSMSArray = prepareSMS(sms)
	dataEmailMap = prepareEmail(email)
	supportData = prepareSupport(support)

	ResultS.SMS = dataSMSArray
	//ResultS.SMS = nil
	ResultS.MMS = dataMMSArray
	ResultS.VoiceCall = voiceCall
	ResultS.Billing = billing
	ResultS.Incidents = incident
	ResultS.Support = supportData
	ResultS.Email = dataEmailMap

	if ResultS.SMS != nil && ResultS.SMS != nil && ResultS.VoiceCall != nil && (ResultS.Billing) == billing &&
		ResultS.Incidents != nil && ResultS.Support != nil && ResultS.Email != nil {
		Result.Status = true
		Result.Data = ResultS
	} else {
		Result.Status = false
		Result.Error = fmt.Sprintf("Error on collect data")
	}

	return Result
}
