package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8888"
		log.Print("$PORT must be set")
	}

	router := gin.New() // gin 생성
	router.Use(gin.Logger()) // log 사용
	router.LoadHTMLGlob("templates/*.tmpl.html") // 읽어들일 html 지정
	router.Static("/static", "static") // 스타일 지정

	// 기본 localhost:8888 로 접속했을 때 동작하는 코드
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/chapter-1", func(c *gin.Context) {
		sum := Add(1, 3)
		expected := 4
		values := []obj{{Key: "sum?",Value: sum}, {Key: "expected?",Value: expected}, {Key: "test",Value: "123"}}
		c.HTML(http.StatusOK, "basic-learn.tmpl.html", gin.H{
			"chapter": "정수 1111",
			"obj": values,
		})
	})

	router.GET("/chapter-2", func(c *gin.Context) {
		values := []obj{{Key: "반복문 예제 실습하기 : ",Value: Repeat("가나다라")}}
		c.HTML(http.StatusOK, "basic-learn.tmpl.html", gin.H{
			"chapter": "반복",
			"obj": values,
		})
	})

	router.GET("/chapter-3", func(c *gin.Context) {
		values := []obj{
			{Key: "배열과 슬라이스 예제 실습하기 1번 : ",Value: Sum([]int{1, 2, 3, 4, 5})},
			{Key: "배열과 슬라이스 예제 실습하기 3번 : ",Value: SumAll([]int{1, 2}, []int{0, 9})},
			{Key: "배열과 슬라이스 예제 실습하기 4-5 - 1번 : ",Value: SumAllTails([]int{1, 2}, []int{0, 9})},
			{Key: "배열과 슬라이스 예제 실습하기 4-5 - 2번 : ",Value: SumAllTails([]int{}, []int{3, 4, 5})},
		}
		c.HTML(http.StatusOK, "basic-learn.tmpl.html", gin.H{
			"chapter": "배열과 슬라이스",
			"obj": values,
		})
	})

	router.Run(":" + port)
}