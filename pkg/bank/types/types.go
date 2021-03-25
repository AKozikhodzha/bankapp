package types

//Money представляет собой денежную сумму в минимальных единицах(центы, копейки, дирамы и т.д.)
type Money int64

//Cyrrency представляет код валюты
type Currency string

//Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//Payment
type Payment struct {
	ID     int
	Amount Money
}

//PAN представляет номер карты
type PAN string

//Card представляет информацию о платежной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money //использовали Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

type PaymentSource struct {
	Type    string //'card'
	Number  string // номер вида "5058 хххх хххх 8888"
	Balance Money  // баланс в дирамах

}
