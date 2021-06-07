package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/peerfintech/hf-apigen/apiserver/common"
)

func InvokeData(c *gin.Context) {
	req := Request{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Read request failed %s", err.Error())
		Response(c, err, common.RequestFormatErr, nil)
		return
	}
	txId, errCode, err := invokeData(req)
	if err != nil {
		logger.Errorf("Invoke data failed %s", err.Error())
		Response(c, err, errCode, nil)
		return
	}
	logger.Infof("Upload data %s success", req.Key)
	Response(c, nil, common.Success, txId)
	return
}
