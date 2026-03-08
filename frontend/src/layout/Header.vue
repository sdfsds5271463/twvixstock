<script setup>
import { onMounted, ref, watch, watchEffect } from 'vue'
import { Switch } from '@headlessui/vue'
//import HeaderSearch from '../components/HeaderSearch.vue'
import { Menu, MenuButton, MenuItems, MenuItem } from '@headlessui/vue'
import { useStore } from 'vuex'

let store = useStore()

const sideBarToggle = () => {
    let sidenav = store.state.largeSidebar.sidebarToggleProperties.isSideNavOpen

    if (sidenav == false) {
        store.commit('largeSidebar/toggleSidebarProperties')
    } else {
        store.commit('largeSidebar/toggleSidebarProperties')
    }
}
</script>

<template>
    <div class="header-wrapper flex bg-white justify-between px-4">
        <div class="flex items-center">
            <router-link class="logo flex justify-center" to="/">
                <img src="@/../public/images/a_icon.png">
                <span class="text-3xl pl-1 text-primary">TwVix</span>
            </router-link>
            <div class="mx-0">
                <button
                    @click="sideBarToggle"
                    class="
                        menu-toggle
                        cursor-pointer
                        text-gray-500
                        align-middle
                        focus:outline-none
                        p-3
                    "
                    type="button"
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        class="h-6 w-6"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M4 6h16M4 12h16M4 18h16"
                        />
                    </svg>
                </button>
            </div>
            <!-- mega-menu  -->

            <!-- hidden 隱藏起來 20260308-->
            <Menu as="div" class="hidden relative inline-block text-left">
                <div>
                    <MenuButton
                        class="
                            btn-sm
                            mx-3
                            md:block
                            hidden
                            focus:outline-none
                            dropdown-button
                        "
                    >
                        <div class="flex items-center">
                            <p class="mr-2">Mega Menu</p>
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                class="h-4 w-4 text-gray-500"
                                fill="none"
                                viewBox="0 0 24 24"
                                stroke="currentColor"
                            >
                                <path
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    stroke-width="2"
                                    d="M19 9l-7 7-7-7"
                                />
                            </svg>
                        </div>
                    </MenuButton>
                </div>

                <MenuItems
                    class="
                        absolute
                        mega-menu
                        left-0
                        mt-2
                        overflow-hidden
                        origin-top-right
                        bg-white
                        rounded-md
                        custom-box-shadow
                        focus:outline-none
                    "
                >
                    <div class="">
                        <MenuItem>
                            <div class="grid grid-cols-12 gap-5">
                                <div
                                    class="
                                        col-span-4
                                        bg-gradient-to-r
                                        from-purple-500
                                        to-purple-700
                                    "
                                >
                                    <div class="p-4">
                                        <p
                                            class="
                                                text-2xl text-white
                                                font-semibold
                                                mb-3
                                            "
                                        >
                                            Mega Menu Sidebar
                                        </p>
                                        <p class="text-white mb-4">
                                            Lorem ipsum dolor sit amet
                                            consectetur, adipisicing elit.
                                            Asperiores natus laboriosam fugit,
                                            consequatur.
                                        </p>
                                        <p class="text-white mb-4">
                                            Lorem ipsum dolor sit amet
                                            consectetur, adipisicing elit.
                                            Exercitationem odio amet eos dolore
                                            suscipit placeat.
                                        </p>
                                        <BaseBtn
                                            class="
                                                border border-yellow-500
                                                text-yellow-500
                                                hover:bg-yellow-500
                                                hover:text-white
                                            "
                                            rounded
                                            >Learn More</BaseBtn
                                        >
                                    </div>
                                </div>
                                <div class="col-span-4">
                                    <div class="p-4">
                                        <p
                                            class="
                                                text-purple-500
                                                border-b border-purple-500
                                                inline-block
                                                mb-4
                                            "
                                        >
                                            Features
                                        </p>
                                        <div
                                            class="
                                                menu-icon-grid
                                                flex
                                                justify-between
                                                flex-wrap
                                            "
                                        >
                                            <a
                                                class="
                                                    hover:bg-purple-500
                                                    hover:text-white
                                                    text-gray-600
                                                    p-5
                                                    inline-flex
                                                    flex-col
                                                    items-center
                                                    rounded
                                                    w-24
                                                "
                                                href="#"
                                            >
                                                <i class="i-Shop-4 text-2xl">
                                                </i
                                                >Home
                                            </a>
                                            <a
                                                class="
                                                    hover:bg-purple-500
                                                    hover:text-white
                                                    text-gray-600
                                                    p-5
                                                    inline-flex
                                                    flex-col
                                                    items-center
                                                    rounded
                                                    w-24
                                                "
                                                href="#"
                                            >
                                                <i
                                                    class="i-Library text-2xl"
                                                ></i
                                                >UI Kits
                                            </a>
                                            <a
                                                class="
                                                    hover:bg-purple-500
                                                    hover:text-white
                                                    text-gray-600
                                                    p-5
                                                    inline-flex
                                                    flex-col
                                                    items-center
                                                    rounded
                                                    w-24
                                                "
                                                href="#"
                                            >
                                                <i class="i-Drop text-2xl"></i
                                                >Apps
                                            </a>
                                            <a
                                                class="
                                                    hover:bg-purple-500
                                                    hover:text-white
                                                    text-gray-600
                                                    p-5
                                                    inline-flex
                                                    flex-col
                                                    items-center
                                                    rounded
                                                    w-24
                                                "
                                                href="#"
                                            >
                                                <i
                                                    class="
                                                        i-File-Clipboard-File--Text
                                                        text-2xl
                                                    "
                                                >
                                                </i
                                                >Forms
                                            </a>
                                            <a
                                                class="
                                                    hover:bg-purple-500
                                                    hover:text-white
                                                    text-gray-600
                                                    p-5
                                                    inline-flex
                                                    flex-col
                                                    items-center
                                                    rounded
                                                    w-24
                                                "
                                                href="#"
                                            >
                                                <i
                                                    class="
                                                        i-Checked-User
                                                        text-2xl
                                                    "
                                                ></i
                                                >Sessions
                                            </a>
                                            <a
                                                class="
                                                    hover:bg-purple-500
                                                    hover:text-white
                                                    text-gray-600
                                                    p-5
                                                    inline-flex
                                                    flex-col
                                                    items-center
                                                    rounded
                                                    w-24
                                                "
                                                href="#"
                                            >
                                                <i
                                                    class="i-Ambulance text-2xl"
                                                ></i
                                                >Support
                                            </a>
                                        </div>
                                    </div>
                                </div>
                                <div class="col-span-4">
                                    <div class="p-4">
                                        <p
                                            class="
                                                text-purple-500
                                                border-b border-purple-500
                                                inline-block
                                                mb-4
                                            "
                                        >
                                            Components
                                        </p>
                                        <ul class="links">
                                            <li>
                                                <a
                                                    class="
                                                        hover:text-purple-500
                                                        text-gray-600
                                                    "
                                                    href="#"
                                                    >Accordion
                                                </a>
                                            </li>
                                            <li>
                                                <a
                                                    class="
                                                        hover:text-purple-500
                                                        text-gray-600
                                                    "
                                                    href="#"
                                                    >Alerts
                                                </a>
                                            </li>
                                            <li>
                                                <a
                                                    class="
                                                        hover:text-purple-500
                                                        text-gray-600
                                                    "
                                                    href="#"
                                                    >Buttons
                                                </a>
                                            </li>
                                            <li>
                                                <a
                                                    class="
                                                        hover:text-purple-500
                                                        text-gray-600
                                                    "
                                                    href="#"
                                                    >Badges
                                                </a>
                                            </li>
                                            <li>
                                                <a
                                                    class="
                                                        hover:text-purple-500
                                                        text-gray-600
                                                    "
                                                    href="#"
                                                    >Carousels
                                                </a>
                                            </li>
                                            <li>
                                                <a
                                                    class="
                                                        hover:text-purple-500
                                                        text-gray-600
                                                    "
                                                    href="#"
                                                    >Lists
                                                </a>
                                            </li>
                                            <li>
                                                <a
                                                    class="
                                                        hover:text-purple-500
                                                        text-gray-600
                                                    "
                                                    href="#"
                                                    >Popover
                                                </a>
                                            </li>
                                            <li>
                                                <a
                                                    class="
                                                        hover:text-purple-500
                                                        text-gray-600
                                                    "
                                                    href="#"
                                                    >Tables
                                                </a>
                                            </li>
                                            <li>
                                                <a
                                                    class="
                                                        hover:text-purple-500
                                                        text-gray-600
                                                    "
                                                    href="#"
                                                    >Datatables
                                                </a>
                                            </li>
                                            <li>
                                                <a
                                                    class="
                                                        hover:text-purple-500
                                                        text-gray-600
                                                    "
                                                    href="#"
                                                    >Modals
                                                </a>
                                            </li>
                                            <li>
                                                <a
                                                    class="
                                                        hover:text-purple-500
                                                        text-gray-600
                                                    "
                                                    href="#"
                                                    >Sliders
                                                </a>
                                            </li>
                                            <li>
                                                <a
                                                    class="
                                                        hover:text-purple-500
                                                        text-gray-600
                                                    "
                                                    href="#"
                                                    >Tabs
                                                </a>
                                            </li>
                                        </ul>
                                    </div>
                                </div>
                            </div>
                        </MenuItem>
                    </div>
                </MenuItems>
            </Menu>

            <!--HeaderSearch /-->
        </div>
        <div class="flex items-center">
            <a
                class="
                    hidden
                    mr-0
                    btn btn-sm
                    hover:bg-gray-100
                    sm:block
                    rounded
                    p-1
                    md:mr-2
                "
                href="https://github.com/sdfsds5271463/twvixstock"
                target="_blank"
            >
                <svg
                    width="24"
                    height="24"
                    fill="currentColor"
                    class="text-purple-600 text-opacity-50 transform"
                >
                    <path
                        fill-rule="evenodd"
                        clip-rule="evenodd"
                        d="M12 2C6.477 2 2 6.463 2 11.97c0 4.404 2.865 8.14 6.839 9.458.5.092.682-.216.682-.48 0-.236-.008-.864-.013-1.695-2.782.602-3.369-1.337-3.369-1.337-.454-1.151-1.11-1.458-1.11-1.458-.908-.618.069-.606.069-.606 1.003.07 1.531 1.027 1.531 1.027.892 1.524 2.341 1.084 2.91.828.092-.643.35-1.083.636-1.332-2.22-.251-4.555-1.107-4.555-4.927 0-1.088.39-1.979 1.029-2.675-.103-.252-.446-1.266.098-2.638 0 0 .84-.268 2.75 1.022A9.606 9.606 0 0112 6.82c.85.004 1.705.114 2.504.336 1.909-1.29 2.747-1.022 2.747-1.022.546 1.372.202 2.386.1 2.638.64.696 1.028 1.587 1.028 2.675 0 3.83-2.339 4.673-4.566 4.92.359.307.678.915.678 1.846 0 1.332-.012 2.407-.012 2.734 0 .267.18.577.688.48C19.137 20.107 22 16.373 22 11.969 22 6.463 17.522 2 12 2z"
                    ></path>
                </svg>
            </a>

            <!-- hidden 隱藏起來 20260308-->
            <Menu as="div" class="hidden relative inline-block text-left">
                <div>
                    <MenuButton
                        class="
                            btn btn-sm
                            hover:bg-gray-100
                            rounded
                            w-full
                            px-4
                            py-2
                            text-sm
                            font-medium
                        "
                    >
                        <i class="i-Safe-Box text-gray-800 text-xl"></i>
                    </MenuButton>
                </div>

                <MenuItems
                    class="
                        absolute
                        right-0
                        w-40
                        mt-2
                        overflow-hidden
                        origin-top-right
                        bg-white
                        rounded-md
                        custom-box-shadow
                        focus:outline-none
                    "
                >
                    <div class="">
                        <MenuItem>
                            <div class="flex justify-between flex-wrap">
                                <a
                                    class="
                                        hover:bg-purple-500 hover:text-white
                                        text-gray-600
                                        p-4
                                        inline-flex
                                        flex-col
                                        items-center
                                        rounded
                                        w-20
                                    "
                                    href="#"
                                >
                                    <i class="i-Shop-4 text-2xl"> </i>Home
                                </a>
                                <a
                                    class="
                                        hover:bg-purple-500 hover:text-white
                                        text-gray-600
                                        p-4
                                        inline-flex
                                        flex-col
                                        items-center
                                        rounded
                                        w-20
                                    "
                                    href="#"
                                >
                                    <i class="i-Library text-2xl"></i>UI Kits
                                </a>
                                <a
                                    class="
                                        hover:bg-purple-500 hover:text-white
                                        text-gray-600
                                        p-4
                                        inline-flex
                                        flex-col
                                        items-center
                                        rounded
                                        w-20
                                    "
                                    href="#"
                                >
                                    <i class="i-Drop text-2xl"></i>Apps
                                </a>
                                <a
                                    class="
                                        hover:bg-purple-500 hover:text-white
                                        text-gray-600
                                        p-4
                                        inline-flex
                                        flex-col
                                        items-center
                                        rounded
                                        w-20
                                    "
                                    href="#"
                                >
                                    <i
                                        class="
                                            i-File-Clipboard-File--Text
                                            text-2xl
                                        "
                                    >
                                    </i
                                    >Forms
                                </a>
                                <a
                                    class="
                                        hover:bg-purple-500 hover:text-white
                                        text-gray-600
                                        p-4
                                        inline-flex
                                        flex-col
                                        items-center
                                        rounded
                                        w-20
                                    "
                                    href="#"
                                >
                                    <i class="i-Checked-User text-2xl"></i
                                    >Sessions
                                </a>
                                <a
                                    class="
                                        hover:bg-purple-500 hover:text-white
                                        text-gray-600
                                        p-4
                                        inline-flex
                                        flex-col
                                        items-center
                                        rounded
                                        w-20
                                    "
                                    href="#"
                                >
                                    <i class="i-Ambulance text-2xl"></i>Support
                                </a>
                            </div>
                        </MenuItem>
                    </div>
                </MenuItems>
            </Menu>

            <!-- hidden 隱藏起來 20260308-->
            <!-- notification-dropdown  -->
            <Menu as="div" class="hidden relative inline-block text-left">
                <div>
                    <MenuButton
                        class="
                            btn btn-sm
                            hover:bg-gray-100
                            rounded
                            badge-top-container
                            inline-flex
                            justify-center
                            w-full
                            px-4
                            py-2
                            text-sm
                            font-medium
                        "
                    >
                        <span class="badge-count text-white bg-purple-500"
                            >3</span
                        >
                        <i class="i-Bell text-xl header-icon text-gray-800"></i>
                    </MenuButton>
                </div>

                <MenuItems
                    class="
                        absolute
                        right-0
                        w-44
                        mt-2
                        overflow-hidden
                        origin-top-right
                        bg-white
                        rounded-md
                        custom-box-shadow
                        focus:outline-none
                    "
                >
                    <div class="">
                        <MenuItem
                            v-for="(item, index) in 4"
                            :key="index"
                            v-slot="{ active }"
                        >
                            <button
                                :class="[
                                    active
                                        ? 'bg-purple-500 text-white'
                                        : 'text-gray-900',
                                    'group flex  items-center w-full px-4 py-2 text-sm',
                                ]"
                            >
                                <div class="flex flex-1 justify-between">
                                    <div>
                                        <h6>New Message</h6>
                                        <p
                                            :class="[
                                                active
                                                    ? 'text-gray-300'
                                                    : 'text-gray-500',
                                            ]"
                                        >
                                            How are you ?
                                        </p>
                                    </div>
                                    <div>
                                        <p
                                            :class="[
                                                active
                                                    ? 'text-gray-200'
                                                    : 'text-gray-400',
                                            ]"
                                        >
                                            10s
                                        </p>
                                    </div>
                                </div>
                            </button>
                        </MenuItem>
                    </div>
                </MenuItems>
            </Menu>

            <!-- hidden 隱藏起來 20260308-->
            <!-- profile-dropdown  -->
            <Menu as="div" class="hidden relative inline-block text-left">
                <div>
                    <MenuButton
                        class="
                            inline-flex
                            justify-center
                            w-full
                            px-4
                            py-2
                            text-sm
                            font-medium
                        "
                    >
                        <img
                            class="avatar rounded-full"
                            src="/images/faces/1.jpg"
                            alt=""
                        />
                    </MenuButton>
                </div>

                <MenuItems
                    class="
                        absolute
                        right-0
                        w-44
                        mt-2
                        overflow-hidden
                        origin-top-right
                        bg-white
                        rounded-md
                        custom-box-shadow
                        focus:outline-none
                    "
                >
                    <div class="">
                        <MenuItem v-slot="{ active }">
                            <button
                                :class="[
                                    active
                                        ? 'bg-purple-500 text-white'
                                        : 'text-gray-900',
                                    'group flex  items-center w-full px-4 py-2 text-sm',
                                ]"
                            >
                                Account Settings
                            </button>
                        </MenuItem>
                        <MenuItem v-slot="{ active }">
                            <button
                                :class="[
                                    active
                                        ? 'bg-purple-500 text-white'
                                        : 'text-gray-900',
                                    'group flex  items-center w-full px-4 py-2 text-sm',
                                ]"
                            >
                                Billing History
                            </button>
                        </MenuItem>
                        <MenuItem v-slot="{ active }">
                            <button
                                :class="[
                                    active
                                        ? 'bg-purple-500 text-white'
                                        : 'text-gray-900',
                                    'group flex  items-center w-full px-4 py-2 text-sm',
                                ]"
                            >
                                Sign Out
                            </button>
                        </MenuItem>
                    </div>
                </MenuItems>
            </Menu>
        </div>
    </div>
</template>

<style lang="scss" scoped>
.header-wrapper {
    position: fixed;
    top: 0;
    width: 100%;
    height: 50px;
    z-index: 100;
    box-shadow: 0 1px 15px rgb(0 0 0 / 4%), 0 1px 6px rgb(0 0 0 / 4%);
}
.mega-menu {
    width: 1200px;
}
ul.links {
    column-count: 2;
    li {
        margin-bottom: 8px;
    }
}
</style>
