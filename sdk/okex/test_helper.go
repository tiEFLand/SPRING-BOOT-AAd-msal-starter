
package okex

/*
 Get a http client
*/

func GetDefaultConfig() *Config {
	var config Config
	config.Endpoint = "http://192.168.80.113:9300/"
	config.WSEndpoint = "ws://192.168.80.113:10442/"
	config.ApiKey = "bb57a1b3-6257-47ff-b06c-faafc4d28fad"
	config.SecretKey = ""
	config.Passphrase = ""
	config.TimeoutSecond = 45
	config.IsPrint = true
	config.I18n = ENGLISH

	return &config
}

func NewTestClient() *Client {