package okex

/*
 OKEX api result definition
 @author Steven LI
 @date 2018-12-27
 @version 1.0.0
*/

type SpotAccountInfo struct {
	Frozen string `json:"frozen"`
	Currency      string `json:"currency"`
	Balance         string `json:"balance"`
	Hold      string `json:"hold"`
	Holds      string `json:"holds"`
	Available string `json:"available"`
	Id       string `json:"id"`
}

// TODO: remove deprecated
// type SpotAccount struct {
// 	BizWarmTips
// 	Info []SpotAccountInfo `json:"info"`
// }

type SpotOrderInfo struct {
	OrderId string `json:"order_id"`
	ClientOid      string `json:"client_oid"`
	Result         bool `json:"result"`
}

type SpotPendingOrderInfo struct {
	OrderId string `json:"order_id"`
	Price string `json:"price"` 
	Size string `json:"size"` 
	InstrumentId string `json:"instrument_id"` 
	Side string `json:"side"` 
	Type string `json:"type"` 
	Timestamp string `json:"timestamp"` 
	FilledSize string `json:"filled_size"` 
	FilledNotional string `json:"filled_notional"` 
	Status string `json:"status"` 
}
