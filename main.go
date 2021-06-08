/*
 * @Author: zzz
 * @Date: 2021-06-08 12:04:38
 * @LastEditTime: 2021-06-08 17:24:37
 * @LastEditors: zzz
 * @Description: 提供Windows Server API
 * @FilePath: \WinServerAPI\main.go
 */

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/kkzzhizhou/go-windows-server-api/hello"
)

type PowerShell struct {
	powerShell string
}

func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
}

func (p *PowerShell) Execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

// var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/flushdns", func(c *gin.Context) {
		pwsh := New()
		stdout, stderr, err := pwsh.Execute("[Console]::OutputEncoding = [Text.Encoding]::UTF8; Clear-DnsServerCache -Force")
		if err != nil {
			c.String(http.StatusOK, "flush failed.")
			// fmt.Println(err)
		} else {
			fmt.Println(stderr)
			fmt.Println(stdout)
			c.String(http.StatusOK, "flush finish.")
		}
	})

	return r
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	c := exec.Command("cmd", "/C", "Title", "Windows Server AP")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(hello.Greet())
	fmt.Println("作者: zzz")
	fmt.Println("说明: 在Windows Server上提供一些API接口，例如刷新DNS缓存等")
	fmt.Println("使用: http://ip:5000/flushdns // 刷新DNS缓存")
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":5000")
}
