package bybit

import (
	"context"
	"fmt"
)

func ExampleGetRiskLimit() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarketRiskLimits(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetFundingRate() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetFundingRateHistory(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetHistoryVolatility() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET), WithDebug(true))
	params := map[string]interface{}{"category": "option", "baseCoin": "BTC"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetHistoryVolatility(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetInstrumentInfo() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET), WithDebug(true))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetInstrumentInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetLongShortRatio() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT", "period": "5min"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetLongShortRatio(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetIndexPriceKline() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDTT", "interval": "1"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetIndexPriceKline(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetPreimumIndexPriceKline() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT", "interval": "1"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetPremiumIndexPriceKline(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetMarkPriceKline() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT", "interval": "1"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarkPriceKline(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetMarketInsurance() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET), WithDebug(true))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarketInsurance(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetMarketKline() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT", "interval": "1"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarketKline(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetMarketTicker() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarketTickers(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetOpenInterest() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetOpenInterests(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetOrderbookInfo() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET), WithDebug(true))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetOrderBookInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetRecentTrade() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetPublicRecentTrades(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetDeliveryPrice() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetDeliveryPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleGetServerTime() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	serverResult, err := client.NewUtaBybitServiceNoParams().GetServerTime(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}
