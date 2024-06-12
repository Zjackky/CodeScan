package Rule

var XmlSqlBlack = []string{
	"<property", "<value>", "id=\"dataSource\"", "<int",
	"<str", "<bool", "<param-value>", "<import", "<delete", "classpath=",
	"<pathelement", "<javac ", "<fileset", "<fail", "<version", "<directory>",
	"<resultMap", "<resultType", "<file", "<mvc:", "<prop", "<param", "<result",
}

var XmlBlack = []string{
	//sql检测不匹配 框架检测也不匹配
	"pom.xml", "log4j2.xml",
}
