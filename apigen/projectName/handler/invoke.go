// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/peerfintech/hf-apigen/apiserver/common"
)

type TestData struct {
	F1 string
	F2 string
	F3 string
}

func InvokeData(c *gin.Context) {
	info := TestData{}
	err := c.ShouldBindJSON(&info)
	if err != nil {
		logger.Errorf("Read request failed %s", err.Error())
		Response(c, err, common.RequestFormatErr, nil)
		return
	}
	req := &Request{
		Key:   info.F1,
		Value: &info,
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
