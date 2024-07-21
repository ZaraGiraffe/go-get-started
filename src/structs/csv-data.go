package structs

type UserData struct {
	UserId string
	Currency string
	Timestamp int64
	Delta float64
}


type MarketData struct {
	Symbol string
	Timestamp int64
	Price float64
}