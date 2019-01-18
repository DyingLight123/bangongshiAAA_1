package controllers

import (
	"bangongshiAAA_1/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
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
