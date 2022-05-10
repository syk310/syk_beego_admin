/**
syk 练手项目
 */
package controllers

type TestController struct {
	BaseController
}

func (self *TestController) List() {
	self.Data["pageTitle"] = "API接口"
	self.Data["ApiCss"] = true
	self.Data["auditStatus"] = AUDIT_STATUS_TEXT
	self.display()
}

