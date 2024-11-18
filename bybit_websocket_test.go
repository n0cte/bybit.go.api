package bybit

import "fmt"

func ExampleWebSocket_SendSubscription() {
	ws := NewBybitPrivateWebSocket("wss://stream-testnet.bybit.com/v5/private", "d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", func(message string) error {
		fmt.Println("Received:", message)
		return nil
	}, WithPingInterval(2))
	_, _ = ws.Connect().SendSubscription([]string{"order", "position", "wallet"})
	select {}
}

func ExampleWebSocket_SendSubscription_to_channel() {
	ws := NewBybitPublicWebSocket("wss://stream.bybit.com/v5/public/spot", func(message string) error {
		fmt.Println("Received:", message)
		return nil
	})
	_, _ = ws.Connect().SendSubscription([]string{"orderbook.1.BTCUSDT", "orderbook.1.ETHUSDT"})
	select {}
}
