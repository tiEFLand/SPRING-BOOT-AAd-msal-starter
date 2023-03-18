package okex

/*
 OKEX websocket api constants
 @author Lingting Fu
 @date 2018-12-27
 @version 1.0.0
*/

import "errors"

const (
	WS_API_HOST = "okexcomreal.bafang.com:10442"
	WS_API_URL  = "wss://real.okex.com:10442/ws/v3"

	CHNL_FUTURES_TICKER          = "futures/ticker"          // 行情数据频道
	CHNL_FUTURES_CANDLE60S       = "futures/candle60s"       // 1分钟k线数据频道
	CHNL_FUTURES_CANDLE180S      = "futures/candle180s"      // 3分钟k线数据频道
	CHNL_FUTURES_CANDLE300S      = "futures/candle300s"      // 5分钟k线数据频道
	CHNL_FUTURES_CANDLE900S      = "futures/candle900s"      // 15分钟k线数据频道
	CHNL_FUTURES_CANDLE1800S     = "futures/candle1800s"     // 30分钟k线数据频道
	CHNL_FUTURES_CANDLE3600S     = "futures/candle3600s"     // 1小时k线数据频道
	CHNL_FUTURES_CANDLE7200S     = "futures/candle7200s"     // 2小时k线数据频道
	CHNL_FUTURES_CANDLE14400S    = "futures/candle14400s"    // 4小时k线数据频道
	CHNL_FUTURES_CANDLE21600     = "futures/candle21600"     // 6小时k线数据频道
	CHNL_FUTURES_CANDLE43200S    = "futures/candle43200s"    // 12小时k线数据频道
	CHNL_FUTURES_CANDLE86400S    = "futures/candle86400s"    // 1day k线数据频道
	CHNL_FUTURES_CANDLE604800S   = "futures/candle604800s"   // 1week k线数据频道
	CHNL_FUTURES_TRADE           = "futures/trade"           // 交易信息频道
	CHNL_FUTURES_ESTIMATED_PRICE = "futures/estimated_price" //获取预估交