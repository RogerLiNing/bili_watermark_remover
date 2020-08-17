package bili

import (
	"net/url"
	"strings"
)

func WatermarkRemover(raw_url string)(string, error)  {


	var err error
	var epId string
	var link string
	if strings.Contains(raw_url, "bangumi") {

		// https://m.bilibili.com/bangumi/play/ss28777
		if strings.Contains(raw_url, "play/ss") {
			if strings.Contains(raw_url, "https://m.") {
				raw_url = strings.Replace(raw_url, "https://m.", "https://www.", -1)
			}

			html, err := GetVideoHtml(raw_url, false)

			if err != nil {
				return "", err
			}

			epId = GetEpIdFromHtml(html)
		} else if strings.Contains(raw_url, "play/ep") {
			// https://m.bilibili.com/bangumi/play/ep326633

			epId = strings.Split(raw_url, "https://m.bilibili.com/bangumi/play/ep")[1]
			link, err = getVideoInfo(epId)
		}

		if len(epId) > 0 {
			link, err = getVideoInfo(epId)
		}

		if err != nil {
			return "", err
		}


	} else if strings.Contains(raw_url, "bbq.bilibili.com") {
		// https://bbq.bilibili.com/video/?id=1583402526078074470

		u, err := url.Parse(raw_url)
		if err != nil {
			return "", nil
		}
		vid := u.Query().Get("id")

		link, err = getbbqVideo(vid)
	} else {
		// https://www.bilibili.com/video/BV1p5411879s
		// https://b23.tv/Bisisw
		html, err := GetVideoHtml(raw_url, true)

		if err != nil {
			return "", err
		}
		link = ExtractVideoLink(html)
	}


	return link, nil

}

