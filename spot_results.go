package okex

/*
 OKEX api result definition
 @author Steven LI
 @date 2018-12-27
 @version 1.0.0
*/

type SpotAccountInfo struct {
	Currency      string `json:"currency"`
	Balance         string `json:"balance"`
	Hold      string `json:"hold"`
	Available string `json:"available"`
	Id       string `json:"id"`
}

type SpotAccounts struct {
	BizWarmTips
	Info []SpotAccountInfo `json:"info"`
}