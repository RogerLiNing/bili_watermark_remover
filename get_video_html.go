package bilibili

import (
	"io/ioutil"
	"log"
	"net/http"
)


func GetVideoHtml(url string)(string, error)  {

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	response, _ := client.Do(request)

	body, err := ioutil.ReadAll(response.Body)


	if err != nil {
		log.Fatal(err)
		return "", nil
	}

	defer response.Body.Close()

	return string(body), nil
}

