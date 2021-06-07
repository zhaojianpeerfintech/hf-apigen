package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

	"github.com/DeanThompson/ginpprof"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/peerfintech/hf-apigen/apigen/router"
)


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

//func main(){
	//buffer := new(bytes.Buffer)
	//
	//funcs := map[string]interface{}{
	//	//"bindtype":      bindType[lang],
	//	//"bindtopictype": bindTopicType[lang],
	//	//"namedtype":     namedType[lang],
	//	//"capitalise":    capitalise,
	//	//"decapitalise":  decapitalise,
	//}
	//data := &tmplData{
	//	Name:   "TestData",
	//	Fields: []*tmplField{
	//		&tmplField{"string","F1"},
	//		&tmplField{"string","F2"},
	//		&tmplField{"string","F3"},
	//	},
	//	Key:"F1",
	//}
	//tmpl := template.Must(template.New("").Funcs(funcs).Parse(tmplSourceGo))
	//if err := tmpl.Execute(buffer, data); err != nil {
	//	panic(err)
	//	return
	//}
	//code, err := format.Source(buffer.Bytes())
	//if err != nil {
	//	panic(err)
	//	return
	//}
	////fmt.Println(string(code))
	//err = utils.DecompressTGZ("../apiserver.tar.gz","./projectName")
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//f, err := os.OpenFile("./projectName/handler/invoke.go", os.O_CREATE|os.O_WRONLY, 0666) //打开文件
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//defer f.Close()
	//_, err = f.Write(code)
	//if err != nil {
	//	panic(err)
	//	return
	//}

//	makeImage()
//}

var logger = flogging.MustGetLogger("service")
var port = flag.Int("port", 80, "listen port")
func main() {
	var err error
	runtime.GOMAXPROCS(runtime.NumCPU())

	gin.SetMode(gin.ReleaseMode)
	r := router.GetRouter()
	ginpprof.Wrapper(r) // for debug

	logger.Debug("The listen port is ", port)
	server := endless.NewServer(fmt.Sprintf(":%d", *port), r)

	// save pid file
	server.BeforeBegin = func(add string) {
		pid := syscall.Getpid()
		logger.Criticalf("Actual pid is %d", pid)
		pidFile := "apiserver.pid"
		if checkFileIsExist(pidFile) {
			os.Remove(pidFile)
		}
		if err := ioutil.WriteFile(pidFile, []byte(fmt.Sprintf("%d", pid)), 0666); err != nil {
			logger.Fatalf("Api server write pid file failed! err:%v", err)
		}
	}

	err = server.ListenAndServe()
	if err != nil {
		if strings.Contains(err.Error(), "use of closed network connection") {
			logger.Errorf("%v", err)
		} else {
			logger.Errorf("Api server start failed! err:%v", err)
			panic(err)
		}
		panic(err)
	}
	panic(err)
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}


func makeImage() {
	checkStatement := fmt.Sprintf(`cd ./projectName && make apiserver-docker ; echo $?`)
	output, err := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(output))
}
