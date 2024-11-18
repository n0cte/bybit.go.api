package bybit

import (
	"context"
	"fmt"
)

func ExampleBybitClientRequest_PlaceBatchOrder() {
	client := NewBybitHttpClient("8wYkmpLsMg10eNQyPm", "Ouxc34myDnXvei54XsBZgoQzfGxO4bkr2Zsj", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "option",
		"request": []map[string]interface{}{
			{
				"category":    "option",
				"symbol":      "BTC-10FEB23-24000-C",
				"orderType":   "Limit",
				"side":        "Buy",
				"qty":         "0.1",
				"price":       "5",
				"orderIv":     "0.1",
				"timeInForce": "GTC",
				"orderLinkId": "9b381bb1-401",
				"mmp":         false,
				"reduceOnly":  false,
			},
			{
				"category":    "option",
				"symbol":      "BTC-10FEB23-24000-C",
				"orderType":   "Limit",
				"side":        "Buy",
				"qty":         "0.1",
				"price":       "5",
				"orderIv":     "0.1",
				"timeInForce": "GTC",
				"orderLinkId": "82ee86dd-001",
				"mmp":         false,
				"reduceOnly":  false,
			},
		},
	}

	orderResult, err := client.NewClassicalBybitServiceWithParams(params).PlaceBatchOrder(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(orderResult))
}

func ExamplOrder_Do() {
	client := NewBybitHttpClient("d08Wh6P037IXAvcrL2", "gLfd1BLGU9oq6YoRZRlwXkIQRYB4n5KvXDvv", WithBaseURL(TESTNET), WithDebug(true))
	orderResult, err := client.NewPlaceOrderService("linear", "XRPUSDT", "Buy", "Market", "15").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(orderResult))
}

func ExampleBybitClientRequest_PlaceOrder() {
	client := NewBybitHttpClient("xxxx", "xxx", WithBaseURL(TESTNET))
	params := map[string]interface{}{
		"category":    "linear",
		"symbol":      "BTCUSDT",
		"side":        "Buy",
		"positionIdx": 0,
		"orderType":   "Limit",
		"qty":         "0.001",
		"price":       "10000",
		"timeInForce": "GTC",
	}
	orderResult, err := client.NewUtaBybitServiceWithParams(params).PlaceOrder(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(orderResult))
}

func ExampleBybitClientRequest_RequestTestFund() {
	client := NewBybitHttpClient("xxxx", "xxx", WithBaseURL(DEMO_ENV))
	params := map[string]interface{}{
		"adjustType": 0,
		"utaDemoApplyMoney": []map[string]interface{}{
			{
				"coin":      "USDT",
				"amountStr": "109",
			},
			{
				"coin":      "ETH",
				"amountStr": "1",
			},
		},
	}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).RequestTestFund(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}
