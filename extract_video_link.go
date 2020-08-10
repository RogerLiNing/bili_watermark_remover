package bili

import (
	"regexp"
)

func ExtractVideoLink(content string) string {
	//fmt.Println(content)

	//解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile("readyVideoUrl: '//(.*)'")

	//根据规则提取关键信息
	result := reg.FindAllStringSubmatch(content,-1)
	var videoLink string

	if len(result) > 0 {
		videoLink = "https://" + result[0][1]
		return videoLink
	}

	return videoLink
}
