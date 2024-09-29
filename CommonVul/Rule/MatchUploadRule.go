package Rule

var JavaUploadRuleList = []string{
	"Streams.copy(",
	".getOriginalFilename(", ".transferTo(",
	"UploadedFile(", "FileUtils.copyFile(", "MultipartHttpServletRequest", ".getFileName(", ".saveAs(", ".getFileSuffix(", ".getFile", "MultipartFile file",
}

var PHPUploadRuleList = []string{
	"move_uploaded_file(", "file_put_contents(", "$_FILE[", "copy(", "->move(", "request()->file(",
}
