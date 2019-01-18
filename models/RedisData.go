package models

func GetRedisData() ([]interface{}, error) {
	client := ConnRedis()
	defer client.Close()

	field, err := client.HMGet("bangongshi[Cache]", "bangongshidfAAA_1W3P", "bangongshidfAAA_1IA",
		"bangongshidfAAA_1IB", "bangongshidfAAA_1IC", "bangongshidfAAA_1I0", "bangongshidfAAA_1FT",
		"bangongshidfAAA_1F", "bangongshidfAAA_1VAB", "bangongshidfAAA_1VBC", "bangongshidfAAA_1VCA",
		"bangongshidfAAA_1Is", "bangongshidfAAA_1T1", "bangongshidfAAA_1DI12").Result()

	if err != nil {
		return nil, err
	}

	return field, nil
	/*field, err := client.HGetAll("data").Result()
	return field, err*/
}
