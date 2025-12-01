package e621

import "time"

type File struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Extention string `json:"ext"`
	Size      int    `json:"size"`
	Md5       string `json:"md5"`
	Url       string `json:"url"`
}

type Post struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Rating      string    `json:"rating"`
	FavCount    int       `json:"fav_count"`
	Description string    `json:"description"`
	File        *File     `json:"file"`
}

type PostsResponse struct {
	Posts []*Post `json:"posts"`
}
