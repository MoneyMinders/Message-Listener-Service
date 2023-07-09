package commons

type Transaction struct {
	TransactionDate string
	Amount          string
	Merchant        string
}

type KafkaMessage struct {
	Category string `json:"category"`
	AppName  string `json:"appName"`
	Message  string `json:"message"`
}

type Merchant struct {
	CategoryId int `json:"category_id"`
	Name  string `json:"name"`
}