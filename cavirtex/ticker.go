package cavirtex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shopspring/decimal"
)

type TickerData struct {
	Last   decimal.Decimal
	High   decimal.Decimal
	Low    decimal.Decimal
	Sell   decimal.Decimal
	Buy    decimal.Decimal
	Volume decimal.Decimal
}

func (td *TickerData) String() string {
	return fmt.Sprintf("Last: %s, High: %s, Low: %s, Sell: %s, Buy: %s, Volume: %s", td.Last, td.High, td.Low, td.Sell, td.Buy, td.Volume)
}

func (ex *CaVirtex) GetTicker(base, counter string) (*TickerData, error) {
	pair := fmt.Sprintf("%s%s", base, counter)
	resp, err := http.Get(fmt.Sprintf("https://www.cavirtex.com/api2/ticker.json?currencypair=%s", pair))
	if err != nil {
		return nil, err
	}

	result := struct {
		Ticker map[string]TickerData
	}{}

	cnt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(cnt, &result)
	if err != nil {
		return nil, err
	}

	td, present := result.Ticker[pair]
	if !present {
		return nil, fmt.Errorf("pair %q not present in API response", pair)
	}

	return &td, nil
}
