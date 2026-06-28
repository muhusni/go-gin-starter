package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const faviconSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
  <rect width="64" height="64" rx="14" fill="#111827"/>
  <path d="M18 20h28v8H34v8h10v8H34v12H24V20z" fill="#f9fafb"/>
</svg>`

func FaviconHandler(c *gin.Context) {
	c.Header("Cache-Control", "public, max-age=86400")
	c.Data(http.StatusOK, "image/svg+xml", []byte(faviconSVG))
}
