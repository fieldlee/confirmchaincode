package control

import (
	"confirmchaincode/log"
	"confirmchaincode/module"
	"confirmchaincode/services"
	"encoding/json"
	// "jiakechaincode/service"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func (t *ProductTrace) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	log.Logger.Info("Invoke")
	funcation, args := stub.GetFunctionAndParameters()
	lowFuncation := strings.ToLower(funcation)
	log.Logger.Info("lowFuncation", lowFuncation)
	if lowFuncation == "register" { // 资产上链
		return t.Register(stub, args)
	}
	if lowFuncation == "confirm" { // 确权资产上链
		return t.Confirm(stub, args)
	}
	return shim.Error("Invalid invoke function name. " + funcation)
}

/** 资产上链 **/
func (t *ProductTrace) Register(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Register接口开始###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var assetRegister module.RegitserParam
		err := json.Unmarshal([]byte(args[0]), &assetRegister)
		if err != nil {
			log.Logger.Error("Register:err" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return services.toRegister(stub, assetRegister)
		}
	} else {
		log.Logger.Error("Register:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}

/** 确权资产上链 **/
func (t *ProductTrace) Confirm(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	log.Logger.Info("##############调用Confirm接口开始###############")
	returnInfo := module.ReturnInfo{}
	if len(args) >= 1 {
		var confirmParam module.ConfirmParam
		err := json.Unmarshal([]byte(args[0]), &confirmParam)
		if err != nil {
			log.Logger.Error("Confirm:err" + err.Error())
			returnInfo.Success = false
			returnInfo.Info = err.Error()
		} else {
			return services.toConfirm(stub, confirmParam)
		}
	} else {
		log.Logger.Error("Confirm:参数不对，请核实参数信息。")
		returnInfo.Success = false
		returnInfo.Info = "参数不对，请核实参数信息"
	}
	jsonreturn, err := json.Marshal(returnInfo)
	if err != nil {
		return shim.Error("err:" + err.Error())
	}
	return shim.Success(jsonreturn)
}
