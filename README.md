# CodeScan
![image](https://socialify.git.ci/Zjackky/CodeScan/image?description=1&font=Inter&forks=1&issues=1&language=1&logo=https%3A%2F%2Fzjacky-blog.oss-cn-beijing.aliyuncs.com%2Fblog%2F202401021003754.jpg&name=1&owner=1&pattern=Circuit%20Board&pulls=1&stargazers=1&theme=Light)
## 工具概述

该工具目的为对大多数不完整的代码以及依赖快速进行Sink点匹配来帮助红队完成快速代码审计，开发该工具的初衷是以`Sink`​到`Source`​的思路来开发，为了将所有可疑的Sink点匹配出来并且凭借第六感进行快速漏洞挖掘，并且该工具开发可扩展性强，成本极低，目前工具支持的语言有PHP，Java(JSP)

## 编译

```bash
./build.sh

# 会生成所有版本在releases下
```

## 功能

1. 框架识别
2. 涵盖大部分漏洞的Sink点的匹配(如图)
   ![image](https://zjacky-blog.oss-cn-beijing.aliyuncs.com/image-20240928235812-5wlbnbb.png)
3. 可自定义定制化修改黑白名单内容
4. 多模块化多语言化代码审计
5. 进行融于鉴权代码的快速匹配抓取
6. 根据Jar进行静态分析(默认分析)

* mysqlconnect-->jdbc
* Xstream --> xml/json

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

## 高级用法+案例分析

### 高级用法

`以下均以Java作为示例`​

#### 高扩展性

很简单的自定义，如果需要自定义一些匹配规则，首先可以在这里加入

![image](https://zjacky-blog.oss-cn-beijing.aliyuncs.com/image-20240929002903-ypqa197.png)​


其次如果需要新增漏洞类型，只需要三步(这里以Sql为例)

1. 新建SQL目录
2. 定义一个方法叫 SqlCheck
3. 写一个sqlcheck.txt(生成的文件名) + 你自定义的规则
4. 最后在这里加入包名+方法名即可

![image](https://zjacky-blog.oss-cn-beijing.aliyuncs.com/image-20240929003143-7v37o9w.png)​

```go
package SqlTest

import (
	"CodeScan/FindFile"
	"fmt"
)

func SqlCheck(dir string) {
	FindFile.FindFileByJava(dir, "fastjson.txt", []string{".parseObject("})
	fmt.Println("SqlCheck分析完成")

}

```

#### 扫描位置

在打一些闭源代码的时候经常就一个Jar或者Class，反编译的时候会把依赖进行一起反编译，所以为了避免扫描一些依赖的误报，在工具中自带的黑名单中会过滤掉如下黑名单的包名，需要自定义的时候可自行修改，位置在`CommonVul/Rule/MatchPathRule.go`​

```go
var PathBlackJava = []string{
	"apache", "lombok", "microsoft", "solr",
	"amazonaws", "c3p0", "jodd", "afterturn", "hutool",
	"javassist", "alibaba", "aliyuncs", "javax", "jackson",
	"bytebuddy", "baomidou", "google", "netty", "redis", "mysql",
	"logback", "ognl", "oracle", "sun", "junit", "reactor", "github",
	"mchange", "taobao", "nimbusds", "opensymphony", "freemarker", "java", "apiguardian", "hibernate", "javassist", "jboss", "junit", "mybatis",
	"springframework", "slf4j",
}
```

所以这也导致了一个问题，不能从顶层上直接扫描

![image](https://zjacky-blog.oss-cn-beijing.aliyuncs.com/image-20240929124102-qjfancc.png)

`请把CodeScan放在Net同级目录下扫描(否则会忽略掉直接一个Java目录)`​

请`-d`​后面的参数尽量在`/src/main/java`​之后，比如这里就需要把CodeScan放到`net`​目录下开始扫描

```bash
CodeScan_windows_amd64.exe -L java -d ./net
```

#### 过滤字符串(只写了JSP + PHP)

比如现在有一个代码百分百为鉴权代码在JSP中

```java
<%@ include file="../../common/js/CheckSession.jsp"%>
```

此时可以用一下功能来进行快速获取未鉴权代码

```bash
CodeScan_windows_amd64.exe -d ./yuan -m "CheckSession.jsp"
```

此时会将不存在这个代码的文件都放到`NoAuthDir`​目录中，然后可以再扫一遍就可以立刻定位到存在未鉴权并且存在Sink点的函数文件了

```bash
CodeScan_windows_amd64.exe -L java -d ./NoAuthDir
```

#### 静态分析依赖情况

只需要在CodeScan的目录下放入EvilJarList.txt即可匹配出来

`EvilJarList.txt` 内容为存在可打漏洞的`Jar`,模版如下

```bash
fastjson-1.2.47.jar
resin-4.0.63.jar
jackson-core-2.13.3.jar
c3p0-0.9.5.2.jar
commons-beanutils-1.9.4.jar
commons-beanutils-1.9.3.jar
commons-beanutils-1.9.2.jar
commons-collections-3.2.1.jar
mysql-connector-java-8.0.17.jar
commons-collections4-4.0.jar
shiro-core-1.10.1.jar
aspectjweaver-1.9.5.jar
rome-1.0.jar
xstream-1.4.11.1.jar
sqlite-jdbc-3.8.9.jar
vaadin-server-7.7.14.jar
hessian-4.0.63.jar
```

#### 案例
案例请参考我的博客
```bash
https://zjackky.github.io/post/develop-codescan-zwcz53.html
```

## TODO

* [ ] 将结果从TXT转为Excel
* [ ] Sink点继续完善
* [ ] ASP

## 支持项目

* 如果有师傅发现Bug或者有更好的建议请提issue感谢
* 要是各位师傅通过本人的小工具挖到一些好洞记得回头点点Stars诶

## 免责申明

* 如果您下载、安装、使用、修改本工具及相关代码，即表明您信任本工具
* 在使用本工具时造成对您自己或他人任何形式的损失和伤害，我们不承担任何责任
* 如您在使用本工具的过程中存在任何非法行为，您需自行承担相应后果，我们将不承担任何法律及连带责任
* 请您务必审慎阅读、充分理解各条款内容，特别是免除或者限制责任的条款，并选择接受或不接受
* 除非您已阅读并接受本协议所有条款，否则您无权下载、安装或使用本工具
* 您的下载、安装、使用等行为即视为您已阅读并同意上述协议的约束

## 更新日志

**2024/09/29**

* 开源

**2024/10/7**

* 将扫描结果写入result目录中

## 鸣谢

[xiaoqiuxx(github.com)](https://github.com/xiaoqiuxx)
