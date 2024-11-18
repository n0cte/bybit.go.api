package bybit

import (
	"context"
	"fmt"
)

func ExampleBybitClientRequest_GetAccountInfo() {
	client := NewBybitHttpClient("d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", WithBaseURL(TESTNET))
	accountResult, err := client.NewUtaBybitServiceNoParams().GetAccountInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(accountResult))
}

func ExampleBybitClientRequest_GetFeeRates() {
	client := NewBybitHttpClient("d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", WithBaseURL(TESTNET))
	params := map[string]interface{}{"symbol": "BTCUSDT", "category": "linear"}
	accountResult, err := client.NewUtaBybitServiceWithParams(params).GetFeeRates(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(accountResult))
}

func ExampleBybitClientRequest_GetTransactionLog() {
	client := NewBybitHttpClient("d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", WithBaseURL(TESTNET))
	params := map[string]interface{}{"accountType": "UNIFIED", "category": "linear"}
	accountResult, err := client.NewUtaBybitServiceWithParams(params).GetTransactionLog(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(accountResult))
}

func EsamGetAccountWallet() {
	client := NewBybitHttpClient("d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", WithBaseURL(TESTNET))
	params := map[string]interface{}{"accountType": "UNIFIED"}
	accountResult, err := client.NewUtaBybitServiceWithParams(params).GetFeeRates(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(accountResult))
}
