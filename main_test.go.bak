// main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func TestHello(t *testing.T) {
	// Echo のインスタンス作成
	e := echo.New()

	// テスト用のリクエストとレスポンスレコーダーを作成
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Hello ハンドラの実行
	handler := Hello()
	if err := handler(c); err != nil {
		t.Fatalf("handler error: %v", err)
	}

	// HTTP ステータスコードのチェック
	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	// レスポンス内容のチェック
	expected := "hello, world."
	if rec.Body.String() != expected {
		t.Errorf("expected body %q, got %q", expected, rec.Body.String())
	}
}
