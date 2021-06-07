package handler

type tmplData struct{
	//Package   string
	//Structs   map[string]*tmplStruct
	Name   string       // Auto-generated struct name(before solidity v0.5.11) or raw name.
	Fields []*tmplField
	Key    string
}

type tmplStruct struct {
	Name   string       // Auto-generated struct name(before solidity v0.5.11) or raw name.
	Fields []*tmplField // Struct fields definition depends on the binding language.
}

// tmplField is a wrapper around a struct field with binding language
// struct type definition and relative filed name.
type tmplField struct {
	Type    string   // Field type representation depends on target binding language
	Name    string   // Field name converted from the raw user-defined field name
}

const tmplSourceGo = `
// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/peerfintech/hf-apigen/apiserver/common"
)

type {{.Name}} struct {
	{{range $field := .Fields}}
	{{$field.Name}} {{$field.Type}}{{end}}
}

func InvokeData(c *gin.Context) {
	info := {{.Name}}{}
	err := c.ShouldBindJSON(&info)
	if err != nil {
		logger.Errorf("Read request failed %s", err.Error())
		Response(c, err, common.RequestFormatErr, nil)
		return
	}
    req := &Request{
		Key: info.{{.Key}},
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

`