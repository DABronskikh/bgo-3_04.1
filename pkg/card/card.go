package card

import (
	"math/rand"
	"strconv"
)

type Card = struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
}

type Service struct {
	BankName string
	Cards []*Card
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

func (s *Service) IssueCard(issuer string, currency string)  *Card{
	card := &Card{
		Issuer:       issuer,
		Balance:      50_000_00,
		Currency:     currency,
		Number:       randNumber(),
	}

	s.Cards = append(s.Cards, card)
	return card
}

func (s *Service) SearchByNumber(number string) *Card {
	for _, card := range s.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}


type Transaction = struct {
	Id     int64
	Amount int64
	Date   int64
	MCC    string
	status string
	Type   string
	Status string
}

func TranslateMCC(code string) string {
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5533": "Автоуслуги",
		"5912": "Аптеки",
	}

	value, ok := mcc[code]
	if ok {
		return value
	}

	return "Категория не указана"
}

func randNumber() (number string) {
	return strconv.Itoa(rand.Intn(100))
}