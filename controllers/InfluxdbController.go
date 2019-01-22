package controllers

import (
	"bangongshiAAA_1/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"regexp"
)

type InfluxdbController struct {
	beego.Controller
}

type InputTime struct {
	T1   string            `json:"t1"`
	T2   string            `json:"t2"`
	Tags map[string]string `json:"tags"`
}

func (maps *InfluxdbController) Post() {
	t := new(InputTime)
	json.Unmarshal(maps.Ctx.Input.RequestBody, t)

	//2019-01-18T05:40:00+08:00
	reg, err := regexp.Compile(`\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+08:00`)
	if err != nil {
		log.Println(err)
		maps.Data["json"] = map[string]interface{}{"message": err.Error()}
		maps.ServeJSON()
		return
	}
	reg1 := reg.MatchString(t.T1)
	reg2 := reg.MatchString(t.T2)
	if reg1 == false || reg2 == false {
		maps.Data["json"] = map[string]interface{}{"message": "请输入正确的时间格式(2019-01-18T05:40:00+08:00)"}
		maps.ServeJSON()
		return
	}

	if t.Tags["key"] == "" {
		maps.Data["json"] = map[string]interface{}{"message": "请填写key"}
		maps.ServeJSON()
		return
	}

	result, err := models.GetInfluxdbData(t.T1, t.T2, t.Tags)
	if err != nil {
		log.Println(err)
		maps.Data["json"] = map[string]interface{}{"message": err.Error()}
		maps.ServeJSON()
		return
	}

	list := make([]map[string]interface{}, 0)
	for i := 0; i < len(result); i++ {
		m := make(map[string]interface{})
		j, _ := json.Marshal(&result[i])
		json.Unmarshal(j, &m)
		list = append(list, m)
	}
	maps.Data["json"] = map[string]interface{}{"data": list, "length": len(list)}
	maps.ServeJSON()
}
