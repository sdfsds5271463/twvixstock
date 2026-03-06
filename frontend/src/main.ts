import { createApp } from 'vue'
import type { App as VueApp } from "vue"
import App from './App.vue'
import router from './router'
import store from './store'
import './assets/scss/global.scss'
import './index.css'
import BaseCard from './components/Base/BaseCard.vue'
import BaseBtn from './components/Base/BaseBtn.vue'

// perfectscrollbar plugins 
import PerfectScrollbar from 'vue3-perfect-scrollbar'
import 'vue3-perfect-scrollbar/dist/vue3-perfect-scrollbar.css'
import VueApexCharts from "vue3-apexcharts";

// create app
const app: VueApp = createApp(App)

// global components
app.component("BaseCard", BaseCard)
app.component("BaseBtn", BaseBtn)

// plugins
app.use(PerfectScrollbar)
app.use(VueApexCharts)
app.use(store)
app.use(router)

// mount
app.mount("#app")