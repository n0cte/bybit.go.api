package bybit

import (
	"context"
	"fmt"
)

func ExampleGetPositionList() {
	client := NewBybitHttpClient("d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "settleCoin": "USDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetPositionList(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleSetPositionLeverage() {
	client := NewBybitHttpClient("d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "XRPUSDT", "buyLeverage": "20", "sellLeverage": "20"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).SetPositionLeverage(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}
