package controllers

import (
    // "blog/app/models"
    "github.com/revel/revel"
    "blog/app/models"
)

type PostController struct {
    GormController
}

func ( c PostController ) List() revel.Result {
    var posts []models.Post
    Db.Where( "published = 1" ).Order( "created_at desc" ).Limit(5).Find(&posts)
    return c.Render(posts) // TODO
}

func ( c PostController ) Show ( post_token string ) ( revel.Result ) {

    //post := models.Post{}
    //Txn.Where( "token = ?", post_token ).First( &post )
    return c.Render( nil )
}
