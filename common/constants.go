package common

var ERR = map[string]string{
	"NONE": "000",
}

var STATUS = map[string]string{
	"Init":       "Inited",
	"Confirming": "Confirming",
	"Confirm":    "Confirmed",
}

const (
	//下划线
	ULINE = "_"
	//产品信息KEY
	ASSET_INFO = "ASSET_INFO"
	// 产品交易信息
	ASSET_ACTION = "ASSET_ACTION"
)
