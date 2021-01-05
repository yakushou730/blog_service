package service

//type Article struct {
//	*Model
//	Title         string `json:"title"`
//	Desc          string `json:"desc"`
//	Content       string `json:"content"`
//	CoverImageUrl string `json:"cover_image_url"`
//	State         uint8  `json:"state"`
//}

type CountArticleRequest struct {
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=100"`
	Content       string `form:"content" binding:"max=100"`
	CoverImageUrl string `form:"cover_image_url" binding:"max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagArticleRequest struct {
	Title         string `form:"title" binding:"max=100"`
	Desc          string `form:"desc" binding:"max=100"`
	Content       string `form:"content" binding:"max=100"`
	CoverImageUrl string `form:"cover_image_url" binding:"max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"desc" binding:"required,min=3,max=100"`
	Content       string `form:"content" binding:"required,min=3,max=100"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,min=3,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"desc" binding:"required,min=3,max=100"`
	Content       string `form:"content" binding:"required,min=3,max=100"`
	CoverImageUrl string `form:"cover_image_url" binding:"required,min=3,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
