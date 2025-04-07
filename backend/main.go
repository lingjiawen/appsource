package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "mangosmithy/internal/app/boot"
	_ "mangosmithy/internal/app/system/packed"
	"mangosmithy/internal/cmd"
	_ "mangosmithy/library/libValidate"
	_ "mangosmithy/task"
)

func main() {
	cmd.Main.Run(gctx.New())
}
