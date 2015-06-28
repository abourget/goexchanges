package cavirtex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

type Trade struct {
	ID              uint64
	Date            time.Time
	Amount          decimal.Decimal
	Rate            decimal.Decimal
	BaseCurrency    string
	CounterCurrency string
	// We can ignore the rest, as it's just noise.. We can calculate everything based
	// of those numbers.  We are using the common terms here.
}

func (t Trade) String() string {
	return fmt.Sprintf("[%d. %s] %s %s for %s %s per %s", t.ID, t.Date, t.Amount, t.BaseCurrency, t.Rate, t.CounterCurrency, t.BaseCurrency)
}

func (ex *CaVirtex) GetTrades(base, counter string) ([]Trade, error) {
	pair := fmt.Sprintf("%s%s", base, counter)
	resp, err := http.Get(fmt.Sprintf("https://www.cavirtex.com/api2/trades.json?currencypair=%s", pair))
	if err != nil {
		return nil, err
	}

	result := struct {
		Status  string
		Message string
		ApiRate int
		Trades  []struct {
			ID     uint64
			Date   float64
			Amount decimal.Decimal
			Rate   decimal.Decimal
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

	out := make([]Trade, 0)
	for _, tr := range result.Trades {
		t := Trade{
			ID:              tr.ID,
			Date:            time.Unix(int64(tr.Date), 0),
			Amount:          tr.Amount,
			Rate:            tr.Rate,
			BaseCurrency:    base,
			CounterCurrency: counter,
		}
		out = append(out, t)
	}

	return out, nil

}
