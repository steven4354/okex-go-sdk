package okex

/*
 OKEX swap api parameter's definition
 @author Steven Li
 @date 2018-12-27
 @version 1.0.0
*/

// type BasePlaceOrderInfo struct {
// 	ClientOid  string `json:"client_oid"`
// 	Price      string `json:"price"`
// 	MatchPrice string `json:"match_price"`
// 	Type       string `json:"type"`
// 	Size       string `json:"size"`
// }

type PlaceSpotLimitOrderInfo struct {
	Type  string `json:"type"`
	Side      string `json:"side"`
	InstrumentId string `json:"instrument_id"`
	MarginTrading       byte `json:"margin_trading"`
	Size string `json:"size"`
	Price string `json:"price"`
}

type PlaceSpotMarketOrderInfo struct {
	Type  string `json:"type"`
	Side      string `json:"side"`
	InstrumentId string `json:"instrument_id"`
	MarginTrading       byte `json:"margin_trading"`
	QtyToSell string `json:"size"`
	QtyToBuy string `json:"notional"`
}

// type PlaceOrdersInfo struct {
// 	InstrumentId string                `json:"instrument_id"`
// 	OrderData    []*BasePlaceOrderInfo `json:"order_data"`
// }
