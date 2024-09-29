package Rule

var JavaRceRuleList = []string{
	"Runtime.getRuntime().exec", "ProcessBuilder.start",
	"RuntimeUtil.exec(", "RuntimeUtil.execForStr(",
}

var PHPRceRuleList = []string{
	"System(", "shell_exec(", "exec(", "eval(", "passthru(", "proc_open(", "popen(",
	"assert(", "call_user_func(", "call_user_func_array(", "create_function(",
}
