TwVixStock - 全棧台股指數爬蟲與視覺化系統

這是一個基於 Go 語言開發的自動化台股爬蟲系統，整合了後端 API、排程任務、前端視覺化以及完整的雲端 CI/CD 自動化部署流程。

🌟 專案亮點

自動化排程：使用 Kubernetes CronJob 實現每日定時爬取台股盤後數據。

高效能後端：採用 Go Gin 框架，配合 Repository Pattern 進行乾淨的架構管理。

容器化管理：透過 Docker 多階段構建 (Multi-stage build) 優化 Image 大小。

DevOps 實踐：連結 GitHub 至 Google Cloud Build，實現 Push-to-Deploy 的全自動化部署。

HTTPS 保障：透過 Cloudflare 代理實現免費的 SSL 安全加密連線。

🏗️ 專案目錄架構
Plaintext
    twvixstock/
    ├── cmd/
    │   ├── server/          # Gin API 伺服器入口
    │   └── scheduler/       # 爬蟲排程啟動器
    ├── internal/
    │   ├── handler/         # HTTP 請求處理 (Controllers)
    │   ├── service/         # 核心業務邏輯 (Business Logic)
    │   ├── repository/      # 資料庫存取層 (SQL Operations)
    │   ├── model/           # 資料結構定義 (Domain Models)
    │   ├── crawler/         # 各類爬蟲具體邏輯
    │   └── config/          # 設定檔與環境變數讀取
    ├── frontend/            # Vue.js 前端專案
    ├── k8s/                 # Kubernetes 部署設定檔案 (Deployment, Service, CronJob)
    ├── scripts/             # SQL 初始化腳本與自動化腳本
    ├── Dockerfile           # 編譯產生映像檔
    ├── cloudbuild.yaml      # GCP Cloud Build 定義檔 (GitHub推送觸發)
    └── docker-compose.yaml  # 本地開發環境一鍵啟動

🛠️ 技術棧
Backend: Go (Gin Gonic)

Frontend: Vue.js 3 + Sass

Database: MySQL 8.0

Infrastructure: Google Kubernetes Engine (GKE), Artifact Registry

CI/CD: Google Cloud Build, GitHub Actions

Network: Cloudflare, Google Load Balancer

🚀 快速啟動 (本地開發)
複製專案

Bash
git clone https://github.com/your-username/twvixstock.git
cd twvixstock
環境變數設定
複製 .env.example 並修改為你的本地設定：

Bash
cp .env.example .env
使用 Docker Compose 啟動

Bash
docker-compose up -d
API Server: http://localhost:8080

Frontend: http://localhost:8081

☁️ 部署說明
本專案已完全自動化。當變更推送到 main 分支時：

Google Cloud Build 會觸發構建。

根據 Dockerfile 生成 Server 與 Scheduler 映像檔。

映像檔推送到 Artifact Registry。

gke-deploy 自動將最新版本部署至 GKE 叢集。

📬 聯絡資訊
作者: [allen]

專案連結: https://twvixstock.qzz.io/