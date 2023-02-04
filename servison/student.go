package service

import (
	"beegoapi/models"

	"github.com/beego/beego/v2/adapter/logs"
)

func Addstu(v models.Student) error {
	if _, err := models.AddStudent(&v); err == nil {
		logs.Info("success")
		return nil
	} else {
		logs.Info(err)
		return err
	}
}
