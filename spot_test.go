package okex

/*
 OKEX Spot Rest Api. Unit test
 @author Steven Li
 @date 2018-12-27
 @version 1.0.0
*/

// import (
// 	"fmt"
// 	"github.com/stretchr/testify/assert"
// 	"testing"
// )

// func TestGetSpotAccount(t *testing.T) {
// 	c := NewTestClient()
// 	sp, err := c.GetSpotAccounts()
// 	assert.True(t, err == nil)
// 	fmt.Printf("%+v err %+v", sp, err)
// }

// func TestPlaceSpotOrder(t *testing.T) {
// 	c := NewTestClient()
// 	sp, err := c.PostSpotOrder(true, false, "ETH-USDT", "0.5", "0.04")
// 	assert.True(t, err == nil)
// 	fmt.Printf("%+v err %+v", sp, err)
// }

// func TestGetAllPendingSpotOrders(t *testing.T) {
// 	c := NewTestClient()
// 	sp, err := c.GetAllPendingOrders()
// 	assert.True(t, err == nil)
// 	fmt.Printf("%+v err %+v", sp, err)
// }