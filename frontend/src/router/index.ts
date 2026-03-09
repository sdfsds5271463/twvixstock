import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import NotFound from '../views/NotFound.vue'
import store from '../store'

const routes: RouteRecordRaw[] = [
  {
    path: '/layout',
    name: 'Home',
    component: () => import('../layout/index.vue'),
    redirect: '/',
    meta: {
      title: 'Home',
    },

    children: [
      {
        path: '/dashboards',
        alias: '/',
        name: 'Dashboards',
        component: () => import('../views/dashboards/Dashboards.vue'),
        meta: {
          title: 'Dashboard',
        },
      },
      {
        path: '/stockapi',
        name: 'stockapi',
        component: () => import('../views/stockapi/Stockapi.vue'),
        meta: {
          title: 'Stock Api',
        },
      },
      {
        path: '/programing',
        name: 'programing',
        component: () => import('../views/programing/Programing.vue'),
        meta: {
          title: 'Programing',
        },
      },
      {
        path: '/profile',
        name: 'profile',
        component: () => import('../views/profile/Profile.vue'),
        meta: {
          title: 'Profile',
        },
      },
      {
        path: '/components',
        name: 'components',
        component: () => import('../views/components/Button.vue'),
        meta: {
          title: 'Components',
        },
      },
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
    const sidenav =
      store.state.largeSidebar.sidebarToggleProperties.isSideNavOpen

    store.commit('largeSidebar/toggleSidebarProperties')
  }
})

export default router