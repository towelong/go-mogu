# 蘑菇钉自动打卡
![Go](https://github.com/ToWeLong/go-mogu/workflows/Go/badge.svg)
![schedule-start](https://github.com/ToWeLong/go-mogu/workflows/schedule-start/badge.svg)
![schedule-end](https://github.com/ToWeLong/go-mogu/workflows/schedule-end/badge.svg)
[![GitHub language](https://img.shields.io/badge/language-golang-orange.svg)](https://golang.org/)
[![GitHub license](https://img.shields.io/github/license/ToWeLong/zhihu-hot-questions)](https://github.com/ToWeLong/go-mogu/blob/main/LICENSE)
> 用法：首先找到`.env.example`文件，打开并修改其中的隐私信息，去掉后缀`.example`，存储为`.env`文件即可。

> 重要： 由于蘑菇钉最近更新了接口，导致之前的接口都不可用了，然后还新增了一个sign，也就是签名。
> 非常感谢 `@晓宇` 同学提供的sign算法

    
# 完成
- [X] 自动签到
- [X] 周报自动编写
- [X] 日志记录


# 测试单个函数(去除cached)
```bash
go test -count=1 -v -run ^TestWeek$ towelong/mogu/test
```
