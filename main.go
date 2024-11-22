package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

func main() {
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
