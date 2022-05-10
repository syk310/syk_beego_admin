package main

import (
	"time"

	"github.com/syk310/syk_beego_admin/models"
	_ "github.com/syk310/syk_beego_admin/routers"

	"github.com/astaxie/beego"
	"github.com/syk310/syk_beego_admin/utils"
	cache "github.com/patrickmn/go-cache"
)

func main() {
	models.Init()
	utils.Che = cache.New(60*time.Minute, 120*time.Minute)
	beego.Run()
}
