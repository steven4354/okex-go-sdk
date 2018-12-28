package okex

/*
 OKEX Swap Api
 @author Steven Li
 @date 2018-12-27
 @version 1.0.0
*/

import (
	// "errors"
	"log"
)

/*
GET /api/spot/v3/accounts
*/
func (client *Client) getSpotAccounts(uri string) (*[]SpotAccountInfo, error) {
	var sa []SpotAccountInfo
	sa = make([]SpotAccountInfo, 0)
	if _, err := client.Request(GET, uri, nil, &sa); err != nil {
		return nil, err
	}
	return &sa, nil
}

/*
GET /api/spot/v3/accounts
*/
func (client *Client) GetSpotAccounts() (*[]SpotAccountInfo, error) {
	return client.getSpotAccounts(SPOT_ACCOUNTS)
}

/*
POST /api/spot/v3/orders
*/
func (client *Client) PostSpotOrder(limit bool, buySide bool, instrumentId string, price string, qty string) (*SpotOrderInfo, error) {
	// by default margin trading is off
	var req interface{}
	if limit == true {
		req = PlaceSpotLimitOrderInfo{
			getMarketLimit(limit), 
			getBuySell(buySide), 
			instrumentId,
			1,
			qty,
			price,
		}	
	} else if limit == false && buySide == true {
		req = PlaceSpotMarketOrderInfo{
			getMarketLimit(limit), 
			getBuySell(buySide), 
			instrumentId,
			1,
			"0",
			qty,
		}	
	} else if limit == false && buySide == false {
		req = PlaceSpotMarketOrderInfo{
			getMarketLimit(limit), 
			getBuySell(buySide), 
			instrumentId,
			1,
			qty,
			"0",
		}	
	}
	
	log.Printf("req %+v", req)

	// return nil, nil

	res := SpotOrderInfo{}
	
	if _, err := client.Request(POST, SPOT_ORDER, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

/*
POST /api/spot/v3/orders_pending
*/
func (client *Client) GetAllPendingOrders() (*[]SpotPendingOrderInfo, error) {
	var sa []SpotPendingOrderInfo
	sa = make([]SpotPendingOrderInfo, 0)
	
	if _, err := client.Request(GET, SPOT_PENDING_ORDERS, nil, &sa); err != nil {
		return nil, err
	}
	return &sa, nil
}

func (client *Client) GetSpecificPendingOrders() () {
	// TODO: created this
	// https://www.okex.com/docs/en/#spot-huo-qu-wei-wan-cheng-ding-dan
}
