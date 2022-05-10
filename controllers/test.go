/**
syk 练手项目
 */
package controllers

import (
	"fmt"
	"time"

	"github.com/george518/PPGo_ApiAdmin/models"
)

type TestController struct {
	BaseController
}

func (self *TestController) List() {
	self.Data["pageTitle"] = "API文档"
	self.Data["ts"] = time.Now()

	grouplists := groupLists()
	self.Data["grouplists"] = grouplists

	groupId, _ := self.GetInt("id", 1)
	groupInfo, err := models.GroupGetById(groupId)

	if err != nil {
		fmt.Println("数据不存在")
	}

	//公共文档
	apiPublic, err := models.ApiPublicGetByIds(groupInfo.ApiPublicIds)
	self.Data["apiPublic"] = apiPublic

	//环境
	// env, err := models.EnvGetByIds(groupInfo.EnvIds)
	// self.Data["env"] = env

	// //状态码
	// code, err := models.CodeGetByIds(groupInfo.CodeIds)
	// self.Data["code"] = code

	//接口
	apiMenu, _ := models.ApiTreeData(groupId)
	self.Data["apiMenu"] = apiMenu
	self.Data["groupId"] = groupId

	self.TplName = "apidoc/index.html"
}

