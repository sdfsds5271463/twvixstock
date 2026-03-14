import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import NotFound from '../views/NotFound.vue'
import store from '../store'

import Layout from '../layout/index.vue'
import Dashboards from '../views/dashboards/Dashboards.vue'
import stockapi from '../views/stockapi/Stockapi.vue'
import programing from '../views/programing/Programing.vue'
import profile from '../views/profile/Profile.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/layout',
    name: 'Home',
    component: Layout,
    //component: () => import('../layout/index.vue'),
    redirect: '/',
    meta: {
      title: 'Home',
    },

    children: [
      {
        path: '/dashboards',
        alias: '/',
        name: 'Dashboards',
        component: Dashboards,  //直接載入
        //component: () => import('../views/dashboards/Dashboards.vue'),   //延遲載入
        meta: {
          title: 'Dashboard',
        },
      },
      {
        path: '/stockapi',
        name: 'stockapi',
        component: stockapi,  //直接載入
        //component: () => import('../views/stockapi/Stockapi.vue'),   //延遲載入
        meta: {
          title: 'Stock Api',
        },
      },
      {
        path: '/programing',
        name: 'programing',
        component: programing,  //直接載入
        //component: () => import('../views/programing/Programing.vue'),   //延遲載入
        meta: {
          title: 'Programing',
        },
      },
      {
        path: '/profile',
        name: 'profile',
        component: profile,  //直接載入
        //component: () => import('../views/profile/Profile.vue'),   //延遲載入
        meta: {
          title: 'Profile',
        },
      },
      /*{
        path: '/components',
        name: 'components',
        component: () => import('../views/components/Button.vue'),
        meta: {
          title: 'Components',
        },
      },*/
    ],
  },

  { path: '/:pathMatch(.*)*', component: NotFound },
]

const router = createRouter({
  history: createWebHistory(),
  routes,

  scrollBehavior() {
    return { left: 0, top: 0 }
  },
})

router.afterEach(() => {
  if (window.innerWidth <= 1200) {
    const sidenav = store.state.largeSidebar.sidebarToggleProperties.isSideNavOpen;
    if(sidenav){
      store.commit('largeSidebar/toggleSidebarProperties');
    }
  }
})

export default router