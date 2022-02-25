package main

import (
	"diplom/internal"
	"diplom/internal/server"
	"fmt"
)

func main() {
	sms := internal.SMSSystem()
	mms := internal.MMSSystem()
	voiceCall := internal.VoiceSystem()
	email := internal.EmailSystem()
	billing := internal.BillingSystem()
	support := internal.SupportSystem()
	incident := internal.IncidentSystem()
	result := internal.GetResultData(sms, mms, email, voiceCall, billing, support, incident)

	server := server.NewApp()
	server.Run(result)

	fmt.Println("SMS:")
	fmt.Println(internal.ResultS.SMS)
	fmt.Println("MMS:")
	fmt.Println(internal.ResultS.MMS)
	fmt.Println("Email:")
	fmt.Println(internal.ResultS.Email)
	fmt.Println("Voice Call:")
	fmt.Println(internal.ResultS.VoiceCall)
	fmt.Println("Billing:")
	fmt.Println(internal.ResultS.Billing)
	fmt.Println("Incident:")
	fmt.Println(internal.ResultS.Incidents)
	fmt.Println("Support:")
	fmt.Println(internal.ResultS.Support)
}
