package bili

import (
	"regexp"
)

type VideoJson struct {
	Id int `json:"id"`
}

func GetEpIdFromHtml(content string) string {
	//解析正则表达式，如果成功返回解释器
	// reg := regexp.MustCompile(`"epInfo":(.*),"epList`)
	// fmt.Println(content)
	reg := regexp.MustCompile(`"epList":\[{"loaded":false,"id":(.*?),"badge":""`)
	//根据规则提取关键信息
	result := reg.FindAllStringSubmatch(content,-1)
	//var videoJsonStr string

	// v := &VideoJson{}
	videoId := ""
	if len(result) > 0 {
		videoId = result[0][1]

		//err := json.Unmarshal([]byte(videoJsonStr), v)
		//if err != nil {
		//	fmt.Println(err)
		//}
	}

	return videoId //strconv.Itoa(v.Id)
}
