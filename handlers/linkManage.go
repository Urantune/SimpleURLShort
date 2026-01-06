package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"SimpleURLShortener/linkStorage"
	"SimpleURLShortener/utils"
)

func CreateShortLink(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url khong rong"})
		return
	}

	req.URL = CheckLinkHTTP(req.URL)

	code := ""
	//for true {
	for i := 0; i < 20; i++ {
		tmp := utils.GenerateCode(6)
		if !linkStorage.CodeExists(tmp) {
			code = tmp
			break
		}
	}

	if err := linkStorage.SaveLink(code, req.URL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "save that bai"})
		return
	}

	base := "http://" + c.Request.Host
	c.JSON(http.StatusOK, gin.H{
		"code":         code,
		"short_url":    base + "/" + code,
		"original_url": req.URL,
	})
}

func CreateCodeByLink(c *gin.Context) {
	link := c.Query("url")
	if link == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "thieu query param url"})
		return
	}

	link = CheckLinkHTTP(link)

	code := ""
	for i := 0; i < 20; i++ {
		tmp := utils.GenerateCode(6)
		if !linkStorage.CodeExists(tmp) {
			code = tmp
			break
		}
	}
	if code == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "khong tao duoc code"})
		return
	}

	if err := linkStorage.SaveLink(code, link); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "save that bai"})
		return
	}

	base := "http://" + c.Request.Host
	c.JSON(http.StatusOK, gin.H{
		"code":         code,
		"short_url":    base + "/" + code,
		"original_url": link,
	})
}

func CheckLinkHTTP(rootString string) string {
	if !strings.HasPrefix(rootString, "http://") && !strings.HasPrefix(rootString, "https://") {
		return "https://" + rootString
	}

	return rootString
}

func ConnectLink(c *gin.Context) {
	code := c.Param("code")

	link, err := linkStorage.GetCode(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	_ = linkStorage.IncreaseVisit(code)
	c.Redirect(http.StatusFound, link.OriginalURL)
}

func ListLinks(c *gin.Context) {
	links, err := linkStorage.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error roi"})
		return
	}
	c.JSON(http.StatusOK, links)
}

func Stats(c *gin.Context) {
	code := c.Param("code")

	link, err := linkStorage.GetCode(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, link)
}
