package main

import (
	"bytes"
	"fmt"
	"io"
	"myapp/route"
	"net/http"
	"time"
)

// ログ出力用のミドルウェア
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jpn, _ := time.LoadLocation("Asia/Tokyo")
		start := time.Now().In(jpn)

		bodyBytes, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // ハンドラーでも読めるよう復元

		// 次の処理（実際のハンドラー）を実行
		next.ServeHTTP(w, r)

		// 処理が終わった後にログを出力
		fmt.Printf("[%s] %s %s %s %s\n",
			start.Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			string(bodyBytes), // クエリパラメータ
			time.Since(start), // 処理にかかった時間
		)
	})
}

func main() {
	fmt.Println("Server Start !")
	http.ListenAndServe(":8080", loggingMiddleware(route.SetRoute()))
}
