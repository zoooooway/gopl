package main

import (
	goods "gopl/ch7/eg7.11"
	"log"
	"net/http"
)

// 增加额外的handler让客户端可以创建，读取，更新和删除数据库记录。
// 例如，一个形如 /update?item=socks&price=6 的请求会更新库存清单里一个货品的价格并且当这个货品不存在或价格无效时返回一个错误值。
// （注意：这个修改会引入变量同时更新的问题）
func main() {
	db := goods.Database{"shoes": 50, "socks": 5}

	// http.HandleFunc("/list", db.List)
	http.HandleFunc("/list", db.Table)
	http.HandleFunc("/price", db.Price)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/create", db.Create)
	http.HandleFunc("/delete", db.Delete)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
