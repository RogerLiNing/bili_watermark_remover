package bili

import (
	"strings"
)

func WatermarkRemover(url string)(string, error)  {

	html, err := GetVideoHtml(url)

	if err != nil {
		return "", err
	}

	var epId string
	var link string
	if strings.Contains(url, "bangumi") {
		// https://m.bilibili.com/bangumi/play/ep326633?spm_id_from=333.851.b_7265706f7274466972737431.1
		epId = GetEpIdFromHtml(html)
		if len(epId) > 0 {
			link, err = getVideoInfo(epId)
		}

		if err != nil {
			return "", err
		}

	} else if strings.Contains(url, "video") {
		link = ExtractVideoLink(html)
	}


	return link, nil

}

