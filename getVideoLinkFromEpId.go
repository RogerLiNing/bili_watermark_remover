package bili

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type JsonData struct {
	Result Result `json:"result"`
}

type Result struct {
	Durl []Durl `json:"durl"`
}

type Durl struct {
	Url string `json:"url"`
}
/*
{
    "code": 0,
    "message": "success",
    "result": {
        "accept_format": "mp4",
        "code": 0,
        "durl": [
            {
                "size": 45618850,
                "ahead": "",
                "length": 360093,
                "vhead": "",
                "backup_url": [],
                "url": "https://upos-sz-mirrorks3.bilivideo.com/upgcxcode/39/79/214607939/214607939-1-29.mp4?e=ig8euxZM2rNcNbh17WdVhoMzhWUVhwdEto8g5X10ugNcXBMvNC8xNbLEkF6MuwLStj8fqJ0EkX1ftx7Sqr_aio8_&uipk=5&nbs=1&deadline=1597083964&gen=playurl&os=ks3bv&oi=1896675426&trid=018f4de10a7247548ab21d13962de41ep&platform=html5&upsig=2c684941bab61f9d68452ed45a90e23d&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,platform&mid=0&orderid=0,1&logo=80000000",
                "order": 1,
                "md5": ""
            }
        ],
*/


func getVideoInfo(EpId string) (string, error) {

	client := &http.Client{}
	// 通过这个接口获取视频信息，其中包括带有水印的链接
	url := "https://api.bilibili.com/pgc/player/web/playurl/html5?bsource=&ep_id=" + EpId

	request, err := http.NewRequest("GET", url, nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Content-type", "application/json")

	response, err := client.Do(request)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)

	}

	jsonData := &JsonData{}
	err = json.Unmarshal(body, jsonData)

	if err != nil {
		fmt.Println(err)
	}

	var videolInk string

	if len(jsonData.Result.Durl) > 0 {
		videolInk = jsonData.Result.Durl[0].Url
	}

	return videolInk, nil

}
