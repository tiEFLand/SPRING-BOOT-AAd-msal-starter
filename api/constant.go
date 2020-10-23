
package api

type JSONReply struct {
	ErrorCode        int         `json:"error_code"`
	ErrorDescription string      `json:"error_desc"`
	Payload          interface{} `json:"payload"`
}