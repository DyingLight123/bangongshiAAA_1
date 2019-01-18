package controllers

import (
	"bangongshiAAA_1/models"
	"errors"
	"github.com/astaxie/beego"
	"github.com/robfig/cron"
	"log"
	"time"
)

type RefreshController struct {
	beego.Controller
}

var pause1 = make(chan string, 1)
var status1 int

func (data *RefreshController) Get() {
	go RefreshInfluxdb()

	data.Data["json"] = map[string]interface{}{}
	data.ServeJSON()
}

func (data *RefreshController) Post() {
	err := PauseInfluxdb()
	if err != nil {
		data.Data["json"] = map[string]interface{}{"message": "可以开始刷新了！"}
	} else {
		data.Data["json"] = map[string]interface{}{"message": "刷新停止！"}
	}
	data.ServeJSON()
}

func RefreshInfluxdb() {
	if status1 == 1 {
		log.Println("refreshing! please pause! ")
		return
	}
	status1 = 1
	c := cron.New()
	c.AddFunc("@every "+"30s", func() {
		t := time.Now()
		err := models.AddInfluxdbData()
		if err != nil {
			return
		}
		tt := time.Since(t)
		log.Println("总的时间：", tt)
		log.Println("")
	})
	c.Start()
	<-pause1
	status1 = 0
	log.Println("continue")
	c.Stop()
	//time.AfterFunc(30 * time.Second, c.Stop)
}

func PauseInfluxdb() error {
	if status1 != 1 {
		return errors.New("please begin refresh! ")
	}
	log.Println("pause")
	pause1 <- "continue"
	return nil
}
