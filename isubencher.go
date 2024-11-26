package isubencher

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

type BenchMarkerFunc func(targetIP string, sslEnabled bool, executer func() error) (string, error)

// 環境変数からISUCONの大会名を取得する
func GetIsuconName() (string, error) {
	iusuconName := os.Getenv("ISUCON_NAME")

	if iusuconName == "" {
		return "", fmt.Errorf("ISUCON_NAME is not set")
	}

	return iusuconName, nil
}

func getBenchMarker(isuconName string) BenchMarkerFunc {
	switch isuconName {
	case "kayac-isucon-listen80":
		return KayacIsuconListen80
	default:
		return NoCompetition
	}
}

func KayacIsuconListen80(targetIP string, sslEnabled bool, executer func() error) (string, error) {
	return "", fmt.Errorf("No competition")
}

func NoCompetition(targetIP string, sslEnabled bool, executer func() error) (string, error) {
	return "", fmt.Errorf("No competition")
}

func BenchHandler(w http.ResponseWriter, r *http.Request) {
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

	cmd := exec.Command("./bench", "-target-url", benchProtcol+"://"+targetIP)

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
