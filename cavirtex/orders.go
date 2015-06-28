package cavirtex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/shopspring/decimal"
)

type OrderMode int

const ORDER_SELL = OrderMode(0)
const ORDER_BUY = OrderMode(1)

func (ex *CaVirtex) PlaceOrder(mode OrderMode, amount decimal.Decimal, price decimal.Decimal, base, counter string) (status, id string, err error) {
	textMode := "sell"
	if mode == ORDER_BUY {
		textMode = "buy"
	}

	resp, err := ex.AuthenticatedPost("order", map[string]string{
		"currencypair": fmt.Sprintf("%s%s", base, counter),
		"mode":         textMode,
		"amount":       amount.String(),
		"price":        price.String(),
	})
	if err != nil {
		return "", "", err
	}

	result := struct {
		Status  string
		Message string
		ApiRate int
		Order   struct {
			Status string
			ID     int
		}
	}{}

	cnt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	err = json.Unmarshal(cnt, &result)
	if err != nil {
		return "", "", err
	}

	if result.Status != "ok" {
		return "", "", fmt.Errorf("status=%s: %s", result.Status, result.Message)
	}

	return result.Order.Status, fmt.Sprintf("%d", result.Order.ID), nil
}

func (ex *CaVirtex) CancelOrder(id string) error {
	resp, err := ex.AuthenticatedPost("order_cancel", map[string]string{
		"id": id,
	})
	if err != nil {
		return err
	}

	result := struct {
		Status  string
		Message string
		ApiRate int
		ID      int
	}{}

	cnt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(cnt, &result)
	if err != nil {
		return err
	}

	if result.Status != "ok" {
		return fmt.Errorf("status=%s: %s", result.Status, result.Message)
	}

	return nil
}
