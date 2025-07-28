package redirect

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Redirect() {
	r := gin.Default()

	r.GET("/porto/callback", RedirectPorto)

	_ = r.Run(":8080")
}

func RedirectPorto(c *gin.Context) {
	c.Redirect(http.StatusFound, "https://riyanathariq.github.io/portofolio/")
}
