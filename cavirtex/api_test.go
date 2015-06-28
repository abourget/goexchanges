package cavirtex

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/abourget/cavirtex"
	"github.com/shopspring/decimal"
)

// Example of a main() function using this lib.
func Example() {
	virtex := &CaVirtex{
		APIKey:   "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		APIToken: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
	}

	// ***

	td, err := virtex.GetTicker("BTC", "CAD")
	if err != nil {
		log.Println("Error getting ticker:", err)
		os.Exit(1)
	}
	fmt.Println("Got it:", td)

	// ***

	trades, err := virtex.GetTrades("BTC", "CAD")
	if err != nil {
		log.Println("Error getting trades:", err)
		os.Exit(1)
	}
	fmt.Println("Got trades:")
	for _, tr := range trades {
		fmt.Printf("  %s\n", tr)
	}

	// ***

	depth, err := virtex.GetDepth("BTC", "CAD")
	if err != nil {
		log.Println("Error getting depth:", err)
		os.Exit(1)
	}
	fmt.Println("Got depth:")
	fmt.Println("  ", depth)

	// ***

	balances, err := virtex.Balances()
	if err != nil {
		log.Println("Error getting balances:", err)
		os.Exit(1)
	}
	fmt.Println("Got balances:", balances)

}

func Example_orders() {
	virtex := &CaVirtex{
		APIKey:   "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		APIToken: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
	}

	status, id, err := virtex.PlaceOrder(
		cavirtex.ORDER_SELL,
		decimal.New(5, -2),
		decimal.New(1, 3),
		"BTC", "CAD",
	)
	if err != nil {
		log.Println("Error placing order:", err)
		os.Exit(1)
	}
	fmt.Println("Got order reply, status:", status, "id:", id)

	time.Sleep(1 * time.Second)

	err = virtex.CancelOrder(id)
	if err != nil {
		log.Println("Error canceling order:", err)
	}
	fmt.Println("Canceled order", id)
}
