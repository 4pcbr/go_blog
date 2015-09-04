package controllers

import (
    "github.com/coopernurse/gorp"
    "database/sql"
    "github.com/revel/revel"
)

var (
    Dbm *gorp.DbMap
)

type GorpController struct {
    *revel.Controller
    Txn *gorp.Transaction
}

func ( c *GorpController ) Begin() revel.Result {
    if txn, err := Dbm.Begin(); err != nil {
        panic( err )
    } else {
        c.Txn = txn
    }
    return nil
}

func ( c * GorpController ) Commit() revel.Result {
    if c.Txn == nil {
        return nil
    }
    if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
        panic( err )
    }
    c.Txn = nil
    return nil
}

func ( c * GorpController ) Rollback() revel.Result {
    if c.Txn == nil {
        return nil
    }
    if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
        panic( err )
    }
    c.Txn = nil
    return nil
}
