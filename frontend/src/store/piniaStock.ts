import { defineStore } from 'pinia'

export const piniaStock = defineStore('piniaStock', {
  /* pinia 使用方法:
      1.  npm install pinia
      2.  src/main.ts 加入
            import { createPinia } from 'pinia'
            const pinia = createPinia()
            app.use(pinia)
      3.  創建這支範例檔案( src/store/piniaStock.ts )
      4.  在要引入的地方
            import { onMounted,ref } from "vue"
            import { piniaStock } from '@/store/piniaStock'
            const piniaStockMain = piniaStock()
            onMounted(async (): Promise<void> => {
              piniaStockMain.submitStockDB({ start: "2026-02-01", end: "2026-02-28", type: "TAIEX" });
            })
  */
  state: () => ({
    ajaxData: null,
    loading: false,
    error: null,
  }),

  actions: {

    async submitStockDB(payload) {
      this.loading = true
      this.error = null

      try {
        const response = await fetch('/api/v1/stockDB', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(payload)
        });

        if (!response.ok) throw new Error('請求失敗');
        const data = await response.json();
        this.ajaxData = data; // 將結果存入 state

        console.log("pinia AJAX:", data);

      } catch (err) {
        this.error = err.message
      } finally {
        this.loading = false
      }
    },
  },

})