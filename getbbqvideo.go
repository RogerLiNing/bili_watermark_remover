package bili

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type JsonResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Play Play `json:"play"`
}

type Play struct {
	Url string `json:"url"`
}
/*
"code": 0,
    "message": "0",
    "ttl": 1,
    "data": {
        "play": {
            "svid": 1583402526078074470,
            "expire_time": 1597689914,
            "url": "http://upos-sz-mirrorkodo.bilivideo.com/bbqxcode2/70/44/m200305wssvid1583402526078074470-1-720p-2m-264.mp4?e=ig8euxZM2rNcNbdlhoNvNC8BqJIzNbfq9rVEuxTEnE8L5F6VnEsSTx0vkX8fqJeYTj_lta53NCM=&uipk=5&nbs=1&deadline=1597689914&gen=playurl&os=kodobv&oi=1962730905&trid=c1c8027f93c849e3a621e115565882d8b&platform=html5&upsig=dd37d52fae572d2003f699870f218d04&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,platform&mid=0&orderid=0,2&logo=00000000",
            "current_time": 1597682714,
            "seek_bar": 1

*/


func getbbqVideo(vid string) (string, error) {

	client := &http.Client{}
	// 通过这个接口获取视频信息，其中包括带有水印的链接
	api := "https://bbq.bilibili.com/bbq/app-bbq/sv/detail?svid=" + vid + "&version=1.3.0&platform=h5"

	request, err := http.NewRequest("GET", api, nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Content-type", "application/json")

	response, err := client.Do(request)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)

	}

	jsonData := &JsonResponse{}
	err = json.Unmarshal(body, jsonData)

	if err != nil {
		fmt.Println(err)
	}

	var videoLink string

	if len(jsonData.Data.Play.Url) > 0 {
		videoLink = jsonData.Data.Play.Url
	}

	return videoLink, nil

}
