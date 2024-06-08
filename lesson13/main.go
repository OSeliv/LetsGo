package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type contract struct { //13.1
	Number   int
	Landlord string
	Tenat    string
}

type contract2 struct { //13.2 + 13.4
	Number   int
	Landlord string
	tenat    string //not in json/xml
}

type contract3 struct { // 13.3
	Number   int    `xml:"number"`
	SignDate string `xml:"sign_date"`
	Landlord string `xml:"landlord"`
	Tenat    string `xml:"tenat"`
}

type contracts struct { //13.3
	List []contract3 `xml:"contract"`
}

type contracts4 struct { //13.4
	List []contract2 `xml:"contract"`
}

func main() {
	fmt.Println("Task 13.1")
	str := `{"number":1,"landlord":"Остап Бендер","tenat":"Шура Балаганов" }`

	c := contract{}

	err := json.Unmarshal([]byte(str), &c)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", c)

	fmt.Println("Task 13.2")
	c2 := contract2{
		Number:   2,
		Landlord: "Остап",
		tenat:    "Шура", //will not in json
	}
	js, err := json.Marshal(c2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%+v\n", string(js))

	fmt.Println("Task 13.3")
	str3 := `
<?xml version="1.0" encoding="UTF-8"?>
<contracts>
<contract>
<number>1</number>
<sign_date>2023-09-02</sign_date>
<landlord>Остап</landlord>
<tenat>Шура</tenat>
</contract>
<contract>
<number>2</number>
<sign_date>2023-09-03</sign_date>
<landlord>Бендер</landlord>
<tenat>Балаганов</tenat>
</contract>
</contracts>`

	data := contracts{}

	err = xml.Unmarshal([]byte(str3), &data)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%+v\n", data)

	fmt.Println("Task 13.4")
	c2 = contract2{
		Number:   1,
		Landlord: "Остап Бендер",
		tenat:    "Шура Балаганов", //will not in xml
	}

	d := contracts4{}
	d.List = append(d.List, c2)
	res, err := xml.Marshal(d)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(xml.Header, string(res))
}
