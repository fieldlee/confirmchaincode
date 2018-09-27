package services

import (
	"confirmchaincode/common"
	"confirmchaincode/log"
	"confirmchaincode/module"
	"encoding/json"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func ToRegister(stub shim.ChaincodeStubInterface, param module.RegitserParam) (tChan module.ChanInfo) {
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
	asset.Operation = param.Operation
	asset.Operator = param.Operator
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

	loged := TransferLog(stub, param.AssetId, param, nil)

	if loged == false {
		log.Logger.Error("Register -- 操作日志保存错误" + "	assetid:" + param.AssetId)
	}

	tChan.AssetId = param.AssetId
	tChan.Status = true
	tChan.Error = "完成"
	return
}

func ToConfirm(stub shim.ChaincodeStubInterface, param module.ConfirmParam) (tChan module.ChanInfo) {
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
	confirm.Signature = param.Signature

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

	loged := TransferLog(stub, param.AssetId, nil, param)

	if loged == false {
		log.Logger.Error("Confirm -- 操作日志保存错误" + "	assetid:" + param.AssetId)
	}

	tChan.AssetId = param.AssetId
	tChan.Status = true
	tChan.Error = "完成"
	return
}

/** 记录日志 **/
func TransferLog(stub shim.ChaincodeStubInterface, assetid string, registerparam module.RegitserParam, confimparam module.ConfirmParam) bool {
	curuser := common.GetUserFromCertification(stub)
	tran := module.Transfer{}
	tran.TxHash = stub.GetTxID()
	tran.OperateTime = time.Now().Unix()
	tran.Operation = operation
	tran.Operator = curuser
	tran.Confirm = confimparam
	tran.Register = registerparam

	jsonByte, err := json.Marshal(tran)
	if err != nil {
		log.Logger.Error("TransferLog --err:" + err.Error())
		return false
	}
	err = stub.PutState(common.ASSET_ACTION+common.ULINE+assetid, jsonByte)
	if err != nil {
		log.Logger.Error("TransferLog -- putState:" + err.Error())
		return false
	}
	return true
}

/** 查找溯源 **/
func QueryHistory(stub shim.ChaincodeStubInterface, param module.QueryParam) (tChan module.QueryLog) {
	resultsIterator, err := stub.GetHistoryForKey(common.ASSET_ACTION + common.ULINE + param.AssetId)
	if err != nil {
		tChan.Info = err.Error()
		tChan.Success = false
		return
	}
	defer resultsIterator.Close()

	results := make([]module.Transfer, 0)
	for resultsIterator.HasNext() {
		result, err := resultsIterator.Next()
		if err != nil {
			tChan.Info = err.Error()
			tChan.Success = false
			return
		}
		tran := module.Transfer{}
		err = json.Unmarshal(result.Value, &tran)
		if err != nil {
			tChan.Info = err.Error()
			tChan.Success = false
			return
		} else {
			results = append(results, tran)
		}
	}
	tChan.Actions = results
	tChan.Success = true

	return
}
