
package db

import "crypto-user/utils"

var (
	DB             = "crypto"
	CollectionUser = "user"
)

func init() {
	DB, _ = utils.GetConfig().Get("db.name")
}