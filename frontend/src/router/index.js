import { createRouter, createWebHistory } from 'vue-router'
import NotFound from '../views/NotFound.vue'
import store from '../store'
import SignIn from '../views/sessions/SignIn.vue'
import SignUp from '../views/sessions/SignUp.vue'

const routes = [
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
                component: () => import('../views/dashboards/Dashboards.v1.vue'),
                meta: {
                    title: 'Dashboard',
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
            {
                path: '/profile',
                name: 'profile',
                component: () => import('../views/profile/ProfileTwo.vue'),
                meta: {
                    title: 'Profile',
                },
            },
        ],
    },

    { path: '/signIn', component: SignIn },
    { path: '/signUp', component: SignUp },

    { path: '/:path(.*)', component: NotFound },
]

const router = createRouter({
    history: createWebHistory(),
    scrollBehavior(to, from, savedPosition) {
        return { left: 0, top: 0 }
    },
    routes,
})

router.afterEach(() => {
    if (window.innerWidth <= 1200) {
        const sidenav =
            store.state.largeSidebar.sidebarToggleProperties.isSideNavOpen

        store.commit('largeSidebar/toggleSidebarProperties')
    }
})

export default router
