package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pterodactyl/wings/router/middleware"
	"io"
	"mime"
	"net/http"
	"os"
	"errors"
	"strings"
	"github.com/pterodactyl/wings/config"
)

// Download Plugin
func downloadPlugin(c *gin.Context) {
	// Get the requested server
	s := middleware.ExtractServer(c)

	// Get the download url
	var data struct {
		DOWNLOAD_URL string `json:"download_url"`
	}

	// Validate the download url
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Invalid download url."})
		return
	}

	if data.DOWNLOAD_URL == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Invalid download url."})
		return
	}

	// Create plugins folder
	if err := s.Filesystem().CreateDirectory("plugins", "/"); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create the plugins directory."})
		return
	}

	// Download plugin
	filename, err := DownloadFile(fmt.Sprintf("%s/%s/plugins/", config.Get().System.Data, c.Param("server")), data.DOWNLOAD_URL)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "filename": filename})
}

// Download file from URL
func DownloadFile(filepath string, url string) (string, error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	// Get the file name
	_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	if err != nil {
		return "", err
	}

	// Get filename
	split := strings.Split(params["filename"], "#")
	filename := split[0] + ".jar"

	// Modify download path
	filepath += filename

	if filename == "" || filename == ".jar" {
		return "", errors.New("invalid name")
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return "", err
	}

	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return filename, err
}
