package errorcode

import "github.com/kondroid00/sample-server-2022/package/errorcode"

var (
	NOT_FOUND      = errorcode.New("not_found", "指定されたリソースが見つかりませんでした。")
	INVALID_PARAMS = errorcode.New("invalid_params", "指定のパラメーターが不正です。")
)
