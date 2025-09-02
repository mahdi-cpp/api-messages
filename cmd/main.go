package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	main2 "github.com/mahdi-cpp/api-messages/cmd"
	"github.com/mahdi-cpp/api-messages/internal/api/handler"
)

func main() {

	fmt.Println("دو ویژگی اصلی یعنی سرعت و دوام، مستقیماً روی تجربه شما تأثیر می‌گذارند. فلشی که سرعت خواندن و نوشتن پایینی دارد، حتی اگر ظرفیت بالایی داشته باشد، در استفاده روزمره شما را کلافه می‌کند. از طرف دیگر، فلشی که ساختار و برند معتبر ندارد، ممکن است خیلی زود خراب شود یا داده‌هایتان را از دست بدهید.\n\nپس قبل از خرید، باید با برندهای معتبر، ویژگی‌های مهم و نیازهای واقعی خودتان آشنا شوید. در این مقاله از دیجیاتو قصد داریم به این موارد اشاره کنیم. پس تا آخر این مطلب با ما همراه باشید. ")

	//userStorageManager, err := application.NewMainStorageManager()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//chatHandler := handler.NewChatHandler(userStorageManager)
	//assetHandler := handler.NewAssetHandler(userStorageManager)
	//
	//initGin()
	//assetRouter(assetHandler)
	//chatRouter(chatHandler)
	//
	//startServer(router)
}

func assetRouter(assetHandler *handler.AssetHandler) {

	api := main2.router.Group("/api/assets/v5")

	// Asset routes
	api.POST("create", assetHandler.Create)
	api.POST("assets", assetHandler.Upload)
	api.GET(":id", assetHandler.Get)
	api.POST("update", assetHandler.Update)
	api.POST("update_all", assetHandler.UpdateAll)
	api.POST("delete", assetHandler.Delete)
	api.POST("filters", assetHandler.Filters)

	api.GET("download/:filename", assetHandler.OriginalDownload)
	api.GET("download/thumbnail/:filename", assetHandler.TinyImageDownload)
	api.GET("download/icons/:filename", assetHandler.IconDownload)

}

func chatRouter(chatHandler *handler.ChatHandler) *gin.Engine {

	api := main2.router.Group("/api/chat/v2")

	api.POST("create", chatHandler.Create)
	api.POST("update", chatHandler.Update)
	api.POST("delete", chatHandler.Delete)
	api.POST("list", chatHandler.GetList)

	return main2.router
}
