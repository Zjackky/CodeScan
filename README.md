
# USAGE

```bash
NAME:
   CodeScan -> A Quick Code Scan Tool

USAGE:
   CodeScan -L language -d directory [-pb pathBlackRule] [-lb lineBlackRule] [-u uploadRule] [-r rceRule]

VERSION:
   1.3

Author:
   Zjacky
   xiaoqiuxx

OPTIONS:
   --Language value, -L value    Code Type (such as : java, php, net)
   --directory path, -d path     You Want To Scan (such as : ./com)
  
Example:
	CodeScan_windows_amd64.exe -L java -d ./net
	CodeScan_windows_amd64.exe -L php -d ./net
	
	
Tips:
	这里进行了忽略排查--->不检查常见的依赖所反编译出来的包: 
给出黑名单列表

var PathBlackJava = []string{
	"apache", "lombok", "microsoft", "solr",
	"amazonaws", "c3p0", "jodd", "afterturn", "hutool",
	"javassist", "alibaba", "aliyuncs", "javax", "jackson",
	"bytebuddy", "baomidou", "google", "netty", "redis", "mysql",
	"logback", "ognl", "oracle", "sun", "junit", "reactor", "github",
	"mchange", "taobao", "nimbusds", "opensymphony", "freemarker", "java", "apiguardian", "hibernate", "javassist", "jboss", "junit", "mybatis",
	"springframework", "slf4j",
}

常见目录为

└─main
    ├─java
    │  └─net
    │      └─mingsoft
    │          ├─cms
    │          │  ├─action
    │          │  │  └─web
    │          │  ├─aop
    │          │  ├─bean
    │          │  ├─biz
    │          │  │  └─impl
    │          │  ├─constant
    │          │  │  └─e
    │          │  ├─dao
    │          │  ├─entity
    │          │  ├─resources
    │          │  └─util
    │          └─config
    
    
    
请把CodeScan放在Net同级目录下扫描(否则会忽略掉直接一个Java目录)
```

# 功能
1. 快速框架分析
2. 涵盖大部分漏洞的Sink点的匹配(如图)
3. 可自定义定制化修改黑白名单内容
4. 多模块化多语言化代码审计


# Tips
1. 存在标准开发代码以及反编译代码的扫描问题，请-d后面的参数尽量在/src/main/java之后(相对路径比较妥当)
2. 扫描的目录一般为`WEB-INF`下的内容(里面包括了反编译的内容)



# 想法
1. 根据Jar进行分析(默认分析)

+ mysqlconnect-->jdbc
+ Xstream --> xml/json

tips: 只需要在CodeScan的目录下放入EvilJarList.txt即可匹配出来

2. 进行融于鉴权代码的快速匹配抓取

例子：
比如现在有一个代码百分百为鉴权代码
```java
<%@ include file="../../common/js/CheckSession.jsp"%>
```

此时可以用一下功能来进行快速获取未鉴权代码

```bash
CodeScan_windows_amd64.exe -d ./yuan -m "<%@ include file="../../common/js/CheckSession.jsp"%>"
<%@ include file="../../common/js/CheckSession.jsp"%>
<%@ include file="../../../common/jsp/CheckSession.jsp"%>
```
然后可以再扫一遍就可以立刻定位到存在未鉴权并且存在Sink点的函数文件了
```bash
CodeScan_windows_amd64.exe -L java -d ./NoAuthDir
```
