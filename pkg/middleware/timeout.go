package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jacobintern/meme_coin/controller"
	"github.com/jacobintern/meme_coin/pkg/errors"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 建立帶有 Timeout 的 Context
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel() // 確保請求結束時釋放資源

		// 用帶 Timeout 的 Context 替換原本的 Request
		c.Request = c.Request.WithContext(ctx)

		// 使用 Go Routine 監聽是否超時
		done := make(chan struct{})
		go func() {
			c.Next() // 執行後續的 Middleware / Handler
			close(done)
		}()

		select {
		case <-ctx.Done(): // 當 Context 超時時
			controller.Failed(c, errors.GatewayTimeout)
			c.Abort() // 取消請求
			return
		case <-done: // 正常請求完成
		}
	}
}
