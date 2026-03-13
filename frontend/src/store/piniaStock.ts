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
    //AJAX用資料
      //股票資料
      stockData: null,
      stockLoading: false,
      stockError: null,
      stockRetry: 4,
      //gemini資料
      geminiData: null,
      geminiLoading: false,
      geminiError: null,
      geminiRetry: 4,

    //分析用資料
      //stock
      DateArr: [] as string[],
      TAIEX: [] as number[],
      VIXTWN: []as number[],
      TAIPE: [] as number[],
      InitialStockReady: false,
      //gemini
      GeminiPoint: 0,
      InitialGeminiReady: false,
  }),

  actions: {
    //初始化 (請求5年資料 + gemini資料存下來)
    initPiniaStock(){
      //時間決定
      const getDays = 1827;  //抓取幾天資料
      const timeStart = new Date(new Date().getTime() - 86400000 * getDays);  //時間物件
      const timeStr = timeStart.toISOString().split('T')[0];  //2026-02-10

      //抓取
      this.submitStockDB({ start: timeStr, end: "", type: "ALL" })
      this.submitGeminiData({ date:"" })  //留空抓最新
    },
    
    //請求股票資料
    async submitStockDB(payload) {
      this.stockLoading = true;
      this.stockError = null;

      //發送
      try {
//console.log("POST A", this.stockRetry)
        const response = await fetch('/api/v1/stockDB', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(payload)
        });

        if (!response.ok) throw new Error('stockDB 請求失敗');
        const data = await response.json();

        if (data.code != 1000) throw new Error('stockDB 資料狀態異常 ' + data.code);
        this.stockData = data; // 將結果存入 state
        console.log("pinia stockDB AJAX:", this.stockRetry, data);

        this.analyzeStockDB(); //分析
        this.stockLoading = false;
      } catch (err) {
        //出錯
        this.stockError = err.message;
        console.log("pinia stockDB AJAX error", err);

        //重試
        this.stockRetry -= 1;
        if(this.stockRetry > 0){
          setTimeout(() => {
            this.submitStockDB(payload);
          }, 5000);
        }else{
          //沒機會重試，結束
          this.stockLoading = false;
        }
      } finally {}
    },

    //請求Gemini資料
    async submitGeminiData(payload) {
      this.geminiLoading = true;
      this.geminiError = null;

      //發送
      try {
//console.log("POST B", this.geminiRetry)
        const response = await fetch('/api/v1/geminiData', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(payload)
        });

        if (!response.ok) throw new Error('geminiData 請求失敗');
        const data = await response.json();

        if (data.code != 1000) throw new Error('geminiData 資料狀態異常 ' + data.code);
        this.geminiData = data; // 將結果存入 state
        console.log("pinia geminiData AJAX:", this.geminiRetry, data);

        //準備完成
        this.InitialGeminiReady = true;
        this.geminiLoading = false;
        this.GeminiPoint = data.data.Point;
      } catch (err) {
        //出錯
        this.geminiError = err.message;
        console.log("pinia geminiData AJAX error", err);

        //重試
        this.geminiRetry -= 1;
        if(this.geminiRetry > 0){
          setTimeout(() => {
            this.submitGeminiData(payload);
          }, 5000);
        }else{
          //沒機會重試，結束
          this.geminiLoading = false;
        }
      } finally {}
    },

    // 分析取回來的 StockDB 資料
    analyzeStockDB(){
      // 沒資料不過
      if(this.stockData == null){
        return
      }
      
      // 開始分析
      //  1. 所有日期標準依據 TAIEX， 其他有空缺的要計算補上
      //  2. 補法都是等比例縮放，不過 VIX 是倒數


      //第一步 完善日期陣列 DateArr
      let DateArr:string[] = [] //日期
      let Kindex:{ [key: string]: number } = {}  //鍵索引(未來直接全參數對其日期)
      let MonthLast:{ [key:string]: string } = {}  //每月最後一天的日期(用來調整本益比的實際結算日)

      this.stockData.data.forEach(v => {
        //記錄下所有日期
        if(v.Type == "TAIEX"){
          DateArr.push(v.Date); //所有日期
          MonthLast[v.Date.substring(0,7)] = v.Date; //每月最後日
        }

        //紀錄下所有鍵索引
        if(v.Type == "TAIEX" || v.Type == "VIXTWN"){
          Kindex[v.Date + v.Type] = v.Close  //月日都記下
        }else if(v.Type == "TAIPE"){
          Kindex[v.Date.substring(0,7) + v.Type] = v.Close  //PE只記錄月 (日無意義)
        }
      });
      //console.log(DateArr, Kindex);


      //第二步 完善其餘陣列
      // 先將 TAIPE 給予正確的 每月最後一天
      for(const[k,v] of Object.entries(MonthLast)){
        if(Kindex[k + "TAIPE"] != undefined){  //有該月的日期
          Kindex[v + "TAIPE"] = Kindex[k + "TAIPE"];  //給予最後一天
        }
      }

      // 用鍵 填入所有值 (該日期沒值，會留空)
      let TAIEXArr:number[] = []
      let VIXTWNArr:number[] = []
      let TAIPEArr:number[] = []
      DateArr.forEach((v,k) => {
        //完善 TAIEXArr VIXTWNArr
        TAIEXArr[k] = Kindex[v + "TAIEX"]?Kindex[v + "TAIEX"]:0;
        VIXTWNArr[k] = Kindex[v + "VIXTWN"]?Kindex[v + "VIXTWN"]:0;
        TAIPEArr[k] = Kindex[v + "TAIPE"]?Kindex[v + "TAIPE"]:0;
      })  
//console.log( DateArr, TAIEXArr, VIXTWNArr, TAIPEArr);  //此四個陣列長度一樣，空缺的部分會變成0


      //第三步 估值填入PE有空缺的值
      // 順序權重
      let wAsc:[number,number][] = [];  //值 權重
      let nowV:number = 0; //當前值
      let nowEX:number = 0; //當前EX
      let nowW:number = 0; //當權重
      for(let k in TAIPEArr){
        if(TAIPEArr[k] != 0){  //找到值
          nowV = TAIPEArr[k]; //當前值
          nowW = 1; //權重為1
          nowEX = TAIEXArr[k]; //當前EX 
        }else if(nowV != 0){ //有值時
          nowW += 1; //權重增加
        }
        wAsc[k] = [nowEX?(nowV*TAIEXArr[k]/nowEX):0, nowW]; //紀錄
      }

      // 倒序權重
      let wDesc:[number,number][] = [];  //值 權重
      nowV = 0; nowEX = 0; nowW = 0;
      for(let k = (TAIPEArr.length-1); k>=0; k--){
        if(TAIPEArr[k] != 0){  //找到值
          nowV = TAIPEArr[k]; //當前值
          nowW = 1; //權重為1
          nowEX = TAIEXArr[k]; //當前EX 
        }else if(nowV != 0){ //有值時
          nowW += 1; //權重增加
        }
        wDesc[k] = [nowEX?(nowV*TAIEXArr[k]/nowEX):0, nowW]; //紀錄
      }
//console.log(wAsc, wDesc)
      //填入值
      TAIPEArr.forEach((v,k)=>{
        //計算權重
        if(wAsc[k][0] != 0 && wDesc[k][0] != 0){
          let sumW = wAsc[k][1] + wDesc[k][1]; //總權重
          let newV = (wAsc[k][0] * wDesc[k][1]) + (wDesc[k][0] * wAsc[k][1]); //互相乘對方權重為新值
          newV = newV / sumW;  //除以權重後 .2f
          TAIPEArr[k] = newV;
        }else if(wAsc[k][0] != 0){ //兩端不用算權重
          TAIPEArr[k] = wAsc[k][0];
        }else if(wDesc[k][0] != 0){ //兩端不用算權重
          TAIPEArr[k] = wDesc[k][0];
        }
      })
//console.log(TAIPEArr)

      //第四步 填入 vix 值
      // 順序填值(數字很密，不用算權重)
      nowV = 0; nowEX = 0;
      for(let k in VIXTWNArr){
        if(VIXTWNArr[k] != 0){  //找到值
          nowV = VIXTWNArr[k]; //當前值
          nowEX = TAIEXArr[k]; //當前EX 
        }
        if(nowV != 0 && VIXTWNArr[k] == 0){  //當前可填值
          VIXTWNArr[k] = nowV * nowEX / TAIEXArr[k];  //直接給倒EX比例值
        }
      }
      // 倒序填值(數字很密，不用算權重)
      nowV = 0; nowEX = 0;
      for(let k = (VIXTWNArr.length-1); k>=0; k--){
        if(VIXTWNArr[k] != 0){  //找到值
          nowV = VIXTWNArr[k]; //當前值
          nowEX = TAIEXArr[k]; //當前EX 
        }
        if(nowV != 0 && VIXTWNArr[k] == 0){  //當前可填值
          VIXTWNArr[k] = nowV * nowEX / TAIEXArr[k];  //直接給倒EX比例值
        }
      }
//console.log(VIXTWNArr)
//console.log( DateArr, TAIEXArr, VIXTWNArr, TAIPEArr);


      //第五步 回給 state 參數
      //小數點整理
      DateArr.forEach((_,k) => {
        TAIEXArr[k] = Number.parseFloat(TAIEXArr[k].toFixed(2))
        VIXTWNArr[k] = Number.parseFloat(VIXTWNArr[k].toFixed(2))
        TAIPEArr[k] = Number.parseFloat(TAIPEArr[k].toFixed(2))
      });
      //給予數值
      this.DateArr = DateArr;
      this.TAIEX = TAIEXArr;
      this.VIXTWN = VIXTWNArr;
      this.TAIPE = TAIPEArr;

      this.InitialStockReady = true;
    },


  },
})