#### Оглавление:
____
0. [Дипломный проект](https://github.com/CyrilSbrodov/diplom/edit/master/README.md#дипломный-проект).
1. [Системы](https://github.com/CyrilSbrodov/diplom/edit/master/README.md#системы).
2. [Сбор данных](https://github.com/CyrilSbrodov/diplom/edit/master/README.md#сбор-данных).
3. [Отчет о состоянии систем](https://github.com/CyrilSbrodov/diplom/edit/master/README.md#отчет-о-состоянии-систем).
____

## Дипломный проект
____

Это небольшой сетевой сервис, который принимает запросы по сети и возвращает данные о состоянии систем.
____

### Системы: 
- [SMS](https://github.com/CyrilSbrodov/diplom/blob/master/internal/sms.go):
```GO
type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
} 
```

- [MMS](https://github.com/CyrilSbrodov/diplom/blob/master/internal/mms.go):
```GO
type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}
```

- [Voice Call](https://github.com/CyrilSbrodov/diplom/blob/master/internal/voice.go):
```GO
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
```

- [Support](https://github.com/CyrilSbrodov/diplom/blob/master/internal/support.go): 
```GO
type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}
```

- [Email](https://github.com/CyrilSbrodov/diplom/blob/master/internal/email.go): 
```GO
type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}
```

- [Incidents](https://github.com/CyrilSbrodov/diplom/blob/master/internal/incident.go): 
```GO
type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы: active и closed
}
```

- [Billings](https://github.com/CyrilSbrodov/diplom/blob/master/internal/billing.go): 
```GO
type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}
```
____

### Сбор данных:
- [Result](https://github.com/CyrilSbrodov/diplom/blob/master/internal/result.go): 
```GO
type ResultSetT struct {
	SMS       [][]SMSData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   BillingData              `json:"billing"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}
```
____

### Отчет о состоянии систем:
![Отчет:](https://github.com/CyrilSbrodov/diplom/blob/master/2022-05-10_12-00-16.png "Отчет")
____
Для работы димпломного проекта нужен симулятор (git clone https://github.com/antondzhukov/skillbox-diploma.git .)
