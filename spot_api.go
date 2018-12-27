package okex

/*
 OKEX Swap Api
 @author Steven Li
 @date 2018-12-27
 @version 1.0.0
*/

import (
	// "errors"
)

/*
GET /api/spot/v3/accounts
*/

func (client *Client) getSpotAccounts(uri string) (*SpotAccounts, error) {
	sa := SpotAccounts{}
	if _, err := client.Request(GET, uri, nil, &sa); err != nil {
		return nil, err
	}
	return &sa, nil
}

/*
GET /api/spot/v3/accounts
*/
func (client *Client) GetSpotAccounts() (*SpotAccounts, error) {
	return client.getSpotAccounts(SPOT_ACCOUNTS)
}