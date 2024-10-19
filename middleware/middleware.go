package middleware

import(
	"github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "log"
    "time"
    "net/http"
    // "github.com/golang-jwt/jwt/v5"
    "auto_slip_api/pkg/utils"
)

// リクエストのログを記録
func RecordUaAndTime(c *gin.Context){
   logger, err := zap.NewProduction()
   if err != nil{
      log.Fatal(err.Error())
   }
   oldTime := time.Now()
   ua := c.GetHeader("User-Agent")
   c.Next()
    logger.Info("incoming request",
        zap.String("path", c.Request.URL.Path),
        zap.String("Ua", ua),
        zap.Int("status", c.Writer.Status()),
        zap.Duration("elapsed", time.Now().Sub(oldTime)),
    )
	
}

// 認証用ミドルウェア
func AuthMiddleware(ctx *gin.Context) {
	headerAuthorization := ctx.GetHeader("Authorization")
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "Unauthorized",
	// 	})
	// 	ctx.Abort()
	// 	return
	// }

	ok, err := utils.ParseToken(headerAuthorization)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		ctx.Abort()
		return
	}
    if !ok {
        ctx.JSON(http.StatusUnauthorized, gin.H{
            "message": "Unauthorized",
        })
        ctx.Abort()
        return
    }
	ctx.Next()  // エンドポイントの処理を続行
}
