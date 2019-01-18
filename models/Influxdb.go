package models

import (
	"encoding/json"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"strings"
	"time"
)

type CurData struct {
	Value   string `json:"value"`
	Time    string `json:"time"`
	Quality string `json:"quality"`
}

type BanGongShiDFAAA1 struct {
	Uuid    string `json:"uuid"`
	CurData `json:"curData"`
	Type    int `json:"type"`
}

//influxdb数据库连接
func ConnInfluxdb() client.Client {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://127.0.0.1:8086",
		Username: "admin",
		Password: "admin",
	})
	if err != nil {
		log.Fatal("influxdb数据库连接错误： ", err)
	}
	return cli
}

//influxdb查询
func QueryDB(cli client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: "test",
	}
	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			return nil, response.Error()
		}
		res = response.Results
	} else {
		return nil, err
	}
	return res, nil
}

//influxdb写入
func WritesPoints(cli client.Client, field []interface{}) error {
	//t := time.Now()
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "test",
		Precision: "s",
	})
	if err != nil {
		return err
	}
	for i := 0; i < len(field); i++ {
		m := new(BanGongShiDFAAA1)
		err := json.Unmarshal([]byte(field[i].(string)), m)

		tags := make(map[string]string)
		n := strings.Split(m.Uuid, ",")
		tags["key"] = n[0] + n[1] + n[2] + n[3]

		fields := make(map[string]interface{})
		fields["value"] = m.Value
		fields["quality"] = m.Quality

		pt, err := client.NewPoint(
			"bangongshi",
			tags,
			fields,
			time.Now(),
		)
		if err != nil {
			return err
		}
		bp.AddPoint(pt)
		/*tags := map[string]string{"key": key}
		fields := map[string]interface{}{
			"value": value,
		}
		pt, err := client.NewPoint(
			"map",
			tags,
			fields,
			time.Now(),
		)
		if err != nil {
			return err
		}
		bp.AddPoint(pt)*/
	}
	if err := cli.Write(bp); err != nil {
		return err
	}
	//elapsed := time.Since(t)
	//fmt.Println(elapsed)
	return nil
}
