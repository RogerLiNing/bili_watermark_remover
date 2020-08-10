package bili

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

type VideoJson struct {
	Id int `json:"id"`
}

func GetEpIdFromHtml(content string) string {
	//解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile(`"epInfo":(.*),"epList`)

	//根据规则提取关键信息
	result := reg.FindAllStringSubmatch(content,-1)
	var videoJsonStr string

	v := &VideoJson{}

	if len(result) > 0 {
		videoJsonStr = result[0][1]

		err := json.Unmarshal([]byte(videoJsonStr), v)
		if err != nil {
			fmt.Println(err)
		}
	}

	return strconv.Itoa(v.Id)
}
