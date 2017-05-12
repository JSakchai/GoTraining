package main
import(
	"fmt"
	"encoding/json"
	//"bytes"
	//"log"
	//"os"
)

type Info struct {
	Cardid string   `json:"cardid"`
	FirstName string `json:"name"`
	Surname string `json:"surname"`
}
var indexStr = "indexStr"

func main()  {
		str := "{ \"id\": \"0001\", \"type\": \"donut\", \"name\": \"Cake\", \"ppu\": 0.55, \"batters\": { " +
			"\"batter\": [{ \"id\": \"1001\", \"type\": \"Regular\" }, { \"id\": \"1002\", \"type\": " +
			"\"Chocolate\" }] }, \"topping\": [{ \"id\": \"5001\", \"type\": \"None\" }, { \"id\": \"5002\", " +
			"\"type\": \"Glazed\" }, { \"id\": \"5005\", \"type\": \"Sugar\" }, { \"id\": \"5007\", \"type\": " +
			"\"Powdered Sugar\" }] }"

		jsonMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(str), &jsonMap)
		if err != nil {
			panic(err)
		}

		batters := jsonMap["batters"].(map[string]interface{})["batter"]

		for _, b := range(batters.([]interface{})) {
			b.(map[string]interface{})["id"] = "other id"
		}

		str2, err := json.Marshal(jsonMap)
		fmt.Println(string(str2))


}
/*
func updateData() {
	info := Info{}
	var infoString []byte
	json.Unmarshal(info,&infoString)

	info.Cardid ="131314546454312"
	info.FirstName = "name0"
	info.Surname = "surname0"

	*/
/*info[1].Cardid ="1215454554541351"
	info[1].FirstName = "name1"
	info[1].Surname = "surname1"
*//*

	toJson,err := json.NewEncoder(info)
	if err != nil{
		fmt.Println("error on bytes to json ")
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out,toJson, "", "\t")
	out.WriteTo(os.Stdout)

}*/
