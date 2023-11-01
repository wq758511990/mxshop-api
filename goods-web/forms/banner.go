package forms

type BannerForm struct {
	ID    int32  `form:"id" json:"id"`
	Image string `form:"image" json:"image" binding:"url"`
	Index int32  `form:"index" json:"index" binding:"required"`
	Url   string `form:"url" json:"url" binding:"url"`
}
