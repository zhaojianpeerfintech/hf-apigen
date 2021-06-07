package handler

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/peerfintech/hf-apigen/apiserver/common"

	"github.com/peerfintech/hf-apigen/apiserver/logging"
	"github.com/peerfintech/hf-apigen/apiserver/sdk"
)

var logger = logging.NewLogger("debug", "handler")

type Request struct{
	Key string `json:"key"`
	Value interface{}
}

func invokeData(v interface{}) (string, int, error) {
	args, err := json.Marshal(v)
	if err != nil {
		return "", common.RequestFormatErr, err
	}
	req := []string{common.SaveData, string(args)}
	res, err := sdk.Invoke(req)
	if err != nil {
		return "", common.InvokeErr, err
	}
	return res.TxID, common.Success, nil
}

func QueryData(c *gin.Context) {
	key := c.Query("key")
	args := []string{common.QueryData, key}
	bytes, err := sdk.Query(args)
	if err != nil {
		logger.Errorf("Fabric query data %s failed %s", key, err.Error())
		Response(c, err, common.QueryErr, nil)
		return
	}
	logger.Infof("Get value %s", string(bytes))
	Response(c, nil, common.Success, bytes)
	return
}
