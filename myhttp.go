package myhttp

import (
	"fmt"
	"net/http"
	"github.com/rosssalas/mydb"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler:",mydb.Db)
}

