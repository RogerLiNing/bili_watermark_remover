package bili

import (
	"strings"
)

func WatermarkRemover(url string)(string, error)  {


	var err error
	var epId string
	var link string
	if strings.Contains(url, "bangumi") {

		// https://m.bilibili.com/bangumi/play/ss28777
		if strings.Contains(url, "play/ss") {
			if strings.Contains(url, "https://m.") {
				url = strings.Replace(url, "https://m.", "https://www.", -1)
			}

			html, err := GetVideoHtml(url, false)

			if err != nil {
				return "", err
			}

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
	// https://b23.tv/Bisisw
	} else {
		html, err := GetVideoHtml(url, true)

		if err != nil {
			return "", err
		}
		link = ExtractVideoLink(html)
	}


	return link, nil

}

