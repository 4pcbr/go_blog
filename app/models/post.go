package models

import "time"

type Post struct {
    Id          int64       `db:"id" json:"id"`
    Title       string      `db:"title" json:"title"`
    Token       string      `db:"token" json:"token"`
    Body        string      `db:"body" json:"body"`
    BlogId      string      `db:"blog_id" json:"blog_id"`
    Tags        string      `db:"tags" json:"tags"`
    Published   bool        `db:"published" json:"published"`
    CreatedAt   time.Time   `db:"created_at" json:"created_at"`
    UpdatedAt   time.Time   `db:"updated_at" json:"updated_at"`
}
