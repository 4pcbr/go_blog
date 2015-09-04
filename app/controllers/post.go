package controllers

import (
    // "blog/app/models"
    "github.com/revel/revel"
)

type PostController struct {
    GorpController
}

func ( c PostController ) Show ( post_token string ) ( revel.Result ) {
    //post := models.Post{}
    //Txn.Where( "token = ?", post_token ).First( &post )
    return c.Render( nil )
}
