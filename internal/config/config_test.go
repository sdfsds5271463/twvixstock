package config

import "testing"
import "os"

func TestLoadConfig(t *testing.T) {
	os.Chdir("../..")
	LoadConfig("")
	if AppConfig.Database.Host == "" {
		t.Error("未能讀取到資料庫 Host 設定")
	} else {
		t.Logf("讀取成功，Host 為: %s", AppConfig.Database.Host)
	}
}