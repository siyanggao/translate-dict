package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	_ "translate-dict/routers"

	"github.com/astaxie/beego"
	"github.com/zzc-tongji/mydictionary"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	initDict()

	beego.Run()

}

func initDict() {
	currentPath := getCurrentDirectory()
	currentPath = currentPath + string(filepath.Separator) + "document"
	path := []string{currentPath}
	err, msg := mydictionary.Initialize(path)
	if !err {
		log.Fatal(msg)
		return
	}
	err, msg = mydictionary.CheckNetwork()
	if !err {
		log.Fatal(msg)
		return
	}
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
