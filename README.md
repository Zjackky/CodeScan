# README

‍

# CodeScan

## 工具概述

该工具目的为对大多数不完整的代码以及依赖快速进行Sink点匹配来帮助红队完成快速代码审计，开发该工具的初衷是以`Sink`​到`Source`​的思路来开发，为了将所有可疑的Sink点匹配出来并且凭借第六感进行快速漏洞挖掘，并且该工具开发可扩展性强，成本极低，目前工具支持的语言有PHP，Java(JSP)

‍

## 编译

```bash
./build.sh

# 会生成所有版本在releases下
```

‍

## 功能

1. 框架识别
2. 涵盖大部分漏洞的Sink点的匹配(如图)

    ​![image](https://zjacky-blog.oss-cn-beijing.aliyuncs.com/image-20240928235812-5wlbnbb.png)​
3. 可自定义定制化修改黑白名单内容
4. 多模块化多语言化代码审计
5. 进行融于鉴权代码的快速匹配抓取
6. 根据Jar进行静态分析(默认分析)

* mysqlconnect-->jdbc
* Xstream --> xml/json

‍

## 使用

```bash
Usage of ./CodeScan_darwin_arm64:
  -L string
        审计语言
  -d string
        要扫描的目录
  -h string
        使用帮助
  -lb string
        行黑名单
  -m string
        过滤的字符串
  -pb string
        路径黑名单
  -r string
        RCE规则
  -u string
        文件上传规则


Example:
	CodeScan_windows_amd64.exe -L java -d ./net
	CodeScan_windows_amd64.exe -L php -d ./net
	CodeScan_windows_amd64.exe -d ./net -m "CheckSession.jsp"
```

‍

## TODO

* [ ] 将结果从TXT转为Excel
* [ ] Sink点继续完善
* [ ] ASP

## 支持项目

* 如果有师傅发现Bug或者有更好的建议请提issue感谢
* 要是各位师傅通过本人的小工具挖到一些好洞记得回头点点Stars诶

‍

## 高级用法+案例分析

高级用法请参考我的博客

‍
