package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.)
type Money int64

// Category представляет категорию
type Category string

// Status представляет статус платежа
type Status string

// Предопределеннве статусы платежей
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

//Payment представляет информацию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}
