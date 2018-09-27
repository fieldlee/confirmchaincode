package module

// 资产上链
type RegitserParam struct {
	AssetId       string `json:"assetId"`
	AssetName     string `json:"assetName"`
	AssetAbstract string `json:"assetAbstract"`
	Operation     string `json:"operation"` // 操作内容
	Operator      string `json:"operator"`  // 资产操作人
	Files         []File `json:"files"`
}

// 确权信息
type ConfirmParam struct {
	AssetId   string `json:"assetId"`
	Opinion   string `json:"opinion"`
	Operation string `json:"operation"` // 操作内容
	Operator  string `json:"operator"`  // 确权信息操作人
	Signature string `json:"signature"`
	Files     []File `json:"files"`
}

// 上链信息
type PlatParam struct {
	AssetId   string `json:"assetId"`
	Operation string `json:"operation"` // 操作内容
	Operator  string `json:"operator"`  // 平台上链操作人
}

type QueryParam struct {
	AssetId string `json:"assetId"`
}
