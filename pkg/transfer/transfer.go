package transfer

import "github.com/DABronskikh/bgo-3_04.1/pkg/card"

type Commission struct {
	from       bool
	to         bool
	percentage float64
	minAmount  int64
}

type Service struct {
	CardSvc    *card.Service
	Commission []*Commission
}

func NewService(cardSvc *card.Service) *Service {
	return &Service{
		CardSvc: cardSvc,
	}
}

func (s *Service) IssueCommission(from, to bool, percentage float64, minAmount int64) {
	commission := &Commission{
		from:       from,
		to:         to,
		percentage: percentage,
		minAmount:  minAmount,
	}
	s.Commission = append(s.Commission, commission)
}

func (s *Service) Card2Card(from string, to string, amount int64) (total int64, ok bool) {
	fromCard := s.CardSvc.SearchByNumber(from)
	fromBool := fromCard != nil
	toCard := s.CardSvc.SearchByNumber(to)
	toBool := toCard != nil

	commission := s.searchCommission(fromBool, toBool)
	percentage := commission.percentage
	minAmount := commission.minAmount

	var sumCommission int64
	if percentage != 0 {
		sumCommission = int64(float64(amount) * percentage / 100)
	}

	if sumCommission < minAmount {
		sumCommission = minAmount
	}

	total = amount + sumCommission
	if (!fromBool && !toBool) || !fromBool {
		ok = true
	}

	if fromBool {
		newBalance := fromCard.Balance - total
		if newBalance > 0 {
			fromCard.Balance = newBalance
			ok = true
		}
	}

	if toBool {
		toCard.Balance += amount
	}

	return total, ok
}

func (s *Service) searchCommission(from bool, to bool) *Commission {
	for _, candidate := range s.Commission {
		if candidate.from == from && candidate.to == to {
			return candidate
		}
	}

	return &Commission{
		from:       false,
		to:         false,
		percentage: 0,
		minAmount:  0,
	}
}
