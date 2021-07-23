package transfer

import (
	"github.com/DABronskikh/bgo-3_04.1/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	testBank := card.NewService("Test bank")
	testBank.IssueCard("Visa", "RUB")
	testBank.IssueCard("MasterCard", "RUB")

	testBankTransfer := NewService(testBank)
	testBankTransfer.IssueCommission(true, true, 0, 0)
	testBankTransfer.IssueCommission(true, false, 0.5, 10_00)
	testBankTransfer.IssueCommission(false, false, 1.5, 30_00)

	testBankTransfer1 := testBankTransfer
	testBankTransfer2 := testBankTransfer
	testBankTransfer3 := testBankTransfer
	testBankTransfer4 := testBankTransfer
	testBankTransfer5 := testBankTransfer
	testBankTransfer6 := testBankTransfer

	type fields struct {
		CardSvc    *card.Service
		Commission []*Commission
	}
	type args struct {
		from   string
		to     string
		amount int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantOk    bool
	}{
		{
			name: "Карта своего банка -> Карта своего банка (денег достаточно)",
			fields: fields{
				CardSvc:    testBankTransfer1.CardSvc,
				Commission: testBankTransfer1.Commission,
			},
			args: args{
				from:   testBankTransfer1.CardSvc.Cards[0].Number,
				to:     testBankTransfer1.CardSvc.Cards[1].Number,
				amount: 1000_00,
			},
			wantTotal: 1000_00,
			wantOk:    true,
		},
		{
			name: "Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields: fields{
				CardSvc:    testBankTransfer2.CardSvc,
				Commission: testBankTransfer2.Commission,
			},
			args: args{
				from:   testBankTransfer2.CardSvc.Cards[0].Number,
				to:     testBankTransfer2.CardSvc.Cards[1].Number,
				amount: 100_000_00,
			},
			wantTotal: 100_000_00,
			wantOk:    false,
		},
		{
			name: "Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields: fields{
				CardSvc:    testBankTransfer3.CardSvc,
				Commission: testBankTransfer3.Commission,
			},
			args: args{
				from:   testBankTransfer3.CardSvc.Cards[0].Number,
				to:     "001",
				amount: 100_00,
			},
			wantTotal: 110_00,
			wantOk:    true,
		},
		{
			name: "Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields: fields{
				CardSvc:    testBankTransfer4.CardSvc,
				Commission: testBankTransfer4.Commission,
			},
			args: args{
				from:   testBankTransfer4.CardSvc.Cards[0].Number,
				to:     "001",
				amount: 100_000_00,
			},
			wantTotal: 100_500_00,
			wantOk:    false,
		},
		{
			name: "Карта чужого банка -> Карта своего банка",
			fields: fields{
				CardSvc:    testBankTransfer5.CardSvc,
				Commission: testBankTransfer5.Commission,
			},
			args: args{
				from:   "001",
				to:     testBankTransfer5.CardSvc.Cards[0].Number,
				amount: 1000_00,
			},
			wantTotal: 1000_00,
			wantOk:    true,
		},
		{
			name: "Карта чужого банка -> Карта чужого банка",
			fields: fields{
				CardSvc:    testBankTransfer6.CardSvc,
				Commission: testBankTransfer6.Commission,
			},
			args: args{
				from:   "001",
				to:     "002",
				amount: 100_00,
			},
			wantTotal: 130_00,
			wantOk:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:    tt.fields.CardSvc,
				Commission: tt.fields.Commission,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
