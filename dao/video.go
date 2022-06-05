package dao

type Video struct {
	VideoId       int64    `json:"id,omitempty"`
	Author        User     `json:"author" gorm:"foreignKey:UserId;references:UserId;"`
	PlayUrl       string   `json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	FavoriteCount int64    `json:"favorite_count" gorm:"force:force"`
	CommentCount  int64    `json:"comment_count" gorm:"force:force"`
	Favorite      Favorite ` gorm:"foreignkey:UserId;references:UserId;"`
	IsFavorite    bool     `json:"is_favorite"`
	UserId        int64    `gorm:"not null"`
	Titile        string   `json:"title,omitempty"`
	CreateDate    int64    `gorm:"autoCreateTime"`
}

func newvideo(param *Video) Video {
	video := Video{
		VideoId:       param.VideoId,
		Author:        param.Author,
		PlayUrl:       param.PlayUrl,
		CoverUrl:      param.CoverUrl,
		FavoriteCount: param.FavoriteCount,
		CommentCount:  param.CommentCount,
		Favorite:      param.Favorite,
		IsFavorite:    param.IsFavorite,
		UserId:        param.UserId,
		Titile:        param.Titile,
		CreateDate:    param.CreateDate,
	}
	return user
}
func GetVideoList() []Video {
	videos := []Video{}
	db.Debug().Preload("Author").Order("create_date desc").Find(&videos)

	return videos
}

func VideoFavorite(userid int64) []int64 {
	favorites := []int64{}
	db.Debug().Table("favorites").Select("video_id").Where("user_id = ?", userid).Scan(&favorites)
	return favorites
}

func Videofollow(userid int64) []int64 {
	follow := []int64{}
	db.Debug().Table("follows").Select("follow_id").Where("follower_id = ?", userid).Scan(&follow)
	return follow
}

func PublishList(userid int64) []Video {
	publishlist := []Video{}
	db.Debug().Preload("Author").Where("user_id =?", userid).Find(&publishlist)
	return publishlist
}