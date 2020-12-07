package main

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/PuerkitoBio/goquery"

)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetID(htmlresponse *grequests.Response){
	
	doc, err := goquery.NewDocumentFromReader(htmlresponse)
	checkerr(err)

	doc.Find("table").Each(func(i int, a *goquery.Selection){
		number := a.Find("span#MainContent_labelNIN").Text()
		fmt.Println(number)
	})
	

}


func main() {
	first_name := ""
	last_name := ""
	day := ""
	month := ""
	year := ""
	mom_first_name := ""
	mom_last_name := ""
	url := "https://services.nida.go.tz/nidportal/get_nin.aspx/"
	//data to post to the website is placed here
	data := map[string]string{
		"__VIEWSTATE":                 "/wEPDwULLTE2MTE2NjcxNjAPZBYCZg9kFgICAw9kFgICBw9kFgICAQ9kFgYCCg8PFgIeB1Zpc2libGVnZGQCCw8PFggeCUZvbnRfU2l6ZSgqIlN5c3RlbS5XZWIuVUkuV2ViQ29udHJvbHMuRm9udFVuaXQEMjBwdB4JRm9yZUNvbG9yCmweBFRleHQFFDE5OTkxMjIwMTQxMjgwMDAwMTIwHgRfIVNCAoQIZGQCDA8PFgQfAGceC05hdmlnYXRlVXJsBQ1OSURfQ29weS5hc3B4ZGRkcAf4q10sAzfgJ25L45umLcGVWJ3DjvWfE8EtfUfret0=",
		"__EVENTVALIDATION":           "/wEWBgLvo/emCQKk0qrUBQKK8u7bBAL66eaoDgKS3rH/DgLZ662hCx2C1/SEFPdGgJUuhr+WioED7vX6v/clvv59lA9f//3r",
		"ctl00$MainContent$FIRSTNAME": first_name,
		"ctl00$MainContent$SURNAME":   last_name,
		"DAY":   day,
		"MONTH": month,
		"YEAR":  year,
		"ctl00$MainContent$MOTHERSFIRSTNAME": mom_first_name,
		"ctl00$MainContent$MOTHERSSURNAME":   mom_last_name,
		"ctl00$MainContent$btn":              "Tafuta",
	}

	resp, err := grequests.Post(url, &grequests.RequestOptions{Data: data})
	checkerr(err)
	GetID(resp)
	//data type
	// x := fmt.Sprintf("%T", resp)
	// fmt.Println(x)
}
	
	
