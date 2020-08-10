package bili

func WatermarkRemover(url string)(string, error)  {

	html, err := GetVideoHtml(url)

	if err != nil {
		return "", err
	}

	link := ExtractVideoLink(html)

	return link, nil

}

