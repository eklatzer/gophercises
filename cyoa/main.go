package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"gophercises/cyoa/config"
	"gophercises/cyoa/story"
)

var cfg = config.Config{}

func init() {
	flag.StringVar(&cfg.StoryFile, "story", "gopher.json", "Filepath of the file with the story")
	flag.StringVar(&cfg.TemplatePath, "template", "templates", "Path to the HTML templates")
	flag.IntVar(&cfg.Port, "port", 8080, "Port used for the web application")
	flag.Parse()
}

func main() {
	var story = story.Story{}
	err := story.ReadStoryFromFile(cfg.StoryFile)
	if err != nil {
		log.Fatalf("faild to read story from file: %v", err)
	}

	router := gin.Default()
	log.Println(filepath.Join(cfg.TemplatePath, "*"))
	router.LoadHTMLGlob(filepath.Join(cfg.TemplatePath, "*"))

	router.GET("/story/:name", func(c *gin.Context) {
		nameOfStoryPart := c.Param("name")
		storyPart, exists := story[nameOfStoryPart]

		if !exists {
			c.Redirect(http.StatusFound, "/404")
		}

		c.HTML(http.StatusOK, "template.html", gin.H{
			"Story": storyPart,
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	})

	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
