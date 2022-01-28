package okex

/*
 OKEX futures contract api request params
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

/*
 Create a new order
 ClientOid: You setting order id.(optional)
 Type: The execution type @see file: futures_constants.go
 InstrumentId: The id of the futures, eg: BTC_USD_0331
 Price: The order price: Maximum 1 million
 Amount: The order amount: Maximum 1 million
 MatchPrice: Match best counter party price (BBO)? 0: No 1: Yes   If yes, the 'price' field is ignored
 LeverRate: lever, default 10.
*/
type FuturesNewOrderParams struct {
	InstrumentId string `json:"instrument_id"`
	Leverage     string `json:"leverage"`
	OrderType    string `json:"order_type,omitempty"`
	FuturesBatchNewOrderIte