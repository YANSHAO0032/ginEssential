package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"jianyi.com/ginessential/common"
)



func main()  {
	db := common.InitDb()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}






