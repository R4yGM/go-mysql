package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
type Data struct {
    ID   int    `json:"value1"`
    IP string `json:"value2"`
    Date string `json:"value3"`
    Time string `json:"value4"`
}
func main() {
    fmt.Println("Go Connection")
    
    db, err := sql.Open("mysql", "root:password.@tcp(127.0.0.1:3306)/db")
fmt.Println("Success Connection")
        



       if err != nil {
        panic(err.Error())
    }
defer db.Close()

var aa int = 0;
var p Data
results, err := db.Query("SELECT * FROM (     SELECT * FROM table ORDER BY id DESC LIMIT 3000 ) sub ORDER BY id ASC")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

for results.Next() {
err = results.Scan(&p.value1, &p.value2, &p.value3, &p.value4)

aa = aa +1;
fmt.Println("-------------+",aa,"+----------\n ","value1: ",p.value1, "\n value2: "+p.value2+ "\n value3: "+p.value3 +"\n value4: "+p.value4)
}

    


}

