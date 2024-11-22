package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

const (
	// ベンチマーク実行用のバイナリパス
	binaryPath = "./kayac-isucon-2022/bench"
)

func main() {
	// 実行権限を付与
	// err := os.Chmod(binaryPath+"/bench", 0755)
	// if err != nil {
	// 	fmt.Println("権限変更エラー:", err)
	// 	return
	// }

	s := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/bench", benchHandler)
	s.ListenAndServe()
}

func benchHandler(w http.ResponseWriter, r *http.Request) {
	targetIP := r.URL.Query().Get("target-ip")
	if targetIP == "" {
		http.Error(w, "target-ip is required", http.StatusBadRequest)
		return
	}

	sslEnabled := r.URL.Query().Get("ssl-enabled")
	sslFlag, err := strconv.ParseBool(sslEnabled)
	if err != nil {
		http.Error(w, "ssl-enabled is invalid", http.StatusBadRequest)
		return
	}

	benchProtcol := "http"
	if sslFlag {
		benchProtcol = "https"
	}

	// curlコマンド実行
	cmd := exec.Command("./bench", "-target-url", benchProtcol+"://"+targetIP)
	cmd.Dir = binaryPath

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		if os.IsPermission(err) {
			fmt.Println("実行権限が不足しています。権限を確認してください。")
		}
		http.Error(w, "bench failed", http.StatusInternalServerError)
		return
	}

	w.Write(out)
}
