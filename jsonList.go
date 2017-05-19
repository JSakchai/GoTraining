package main
import (
	"fmt"
	"encoding/json"

)
type listAddress struct {
	Cardid string `json:"cardid"`
	Listaddress []Address `json:"listaddress"`
}
type Address struct {
	firstAdd1  string `json:"firstadd1"`
	firstAdd2  string `json:"firstadd2"`
	firstAdd3  string `json:"firstadd3"`
}



func main(){
	//var err error
	data := "{\"cardid\":\"13133132132\"," +
		"\"listaddress\":[" +
		"{\"firstadd1\":\"333\",\"firstadd2\":\"tungmahamek\",\"firstadd3\":\"bkk\"}," +
		"{\"firstadd1\":\"444\",\"firstadd2\":\"sathon\",\"firstadd3\":\"bkk\"}]}"
	//res:= listAddress{}.Listaddress
	res := listAddress{}
	//res2 := Address{}
	json.Unmarshal([]byte(data),&res)
	//json.Unmarshal([]byte(data),&res2)

	if len(res.Listaddress) >= 0{
		fmt.Println("Have list address:",len(res.Listaddress))

		fmt.Println(string(res.Cardid))
			//res.Listaddress[0].firstAdd1 = "333"
			fmt.Println("Addrass home is :",res.Listaddress[0].firstAdd1,
				res.Listaddress[0].firstAdd2,res.Listaddress[0].firstAdd3)
			fmt.Println("Address office is :",res.Listaddress[1].firstAdd1,
				res.Listaddress[1].firstAdd2,res.Listaddress[1].firstAdd3)

	}else {
		fmt.Println("Can't data from jason ")
	}

}