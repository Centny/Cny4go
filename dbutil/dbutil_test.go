package dbutil

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/Centny/Cny4go/test"
	"github.com/Centny/TDb"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

type TSt struct {
	Tid    int64     `m2s:"TID"`
	Tname  string    `m2s:"TNAME"`
	Titem  string    `m2s:"TITEM"`
	Tval   string    `m2s:"TVAL"`
	Status string    `m2s:"STATUS"`
	Time   time.Time `m2s:"TIME"`
	Fval   float64   `m2s:"FVAL"`
	Uival  int64     `m2s:"UIVAL"`
	Add1   string    `m2s:"ADD1"`
	Add2   string    `m2s:"Add2"`
}

func TestDbUtil(t *testing.T) {
	db, _ := sql.Open("mysql", test.TDbCon)
	defer db.Close()
	res, err := DbQuery(db, "select * from ttable where tid>?", 1)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(res) < 1 {
		t.Error("not data")
		return
	}
	if len(res[0]) < 1 {
		t.Error("data is empty")
		return
	}
	bys, err := json.Marshal(res)
	fmt.Println(string(bys))
	//
	var mres []TSt
	err = DbQueryS(db, &mres, "select * from ttable where tid>?", 1)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(mres) < 1 {
		t.Error("not data")
		return
	}
	fmt.Println(mres, mres[0].Add1)
	//

	_, err = DbQuery(db, "selectt * from ttable where tid>?", 1)
	if err == nil {
		t.Error("not error")
		return
	}
	_, err = DbQuery(db, "select * from ttable where tid>?", 1, 2)
	if err == nil {
		t.Error("not error")
		return
	}
	err = DbQueryS(nil, nil, "select * from ttable where tid>?", 1)
	if err == nil {
		t.Error("not error")
		return
	}
}
func Map2Val2(columns []string, row map[string]interface{}, dest []driver.Value) {
	for i, c := range columns {
		if v, ok := row[c]; ok {
			switch c {
			case "INT":
				dest[i] = int(v.(float64))
			case "UINT":
				dest[i] = uint32(v.(float64))
			case "FLOAT":
				dest[i] = float32(v.(float64))
			case "SLICE":
				dest[i] = []byte(v.(string))
			case "STRING":
				dest[i] = v.(string)
			case "STRUCT":
				dest[i] = time.Now()
			case "BOOL":
				dest[i] = true
			}
		} else {
			dest[i] = nil
		}
	}
}
func TestDbUtil2(t *testing.T) {
	TDb.Map2Val = Map2Val2
	db, _ := sql.Open("TDb", "td@tdata.json")
	defer db.Close()
	res, err := DbQuery(db, "SELECT * FROM TESTING WHERE INT=? AND STRING=?", 1, "cny")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(res)
}
