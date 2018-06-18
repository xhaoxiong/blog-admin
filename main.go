package main

import (
	_ "blogadmin/routers"
	"github.com/astaxie/beego"
	"encoding/gob"
	"os"
	"blogadmin/models"
)

func init() {
	initArgs()
	models.Connect()
	gob.Register(models.AdminUser{})
}

func main() {
	beego.Run()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.SyncDB()
			os.Exit(0)
		}
		if v == "-admin" {
			models.Connect()
			models.AddAdmin()
		}
	}
}
