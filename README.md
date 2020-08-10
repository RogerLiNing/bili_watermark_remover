# 哔哩哔哩视频水印去除库
本仓库是我在另一个项目中使用的解析库，仅作为库源码使用，所以不提供二进制直接运行。毕竟已经有很多去水印的软件了，也不缺我这个。如果你想要在你的go项目中使用，可以在你的代码中导入该包，并按照以下方式调用。

## 安装
```shell script
go get -u github.com/RogerLiNing/bili_watermark_remover
```

## 使用
```go
import (
	bili "github.com/RogerLiNing/bili_watermark_remover"
)

// url 是原始视频，不包含水印
url, _ := bili.WatermarkRemover("https://www.bilibili.com/video/BV1p5411879s")


fmt.Println(url)
// 

```