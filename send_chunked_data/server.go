package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		w := c.Writer
		header := w.Header()
		header.Set("Tranfer-Encoding", "chunked")
		header.Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<html><body><p>test chunked data</p>"))
		w.(http.Flusher).Flush()
		for i := 0; i < 10; i++ {
			w.Write([]byte(fmt.Sprintf("<h1>%d</h1>", i)))
			w.(http.Flusher).Flush()
			time.Sleep(time.Duration(1) * time.Second)
		}
		w.Write([]byte("<p>Done!</p></body></html>"))
		w.(http.Flusher).Flush()
	})

	r.Run("127.0.0.1:8080")
}
