<script setup lang="ts">
  import Breadcrumb from '@/components/Breadcrumbs.vue'
</script>

<template>
  <div class="container mx-auto">
    <Breadcrumb parentTitle='Stock Api' subParentTitle='股票資料接口' />

    <BaseCard class="p-4">
      <div class="apiDesc border-b border-gray-300 pb-8 mb-8">
        <p>
          本站資料從
          <a class="text-primary" href="https://mis.taifex.com.tw/futures/" target="_blank">臺灣期貨交易所</a>、
          <a class="text-primary" href="https://www.twse.com.tw/zh/index.html" target="_blank">臺灣證券交易所</a>
          爬蟲抓下，並儲存於 Google Cloud Platform，程式碼完全開源可見 
          <a class="text-primary" href="https://github.com/sdfsds5271463/twvixstock" target="_blank">GitHub</a>
          。
        </p>
        <p>資料開放 API 進行下載，以下接口請求方式允許 <b>GET</b> 或 <b>POST</b> ，請求格式允許 <b>get-query</b>、<b>form</b> 或 <b>json</b>。</p>
        <p>請求頻率同 IP <b>每2分鐘最多20次</b> ，請求過於頻繁將返回錯誤碼 3001。</p>
        <p>
          開始對接之前，試試看點一下這個連結
          <a class="text-primary" href="https://twvixstock.qzz.io/api/v1/stockDB?start=2026-01-01&end=2026-01-05" target="_blank">
            https://twvixstock.qzz.io/api/v1/stockDB?start=2026-01-01&end=2026-01-05
          </a>
          用 GET 輕鬆完成請求。
        </p>
      </div>

      <div class="">
          <p class="text-xl">域名 Domain</p>
          <p class="text-gray-600 border-b border-gray-300 pb-8 mb-8 pt-2 text-2xl font-bold">
            https://twvixstock.qzz.io/api/v1/
          </p>
      </div>


      <div class="">
          <p class="text-xl">表頭 Header</p>
          <p class="text-gray-600">與伺服器連線之資料請求與返回格式</p>
      </div>

      <div class="border-b border-gray-300 pb-8 mb-8 pt-2">
        <table class="apiTable">
          <tr>
            <td>Request Method</td><td>GET | POST</td>
          </tr>
          <tr>
            <td>Request content-type</td><td>null | application/json | application/x-www-form-urlencoded</td>
          </tr>
          <tr>
            <td>Response content-type</td><td>application/json</td>
          </tr>
        </table>
      </div>

      
      <div class="">
          <p class="text-xl">參數 Parameters</p>
          <p class="text-gray-600">與伺服器連線之資料請求與返回格式</p>
      </div>

      <div class="grid grid-cols-12 border-b border-gray-300 pb-8 mb-8">

        
        <div class="col-span-12 xl:col-span-6 m-2">
          <p class="mb-2">請求參數 Request parameters</p>
          <BaseCard noPadding class="overflow-x-auto">
            <div class="min-w-full bg-white rounded-xl shadow">
              <table class="min-w-full text-sm text-left">
                <thead class="bg-gray-100 text-gray-600 uppercase text-xs">
                  <tr>
                    <th class="pl-4 py-3">NAME</th>
                    <th class="pl-4 py-3">TYPE</th>
                    <th class="pl-4 py-3">SAMPLE</th>
                    <th class="pl-4 py-3">REQUIRED</th>
                    <th class="pl-4 py-3">REMARK</th>
                  </tr>
                </thead>

                <tbody class="divide-y">
                  <tr class="hover:bg-gray-50">
                    <td class="pl-4 py-4 font-medium">start</td>
                    <td class="pl-4 py-4">Date</td>
                    <td class="pl-4 py-4">2006-01-02</td>
                    <td class="pl-4 py-4 text-green-600">✅</td>
                    <td class="pl-4 py-4">起始日期(必填)</td>
                  </tr>

                  <tr class="hover:bg-gray-50">
                    <td class="pl-4 py-4 font-medium">end</td>
                    <td class="pl-4 py-4">Date</td>
                    <td class="pl-4 py-4">2006-01-02</td>
                    <td class="pl-4 py-4 text-red-600">❌</td>
                    <td class="pl-4 py-4">結束日期(預設當日)</td>
                  </tr>

                  <tr class="hover:bg-gray-50">
                    <td class="pl-4 py-4 font-medium">type</td>
                    <td class="pl-4 py-4">String</td>
                    <td class="pl-4 py-4">all</td>
                    <td class="pl-4 py-4 text-red-600">❌</td>
                    <td class="pl-4 py-4">股票類型(請參附錄)</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </BaseCard>
        </div>

        <div class="col-span-12 xl:col-span-6 m-2">
          <p class="mb-2">返回參數 Return parameters</p>
          <BaseCard noPadding class="overflow-x-auto">
            <div class="min-w-full bg-white rounded-xl shadow">
              <table class="min-w-full text-sm text-left">
                <thead class="bg-gray-100 text-gray-600 uppercase text-xs">
                  <tr>
                    <th class="pl-4 py-3">NAME</th>
                    <th class="pl-4 py-3">TYPE</th>
                    <th class="pl-4 py-3">SAMPLE</th>
                    <th class="pl-4 py-3">REQUIRED</th>
                    <th class="pl-4 py-3">REMARK</th>
                  </tr>
                </thead>

                <tbody class="divide-y">
                  <tr class="hover:bg-gray-50">
                    <td class="pl-4 py-4 font-medium">code</td>
                    <td class="pl-4 py-4">Number</td>
                    <td class="pl-4 py-4">1000</td>
                    <td class="pl-4 py-4 text-green-600">✅</td>
                    <td class="pl-4 py-4">狀態碼(請參附錄)</td>
                  </tr>

                  <tr class="hover:bg-gray-50">
                    <td class="pl-4 py-4 font-medium">error</td>
                    <td class="pl-4 py-4">String</td>
                    <td class="pl-4 py-4">too many requests</td>
                    <td class="pl-4 py-4 text-red-600">❌</td>
                    <td class="pl-4 py-4">僅失敗時返回error</td>
                  </tr>

                  <tr class="hover:bg-gray-50">
                    <td class="pl-4 py-4 font-medium">data</td>
                    <td class="pl-4 py-4">Array</td>
                    <td class="pl-4 py-4">[{"Type":"TAIPE","Date":"2026-02-01","Close":27.46}]</td>
                    <td class="pl-4 py-4 text-red-600">❌</td>
                    <td class="pl-4 py-4">僅成功時返回data</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </BaseCard>
        </div>
      </div>

      <div class="">
        <p class="text-xl">範例 Example</p>
        <p class="text-gray-600">請求成功與失敗的響應格式</p>
      </div>

      <div class="grid grid-cols-12 border-b border-gray-300 pb-8 mb-8">
        <div class="col-span-12 xl:col-span-6 m-2">
          ✅請求成功範例:
          <table class="apiTable">
            <tr>
              <td>Url</td><td>/api/v1/stockDB</td>
            </tr>
            <tr>
              <td>Header</td><td>content-type: application/json</td>
            </tr>
            <tr>
              <td>Parameters</td><td>{"start_date": "2026-01-01", "end_date": "2026-01-06", "stock_type": "taiex"}</td>
            </tr>
            <tr>
              <td>Return</td><td>{"code":1000,"data":[{"Type":"TAIEX","Date":"2026-01-02","Close":29349.81},{"Type":"TAIEX","Date":"2026-01-05","Close":30105.04},{"Type":"TAIEX","Date":"2026-01-06","Close":30576.3}]}</td>
            </tr>
          </table>

        </div>
        <div class="col-span-12 xl:col-span-6 m-2">
          ❌請求失敗範例:
          <table class="apiTable">
            <tr>
              <td>Url</td><td>/api/v1/stockDB</td>
            </tr>
            <tr>
              <td>Header</td><td>content-type: application/json</td>
            </tr>
            <tr>
              <td>Parameters</td><td>{"start_date": "", "end_date": "2026-01-06", "stock_type": "taiex"}</td>
            </tr>
            <tr>
              <td>Return</td><td>{"code":2001,"error":"Request is fail: Key: 'ReqStruct.Start' Error:Field validation for 'Start' failed on the 'required' tag"}</td>
            </tr>
          </table>
        </div>
      </div>



      <div class="">
          <p class="text-xl">附錄 Appendix</p>
          <p class="text-gray-600">參數及其他資料細節參考附錄</p>
      </div>

      <div class="grid grid-cols-12 gap-5 border-b border-gray-300 pb-10 mb-10 mt-2">
          <div class="col-span-12 sm:col-span-6 md:col-span-4">
            <p class="text-purple-500 mb-1"><i class="i-Calendar text-base"></i><span class="ml-1">股票類型 type</span></p>
            <table class="apiTable">
              <tr>
                <td>all</td><td>全部 (default)</td>
              </tr>
              <tr>
                <td>taiex</td><td>台灣加權指數</td>
              </tr>
              <tr>
                <td>vixtwn</td><td>台灣恐慌指數</td>
              </tr>
              <tr>
                <td>taipe</td><td>台灣加權本益比</td>
              </tr>
            </table>
          </div>
          <div class="col-span-12 sm:col-span-6 md:col-span-4">
            <p class="text-purple-500 mb-1"><i class="i-Gears text-base"></i><span class="ml-1">狀態碼 code</span></p>
            <table class="apiTable">
              <tr>
                <td>1000</td><td>請求成功 (http status:200)</td>
              </tr>
              <tr>
                <td>2001</td><td>請求參數錯誤 (http status:400)</td>
              </tr>
              <tr>
                <td>2002</td><td>資料傳遞錯誤 (http status:500)</td>
              </tr>
              <tr>
                <td>3001</td><td>請求過於頻繁 (http status:429)</td>
              </tr>
            </table>
          </div>
      </div>

    </BaseCard>



  </div>

</template>

<style lang="scss" scoped>

/*@media (max-width: 767px) {
}
@media (max-width: 414px) {
}*/

  .apiDesc{
    p{
      line-height: 18px;
      padding: 6px;
    }
  }

  .apiTable{
    td{
      padding: 2px 10px;
    }
  }

</style>
