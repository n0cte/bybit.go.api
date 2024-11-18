package bybit

import (
	"context"
	"fmt"
)

func ExampleBybitClientRequest_GetCoinInfo() {
	client := NewBybitHttpClient("d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", WithBaseURL(TESTNET))
	params := map[string]interface{}{"coin": "USDT"}
	assetResult, err := client.NewUtaBybitServiceWithParams(params).GetCoinInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(assetResult))
}

func ExampleBybitClientRequest_GetTransferableCoin() {
	client := NewBybitHttpClient("d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", WithBaseURL(TESTNET))
	params := map[string]interface{}{"fromAccountType": "UNIFIED", "toAccountType": "CONTRACT"}
	response, err := client.NewUtaBybitServiceWithParams(params).GetTransferableCoin(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(response))
}
