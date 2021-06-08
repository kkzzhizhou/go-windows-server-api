/*
 * @Author: zzz
 * @Date: 2021-06-08 17:06:22
 * @LastEditTime: 2021-06-08 17:06:42
 * @LastEditors: zzz
 * @Description:
 * @FilePath: \WinServerAPI\hello\hello_test.go
 */

package hello

import "testing"

func TestGreet(t *testing.T) {
	result := Greet()
	if result != "欢迎使用Windows Server API" {
		t.Errorf("Greet() = %s; Expected 欢迎使用Windows Server API", result)
	}

}
