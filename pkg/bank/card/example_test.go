package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	card := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 6666",
			Balance: 1_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 7777",
			Balance: 1_000_00,
			Active:  false,
		},
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 0,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 9999",
			Balance: 2_00_00,
			Active:  true,
		},
	}
	for _, card := range PaymentSources(card) {
		fmt.Println(card.Number)
	}
	//Output:
	//5058 xxxx xxxx 6666
	//5058 xxxx xxxx 9999
}
