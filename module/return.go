package module

type ChanInfo struct {
	AssetId string `json:"assetId"`
	Status  bool   `json:"status"`
	Error   string `json:"error"`
}

type ReturnInfo struct {
	Success bool   `json:"success"`
	Info    string `json:"info"`
}
