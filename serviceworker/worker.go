package serviceworker

import (
	"fmt"
	"os"
	"text/template"

	"github.com/akkagao/goutil"
	"github.com/akkagao/nvwa/etc"
)

type NameVo struct {
	Package        string
	ServiceName    string
	ServiceNameLow string
}

func Start(name string) {
	context, err := etc.Asset("template/service/genesisService.tem")
	goutil.ChkErr(err)
	t, err := template.New("service").Parse(string(context))
	goutil.ChkErr(err)
	nameVo := NameVo{
		Package:        goutil.CurrentDirName(),
		ServiceName:    goutil.UpperFirst(name),
		ServiceNameLow: name,
	}
	fileName := name + "Service.go"
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	err = t.Execute(fileObj, nameVo)
	goutil.ChkErr(err)
	fmt.Println("gen service success,fileName:", fileName)
}
