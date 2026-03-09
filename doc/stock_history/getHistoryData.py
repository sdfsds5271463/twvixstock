# 透過json資料，初始化DB的歷史紀錄 (20200101~20260304)
# python getHistoryData.py

#執行提示
confirm = input(
'''本程序將會DB刪減表 stock
並且依據json檔案插入 2017-01-01 ~ 2026-03-04 初始指數資料
確定執行嗎(y|Y)? ''')

if(confirm.upper() != "Y"):
    print("取消動作")
    exit()
print("準備 DB stock 初始指數資料")

# pip install python-dotenv mysql-connector-python
import os
import mysql.connector
from dotenv import load_dotenv
import json
import time

# 載入 .env 檔案
load_dotenv()
db_host = os.getenv('DB_HOST')
db_port = 3306
db_user = os.getenv('DB_USER')
db_pass = os.getenv('DB_PASSWORD')
db_name = os.getenv('DB_NAME')

def connect_to_db():
    try:
        # 建立連線
        connection = mysql.connector.connect(
            host=db_host,
            port=db_port,
            user=db_user,
            password=db_pass,
            database=db_name
        )
        if connection.is_connected():
            print(f"成功連線到資料庫: {db_name}")
            return connection
    except mysql.connector.Error as err:
        print(f"連線失敗: {err}")
        return None

if __name__ == "__main__":
    conn = connect_to_db()
    if conn:
        # 資料庫指標
        cursor = conn.cursor()

        # 刪減表(重新初始)
        query = "TRUNCATE TABLE stock;"
        cursor.execute(query)

        #總資料數
        datacount = 0

        # 台灣加權指數 TAIEX
        print(" - 初始化 台灣加權指數 TAIEX")
        with open("taiex.json", "r", encoding="utf-8") as f:
            data = json.load(f)
            for v in data['data']:
                # 最遠只存到 2017-01-01
                if v['date'] < "2017-01-01":
                    continue
                # 執行插入指令
                sql_query = """
                    INSERT INTO stock (type, date, close) VALUES (%s, %s, %s)
                    ON DUPLICATE KEY UPDATE close = VALUES(close)
                """
                data = ["TAIEX", v['date'], v['close']]

                # 執行
                datacount += 1
                cursor.execute(sql_query, data)
            
        # 台灣恐慌指數 VIXTWN
        print(" - 初始化 台灣恐慌指數 VIXTWN")
        with open("vixtwn.json", "r", encoding="utf-8") as f:
            data = json.load(f)
            for v in data['data']['c:46328']['series'][0]:
                # 最遠只存到 2017-01-01
                if v[0] < "2017-01-01":
                    continue
                # 執行插入指令
                sql_query = """
                    INSERT INTO stock (type, date, close) VALUES (%s, %s, %s)
                    ON DUPLICATE KEY UPDATE close = VALUES(close)
                """
                data = ["VIXTWN", v[0], v[1]]

                # 執行
                datacount += 1
                cursor.execute(sql_query, data)

        # 台灣本益比 TAIPE
        print(" - 初始化 台灣本益比 TAIPE")
        with open("taipe.json", "r", encoding="utf-8") as f:
            data = json.load(f)
            for v in data['Data'][0]['List']:

                t = time.localtime(v['UTC_dataDate'] / 1000)  #時間參數(用時間戳轉)
                t_str = time.strftime("%Y-%m-%d", t)

                # 最遠只存到 2017-01-01
                if t_str < "2017-01-01":
                    continue
                # 執行插入指令
                sql_query = """
                    INSERT INTO stock (type, date, close) VALUES (%s, %s, %s)
                    ON DUPLICATE KEY UPDATE close = VALUES(close)
                """
                data = ["TAIPE", t_str, v['value']]

                # 執行
                datacount += 1
                cursor.execute(sql_query, data)


        # 提交
        conn.commit()

        # 讀取資料數
        query = "SELECT * FROM stock"
        cursor.execute(query)
        rows = cursor.fetchall()
        print(f" - 成功插入 {len(rows)} 筆資料\n")

        # 檢查
        if (len(rows) == datacount):
            print("資料數量沒有丟失，程序結束")
        else:
            print(f"資料數量異常 {len(rows)} != {datacount} ，請檢查或再次執行")

        # 關閉連線
        cursor.close()
        conn.close()


'''
JSON資料來源:::

TAIEX
 FROM https://api.finmindtrade.com/api/v4/data?dataset=TaiwanStockPrice&data_id=TAIEX&start_date=2016-01-01
   2016-01-04 ~ 2026-03-05

VIXTWN
 FROM https://www.macromicro.me/charts/46328/tw-vixtwn
  ==> https://www.macromicro.me/charts/data/46328
   2016-11-25 ~ 2026-03-04

TAIPE
 FROM https://www.yuanta-etfadvisor.com/map/chart?area=80008d7a-64b7-450c-9b96-49d7bf712b34&staticCategoryName=MacroEconomicIndex&product1=507f7d79-b6e2-44b1-902d-f304253322fc
  ==> https://api.yuantafunds.com/ECTranslationAI/api/bridge?APIType=ConsultingAPI&FuncId=Index%2FMarcoEconomicIndex&code=Y00114&CompanyName=YUANTAFUNDS&Platform=ETFXAI&AppName=ETFxAI
   2007-05-07 ~ 2026-02-01
'''