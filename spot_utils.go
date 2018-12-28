package okex

/*
 OKEX api result definition
 @author Steven LI
 @date 2018-12-27
 @version 1.0.0
*/

func getBuySell (buySide bool) string {
	if buySide == true {
		return "buy"
	}
	return "sell"
}

func getMarketLimit (limitOrder bool) string {
	if limitOrder == true {
		return "limit"
	}
	return "market"
}

// func getMarketLimit (marketOrder bool) string {
// 	if marketOrder == true {
// 		return "market"
// 	}
// 	return "limit"
// }
