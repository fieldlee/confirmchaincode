package module

type File struct {
	FileName       string `json:"filename"`
	FileHash       string `json:"filehash"`
	FileCreateTime string `json:"filetime"`
	FileUrl        string `json:"fileurl"`
}

type Confirm struct {
	TxId        string `json:"txId"`
	AssetId     string `json:"assetId"`
	Opinion     string `json:"opinion"`
	Operation   string `json:"operation"` // 操作内容
	Operator    string `json:"operator"`  // 确权信息操作人
	Files       []File `json:"files"`
	OperateTime int64  `json:"operateTime"`
	Signature   string `json:"signature"`
	ChainUser   string `json:"chainUser"`
}

// 资产信息
type Asset struct {
	TxId          string    `json:"txId"`
	AssetId       string    `json:"assetId"`
	AssetName     string    `json:"assetName"`
	AssetAbstract string    `json:"assetAbstract"`
	Operation     string    `json:"operation"`
	Operator      string    `json:"operator"`    // 资产操作人
	OperateTime   int64     `json:"operateTime"` // 资产创建时间
	ChainUser     string    `json:"chainUser"`
	Files         []File    `json:"files"`
	Status        string    `json:"status"`
	Confirm       []Confirm `json:"confirms"`
}

// 交易日志结构
type Transfer struct {
	TxHash      string        `json:"txHash"`
	OperateTime int64         `json:"operateTime"`
	Operation   string        `json:"operation"` // 操作内容
	Operator    string        `json:"operator"`  // 确权信息操作人
	Register    RegitserParam `json:"register"`
	Confirm     ConfirmParam  `json:"confirm"`
}
