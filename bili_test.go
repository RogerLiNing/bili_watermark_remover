package bili

import (
	"fmt"
	"testing"
)

func TestAvailableVideo(t *testing.T) {

	url := "https://m.bilibili.com/bangumi/play/ss28777" // https://www.bilibili.com/video/BV1p5411879s //https://m.bilibili.com/bangumi/play/ss28777
	t.Log("测试有效视频短链接：" + url)

	u, err := WatermarkRemover(url)

	fmt.Println(u)


	if err != nil {
		t.Fail()
		t.Log(err)
	}

}
