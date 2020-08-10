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

		// https://m.bilibili.com/bangumi/play/ss28777
		if strings.Contains(url, "play/ss") {
			epId = GetEpIdFromHtml(html)
		} else if strings.Contains(url, "play/ep") {
			// https://m.bilibili.com/bangumi/play/ep326633

			epId = strings.Split(url, "https://m.bilibili.com/bangumi/play/ep")[1]
			link, err = getVideoInfo(epId)
		}

		if len(epId) > 0 {
			link, err = getVideoInfo(epId)
		}

		if err != nil {
			return "", err
		}

	// https://www.bilibili.com/video/BV1p5411879s
	} else if strings.Contains(url, "video") {
		link = ExtractVideoLink(html)
	}


	return link, nil

}

