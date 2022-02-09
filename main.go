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

	// 페이지 내 chapter-1로 라우팅 되는 코드
	router.GET("/chapter-1", func(c *gin.Context) {
		// basic-func.go의 Add 함수
		sum := Add(1, 3)
		expected := 4
		// interface.go의 obj 인터페이스
		// Key에 들어갈 값과 Value에 들어갈 값을 설정
		values := []obj{{Key: "sum?",Value: sum}, {Key: "expected?",Value: expected}, {Key: "test",Value: "123"}}
		// 호출할 html 설정과 변수 입력
		c.HTML(http.StatusOK, "basic-learn.tmpl.html", gin.H{
			// html 코드 내에서 사용할 변수 할당
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

	/* TODO /chapter-4에 들어갈 항목을 정리한다.
		1. 라우터 생성
		2. 연결할 HTML 설정
		3. html 내에서 사용할 변수 설정
	*/

	router.Run(":" + port)
}