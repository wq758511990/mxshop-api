package utils

var (
	OK         = Response(200, "操作成功")
	Err        = Response(500, "Internal Error")
	BadRequest = Response(400, "参数错误")
)
