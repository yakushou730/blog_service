package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yakushou730/blog-service/global"
	"github.com/yakushou730/blog-service/internal/service"
	"github.com/yakushou730/blog-service/pkg/app"
	"github.com/yakushou730/blog-service/pkg/convert"
	"github.com/yakushou730/blog-service/pkg/errcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

// @Summary 取得多個文章
// @Produce json
// @Param title query string false "文章名稱" maxlength(100)
// @Param desc query string false "文章敘述" maxlength(100)
// @Param content query string false "文章內容" maxlength(100)
// @Param cover_image_url query string false "封面圖片連結" maxlength(100)
// @Param state query int false "狀態" Enums(0, 1) default(1)
// @Param page query int false "頁碼"
// @Param page_size query int false "每頁數量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountArticle(&service.CountArticleRequest{Title: param.Title, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountArticleFail)
		return
	}

	articles, err := svc.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleListFail)
		return
	}

	response.ToResponseList(articles, totalRows)
	return
}

// @Summary 新增文章
// @Produce json
// @Param title query string false "文章名稱" maxlength(100)
// @Param desc query string false "文章敘述" maxlength(100)
// @Param content query string false "文章內容" maxlength(100)
// @Param cover_image_url query string false "封面圖片連結" maxlength(100)
// @Param state query int false "狀態" Enums(0, 1) default(1)
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary 更新文章
// @Produce json
// @Param id path int true "標籤ID"
// @Param title query string false "文章名稱" maxlength(100)
// @Param desc query string false "文章敘述" maxlength(100)
// @Param content query string false "文章內容" maxlength(100)
// @Param cover_image_url query string false "封面圖片連結" maxlength(100)
// @Param state query int false "狀態" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/article/{id} [put]
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 刪除文章
// @Produce json
// @Param id path int true "標籤ID"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/article/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteArticle err:  %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
