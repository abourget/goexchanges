package cavirtex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shopspring/decimal"
)

type Depth struct {
	Asks            []Ask
	Bids            []Bid
	BaseCurrency    string
	CounterCurrency string
}

type Ask struct {
	Price  decimal.Decimal
	Amount decimal.Decimal
}
type Bid Ask

func (d Depth) String() string {
	return fmt.Sprintf("Asks: %+v,     Bids: %+v", d.Asks, d.Bids)
}

func (ex *CaVirtex) GetDepth(base, counter string) (*Depth, error) {
	pair := fmt.Sprintf("%s%s", base, counter)
	resp, err := http.Get(fmt.Sprintf("https://www.cavirtex.com/api2/orderbook.json?currencypair=%s", pair))
	if err != nil {
		return nil, err
	}

	result := struct {
		Status  string
		Message string
		ApiRate int
		OrderBook struct {
			Bids    [][]decimal.Decimal
			Asks    [][]decimal.Decimal
		}
	}{}

	cnt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(cnt, &result)
	if err != nil {
		return nil, err
	}

	d := Depth{
		BaseCurrency:    base,
		CounterCurrency: counter,
	}
	for _, bid := range result.OrderBook.Bids {
		d.Bids = append(d.Bids, Bid{
			Price:  bid[0],
			Amount: bid[1],
		})
	}
	for _, ask := range result.OrderBook.Asks {
		d.Asks = append(d.Asks, Ask{
			Price:  ask[0],
			Amount: ask[1],
		})
	}

	return &d, nil
}
