/*
 * Copyright (c) KylinSoft Co., Ltd.2021-2022. All rights reserved.
 * PilotGo is licensed under the Mulan PSL v2.
 * You can use this software accodring to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *     http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN 'AS IS' BASIS, WITHOUT WARRANTIES OF ANY KIND, 
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 * @Author: zhaozhenfang
 * @Date: 2022-02-25 16:33:46
 * @LastEditTime: 2022-05-18 17:45:30
 */
import router from './router'
import store from './store'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css'// progress bar style
import { getRoles, hasPermission } from '@/utils/auth'
NProgress.configure({ showSpinner: false })// NProgress Configuration
const whiteList = ['/login']

router.beforeEach((to, from, next) => {
  console.log(from, to)
  if (to.meta && to.meta.header_title) {
    document.title = to.meta.header_title
  }
  NProgress.start();
  if (getRoles()) {
    if (to.path === '/login') {
      next({ path: '/' })
      NProgress.done()
    } else {
      if (!store.getters.getMenus || store.getters.getMenus.length === 0) {
        store.dispatch('getPermission', store.getters.roles).then(res => {
          store.dispatch('GenerateRoutes').then(() => {
            next({ ...to, replace: true })
          })
        })
      } else {
        if (to.path === "/") {
          let paths = store.getters.getPaths;
          let keys = Object.keys(paths);
          let to = keys.length > 0 ? paths[keys[0]] : "/401"
          next({ path: to.path, replace: true })
        } else {
          if (to.name) {
            if (hasPermission(store.getters.getMenus, to)) {
              store.dispatch('SetActivePanel', to.meta.panel)
              next()
            } else {
              next({ path: '/404', replace: true })
            }
          } else {
            console.log('无name')
            router.selfaddRoutes();
            /* router.addRoute('Home', {
              path: '/plugin3',
              name: 'Plugin3',
              component: () => import('@/views/Plugin/plugin3.vue'),
              meta: {
                title: 'plugin', header_title: "插件管理3", panel: "log", icon_class: 'el-icon-s-order',
                breadcrumb: [
                  { name: '插件管理3' },
                ],
              }
            }) */
            next({ path: '/plugin3', replace: true })
          }
        }
      }
    }
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      next()
    } else {
      next('/login')
      NProgress.done()
    }
  }
})

router.afterEach(route => {
  NProgress.done();
})