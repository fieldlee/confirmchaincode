package services

import (
	"confirmchaincode/common"
	"confirmchaincode/log"
	"confirmchaincode/module"
	"encoding/json"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func toRegister(stub shim.ChaincodeStubInterface, param module.RegitserParam) (tChan module.ChanInfo) {
	// 	verify product if exist or not
	jsonParam, err := stub.GetState(common.ASSET_INFO + common.ULINE + param.AssetId)
	log.Logger.Info("------------------------------------------------------------------")
	log.Logger.Info(string(jsonParam[:]))
	if jsonParam != nil {
		log.Logger.Error("Register -- get asset by assetid -- err: 已经注册" + "	assetid:" + param.AssetId)
		tChan.AssetId = param.AssetId
		tChan.Status = false
		tChan.Error = "已经注册"
		return
	}

	asset := module.Asset{}
	asset.TxId = stub.GetTxID()
	asset.AssetAbstract = param.AssetAbstract
	asset.AssetId = param.AssetId
	asset.AssetName = param.AssetName
	asset.Files = param.Files
	asset.Operation = asset.Operation
	asset.Operator = asset.Operator
	asset.ChainUser = common.GetUserFromCertification(stub)
	asset.OperateTime = time.Now().Unix()
	asset.Status = common.STATUS["Init"]

	jsonByte, err := json.Marshal(asset)
	if err != nil {
		log.Logger.Error("Register -- marshal product err:" + err.Error() + "	assetid:" + param.AssetId)
		tChan.AssetId = param.AssetId
		tChan.Status = false
		tChan.Error = err.Error()
		return
	}

	err = stub.PutState(common.ASSET_INFO+common.ULINE+param.AssetId, jsonByte)
	if err != nil {
		log.Logger.Error("Register -- putState:" + err.Error() + "	assetid:" + param.AssetId)
		tChan.AssetId = param.AssetId
		tChan.Status = false
		tChan.Error = err.Error()

		return
	}

	tChan.AssetId = param.AssetId
	tChan.Status = true
	tChan.Error = "完成"
	return
}

func toConfirm(stub shim.ChaincodeStubInterface, param module.ConfirmParam) (tChan module.ChanInfo) {
	// 	verify product if exist or not
	jsonParam, err := stub.GetState(common.ASSET_INFO + common.ULINE + param.AssetId)
	log.Logger.Info("------------------------------------------------------------------")
	log.Logger.Info(string(jsonParam[:]))
	if jsonParam == nil {
		log.Logger.Error("Confirm -- get asset by assetid -- err: 未注册" + "	assetid:" + param.AssetId)
		tChan.AssetId = param.AssetId
		tChan.Status = false
		tChan.Error = "资产未注册"
		return
	}

	asset := module.Asset{}
	err = json.Unmarshal(jsonParam, &asset)
	if err != nil {
		log.Logger.Error("Confirm -- Unmarshal asset err:" + err.Error() + "	assetid:" + param.AssetId)
		tChan.AssetId = param.AssetId
		tChan.Status = false
		tChan.Error = err.Error()
		return
	}

	// confirm
	confirm := module.Confirm{}
	confirm.TxId = stub.GetTxID()
	confirm.Files = param.Files
	confirm.AssetId = param.AssetId
	confirm.Operation = param.Operation
	confirm.Operator = param.Operator
	confirm.Opinion = param.Opinion
	confirm.OperateTime = time.Now().Unix()
	confirm.ChainUser = common.GetUserFromCertification(stub)

	asset.Status = common.STATUS["Confirming"]
	asset.Confirm = append(asset.Confirm, confirm)

	jsonByte, err := json.Marshal(asset)
	if err != nil {
		log.Logger.Error("Register -- marshal product err:" + err.Error() + "	assetid:" + param.AssetId)
		tChan.AssetId = param.AssetId
		tChan.Status = false
		tChan.Error = err.Error()
		return
	}

	err = stub.PutState(common.ASSET_INFO+common.ULINE+param.AssetId, jsonByte)
	if err != nil {
		log.Logger.Error("Confirm -- putState:" + err.Error() + "	assetid:" + param.AssetId)
		tChan.AssetId = param.AssetId
		tChan.Status = false
		tChan.Error = err.Error()
		return
	}
	tChan.AssetId = param.AssetId
	tChan.Status = true
	tChan.Error = "完成"
	return
}
