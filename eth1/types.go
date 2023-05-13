// description: 只能合约相关
// @author renshiwei
// Date: 2022/11/15 17:23

package eth1

type Tx struct {
	From    string `json:"from,omitempty"`
	To      string `json:"to"`
	Value   string `json:"value"`
	Data    string `json:"data"`
	ChainId int    `json:"chain_id"`
}
