package controllers

// Taken from
// https://rclayton.silvrback.com/revel-gorp-and-mysql

import (
    "github.com/revel/revel"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "strings"
    // "blog/app/models"
)

func init() {
    revel.OnAppStart( InitDB )
    revel.InterceptMethod( (*GorpController).Begin, revel.BEFORE )
    revel.InterceptMethod( (*GorpController).Commit, revel.AFTER )
    revel.InterceptMethod( (*GorpController).Rollback, revel.FINALLY )
}

var InitDB = func() {
    connectionString := getConnectionString();

    if db, err := gorm.Open( "mysql", connectionString ); err != nil {
        revel.ERROR.Fatal(err)
    } else {
        Db = &db
        Db.DB().Ping()
    }
}

// func definePostTable( dbm *gorp.DbMap ) {
//     dbm.AddTable( models.Post{} ).SetKeys( true, "id" )
// }

func getParamString(param string, defaultValue string) string {
    p, found := revel.Config.String(param)
    if !found {
        if defaultValue == "" {
            revel.ERROR.Fatal("Cound not find parameter: " + param)
        } else {
            return defaultValue
        }
    }
    return p
}

func getConnectionString() string {
    host     := getParamString("db.host", "")
    port     := getParamString("db.port", "3306")
    user     := getParamString("db.user", "")
    pass     := getParamString("db.password", "")
    dbname   := getParamString("db.dbname", "WTFFIXME")
    protocol := getParamString("db.protocol", "tcp")
    dbargs   := getParamString("dbargs", " ")

    if strings.Trim(dbargs, " ") != "" {
        dbargs = "?" + dbargs
    } else {
        dbargs = ""
    }
    return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
        user, pass, protocol, host, port, dbname, dbargs)
}
