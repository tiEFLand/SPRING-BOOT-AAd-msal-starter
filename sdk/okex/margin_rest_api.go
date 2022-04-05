
package okex

import (
	"strings"
)

/*
币币杠杆账户信息
获取币币杠杆账户资产列表，查询各币种的余额、冻结和可用等信息。

限速规则：20次/2s
HTTP请求
GET /api/margin/v3/accounts
*/
func (client *Client) GetMarginAccounts() (*[]map[string]interface{}, error) {
	r := []map[string]interface{}{}

	if _, err := client.Request(GET, MARGIN_ACCOUNTS, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*