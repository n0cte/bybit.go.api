package bybit

import (
	"context"
	"fmt"
)

func ExampleBybitClientRequest_GetMarketRiskLimits() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarketRiskLimits(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetFundingRateHistory() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetFundingRateHistory(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetHistoryVolatility() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET), WithDebug(true))
	params := map[string]interface{}{"category": "option", "baseCoin": "BTC"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetHistoryVolatility(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetInstrumentInfo() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET), WithDebug(true))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetInstrumentInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetLongShortRatio() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT", "period": "5min"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetLongShortRatio(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetIndexPriceKline() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDTT", "interval": "1"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetIndexPriceKline(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetPremiumIndexPriceKline() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT", "interval": "1"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetPremiumIndexPriceKline(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetMarkPriceKline() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT", "interval": "1"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarkPriceKline(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetMarketInsurance() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET), WithDebug(true))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarketInsurance(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetMarketKline() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT", "interval": "1"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarketKline(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetMarketTickers() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetMarketTickers(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetOpenInterests() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetOpenInterests(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetOrderBookInfo() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET), WithDebug(true))
	params := map[string]interface{}{"category": "spot", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetOrderBookInfo(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetPublicRecentTrades() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetPublicRecentTrades(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetDeliveryPrice() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	params := map[string]interface{}{"category": "linear", "symbol": "BTCUSDT"}
	serverResult, err := client.NewUtaBybitServiceWithParams(params).GetDeliveryPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}

func ExampleBybitClientRequest_GetServerTime() {
	client := NewBybitHttpClient("", "", WithBaseURL(TESTNET))
	serverResult, err := client.NewUtaBybitServiceNoParams().GetServerTime(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(PrettyPrint(serverResult))
}
