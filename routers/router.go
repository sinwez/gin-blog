package routers

import (
	_ "gin-blog/docs"
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/qrcode"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/upload"
	"gin-blog/routers/api"
	v1 "gin-blog/routers/api/v1"
	"net/http"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)
	//r.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})
	apiv1 := r.Group("api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("tags", v1.GetTags)
		//新建标签
		apiv1.POST("tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("tags/:id", v1.DeleteTag)
		//导出标签
		r.POST("tags/export", v1.ExportTag)
		//导入标签
		r.POST("tags/import", v1.ImportTag)

		//获取文章列表
		apiv1.GET("articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("articles/:id", v1.DeleteArticle)
		//二维码
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}
	return r
}
