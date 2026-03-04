package repository

import (
	"testing"
	"twvixstock/internal/model"
	"twvixstock/internal/config"
	"time"
	"os"

	"github.com/stretchr/testify/assert"
)

func TestInitDBAndCRUD(t *testing.T) {
	os.Chdir("../..")

	// 2. 初始化實體 DB 連線
	config.LoadConfig("")  //載入 .env
	InitDB()  //初始 DB

	// 檢查全域變數 DB 是否成功賦值
	assert.NotNil(t, DB, "DB 連線不應為 nil")

	// 3. 測試：寫入一筆測試資料
	testRecord := model.StockRecord{
		Type:  "TAIEX",
		Date:  time.Now(),
		Close: 18000.50,
	}

	// Create
	result := DB.Create(&testRecord)
	assert.NoError(t, result.Error, "寫入資料時不應有錯誤")
	assert.NotZero(t, testRecord.ID, "寫入後 ID 應自動生成")

	// 4. 測試：讀取剛寫入的資料
	var fetchedRecord model.StockRecord
	DB.First(&fetchedRecord, testRecord.ID)
	assert.Equal(t, "TAIEX", fetchedRecord.Type)
	assert.Equal(t, 18000.50, fetchedRecord.Close)

	// 5. 測試：刪除該測試資料 (保持環境乾淨)
	deleteResult := DB.Delete(&testRecord)
	assert.NoError(t, deleteResult.Error)
}