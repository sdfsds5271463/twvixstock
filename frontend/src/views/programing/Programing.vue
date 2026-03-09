<script setup lang="ts">
  import Breadcrumb from '@/components/Breadcrumbs.vue'
</script>

<template>
  <div class="container mx-auto">
    <Breadcrumb parentTitle='Programing' subParentTitle='本站主程式' />
    <BaseCard class="overflow-hidden relative mb-8 p-4">
      <div class="programing_desc card-body">
        <p>
          本站開源於
          <a class="text-primary" href="https://github.com/sdfsds5271463/twvixstock" target="_blank">https://github.com/sdfsds5271463/twvixstock</a><br>
        </p>
        <p>可執行 <b>git clone https://github.com/sdfsds5271463/twvixstock.git</b> 將專案建置到你的機器上</p>
        <p>本站詳細架構及技術請參照下方 README</p>
      </div>
    </BaseCard>

<BaseCard class="Programing_pre overflow-x-auto p-2 text-gray-300">
<pre><b class="text-sm text-primary">README.md</b>
<b>TwVixStock - 台股恐慌資料站 (台股指數爬蟲與視覺化分析系統)</b>

<hr>
🌈 網站摘要

這是一個開源的台股恐慌資料站，其核心分析資料為 台灣加權指數、台灣恐慌指數(台指選擇權波動指數)、台灣加權本益比。
本站後端基於 Go + Gin 開發，排程爬蟲將會蒐集 臺灣期貨交易所、臺灣證券交易所 之每日指數資料，並提供 API 開源接口可下載。
本站前端使用 Vue3 Vite + TypeScript + Tailwind CSS 撰寫，可進行台股恐慌程度的視覺化分析。
專案自動部屬在 Google Cloud Platform(GCP)，開源代碼可從 GitHub 下載。

<hr>
🌟 專案亮點

現代化前端：使用主流前端框架 Vue3 Vite 搭配 Tailwind CSS 進行高度自由的前端樣式開發。
自動化排程：使用 Kubernetes CronJob 實現每日定時爬取台股盤後數據。
高效能後端：採用 Go Gin 框架，配合 Repository Pattern 進行乾淨的架構管理。
容器化管理：透過 Docker 多階段構建 (Multi-stage build) 優化 Image 大小。
DevOps 實踐：連結 GitHub 至 Google Cloud Build，實現 Push-to-Deploy 的全自動化部署。
HTTPS 保障：透過 Cloudflare 代理實現免費的 SSL 安全加密連線。

<hr>
🏗️ 專案目錄架構

Plaintext
    twvixstock/
    ├── cmd/
    │   ├── server/          # Gin API 伺服器入口
    │   └── scheduler/       # 爬蟲排程啟動器
    ├── internal/
    │   ├── handler/         # HTTP 請求處理
    │   ├── service/         # 核心業務邏輯
    │   ├── repository/      # 資料庫存取層
    │   ├── model/           # 資料結構定義
    │   ├── crawler/         # 各類爬蟲具體邏輯
    │   └── config/          # 設定檔與環境變數讀取
    ├── frontend/            # Vue.js 前端專案
    ├── k8s/                 # Kubernetes 部署設定檔案 (Deployment, Service, CronJob)
    ├── scripts/             # SQL 初始化腳本與自動化腳本
    ├── Dockerfile           # 編譯產生映像檔
    ├── cloudbuild.yaml      # GCP Cloud Build 定義檔 (GitHub推送觸發)
    └── docker-compose.yaml  # 本地開發環境一鍵啟動

<hr>
🛠️ 技術棧

Backend: Go (Gin Gonic)
Frontend: Vue3 Vite + TypeScript + Tailwind CSS + Sass
Database: MySQL 8.0
Infrastructure: Google Kubernetes Engine (GKE), Artifact Registry
CI/CD: Google Cloud Build, GitHub Actions
Network: Cloudflare, Google Load Balancer

<hr>
🚀 快速啟動 (本地開發)

[複製專案]
git clone https://github.com/sdfsds5271463/twvixstock.git
cd twvixstock

[安裝依賴]
go mod tidy
cd frontend
npm install | pnpm install | yarn install

[環境變數設定]
cp .env.example .env

[使用 Docker Compose 啟動]
docker build -t tvs-server --target server .
docker build -t tvs-scheduler --target scheduler .
docker-compose up -d

Server: http://localhost:8080
Api:    http://localhost:8080/api/v1/stockDB?start=2026-01-01&end=2026-01-06&type=taiex

<hr>
☁️ 部署說明

本專案已完全自動化。當變更推送到 main 分支時：
Google Cloud Build 會觸發構建。
根據 Dockerfile 生成 Server 與 Scheduler 映像檔。
映像檔推送到 Artifact Registry。
gke-deploy 自動將最新版本部署至 GKE 叢集。

<hr>
📬 聯絡資訊

作者: [allen.Zheng]
專案連結: https://twvixstock.qzz.io/
專案GitHub: https://github.com/sdfsds5271463/twvixstock

</pre>
</BaseCard>

  </div>
</template>


<style lang="scss" scoped>

/*@media (max-width: 767px) {
}
@media (max-width: 414px) {
}*/
  .programing_desc{
    p{
      line-height: 18px;
      padding: 5px;
    }
  }
  .Programing_pre{
    background: linear-gradient(to bottom, #132c67, #15202e);
  }

</style>
