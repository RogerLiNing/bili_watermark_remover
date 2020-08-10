package bili

import (
	"io/ioutil"
	"log"
	"net/http"
)


func GetVideoHtml(url string, isMobile bool)(string, error)  {

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if isMobile {
		request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
		request.Header.Add("Host", "m.bilibili.com")
	} else {
		request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36 QBCore/4.0.1301.400 QQBrowser/9.0.2524.400 Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2875.116 Safari/537.36 NetType/WIFI MicroMessenger/7.0.5 WindowsWechat")
		request.Header.Add("Host", "www.bilibili.com")
		request.Header.Add("Cookie","finger=-1374999553; bsource=wechat; _uuid=1F105130-A882-9731-F196-8540D7729FFB93901infoc; buvid3=BAE6C96E-1484-4984-BB0B-E3DECF0CD893143078infoc; CURRENT_FNVAL=16; sid=jhfnapdk")
		request.Header.Add("Referer", url)
	}

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	response, _ := client.Do(request)

	body, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body))

	if err != nil {
		log.Fatal(err)
		return "", nil
	}

	defer response.Body.Close()

	return string(body), nil
}

