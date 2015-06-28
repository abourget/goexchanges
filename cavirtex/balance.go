package cavirtex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/shopspring/decimal"
)

func (ex *CaVirtex) Balances() (map[string]decimal.Decimal, error) {
	resp, err := ex.AuthenticatedPost("balance", map[string]string{})
	if err != nil {
		return nil, err
	}

	result := struct {
		Status  string
		Message string
		ApiRate int
		Balance map[string]decimal.Decimal
	}{}

	cnt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(cnt, &result)
	if err != nil {
		return nil, err
	}

	if result.Status != "ok" {
		return nil, fmt.Errorf("status=%s: %s", result.Status, result.Message)
	}

	return result.Balance, nil
}
