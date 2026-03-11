import { Module } from 'vuex'

interface SidebarToggleProperties {
    isSideNavOpen: boolean;
}

interface largeSidebarState {
    sidebarToggleProperties: SidebarToggleProperties;
}

// Create a new store instance.
const largeSidebar:Module<largeSidebarState, any> = {

    namespaced: true,
    state: {
        sidebarToggleProperties: {
            isSideNavOpen: true,
        }
    },
    getters: {
        getSideBarToggleProperties: (state:largeSidebarState) => state.sidebarToggleProperties
    },
    // we cant use async code ---commit
    mutations: {
        toggleSidebarProperties: state =>
        (state.sidebarToggleProperties.isSideNavOpen = !state
          .sidebarToggleProperties.isSideNavOpen),
    },
}

// const app = createApp({ /* your root component */ })
export default largeSidebar;
