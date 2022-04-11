import Vue from 'vue'
import Router from 'vue-router'
const _import = require('./_import')
import Home from '@/views/Home/Home'

Vue.use(Router)
const originalPush = Router.prototype.push
const originalReplace = Router.prototype.replace
Router.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
}
Router.prototype.replace = function replace(location) {
    return originalReplace.call(this, location).catch(err => err)
}
export const constantRouterMap = [
  { path: '/401', component: _import('errorPage/401') },
  { path: '/404', component: _import('errorPage/404') },
  { path: '/', component: _import('errorPage/404') },
]
export const routes = [
  {
    path: '/login',
    name: 'Login',
    component: _import('Login'),
    meta: { title: 'login', header_title: "登录", panel: "login" }
  },
  {
    path: '/home',
    component: Home,
    children: [
      {
        path: '/overview',
        name: 'Overview',
        component: _import('Overview/Overview'),
        meta: { title: 'overview', header_title: "概览", panel: "overview", icon_class: 'el-icon-s-help', icon_code:'&#xe612;' }
      },
      {
        path: '/cluster',
        name: 'Cluster',
        component:  _import('Cluster/Cluster'),
        meta: { title: 'cluster', header_title: "机器管理", panel: "cluster", icon_class: 'el-icon-s-platform' },
        children:[
          {
            path: '/firewall',
            name: 'Firewall',
            component:  _import('Firewall/Firewall'),
            meta: {  
              header_title: "防火墙配置", 
              panel: "cluster", 
              breadcrumb: [
                { name: '机器管理', path: '/cluster' },
                { name: '防火墙配置'}], 
            }
          },
          // {
          //   path: '/prometheus',
          //   name: 'Prometheus',
          //   component: _import('Prometheus/index'),
          //   meta: {
          //     header_title: "性能监控", 
          //     panel: "cluster", 
          //     breadcrumb: [
          //         { name: '机器管理', path: '/cluster' },
          //         { name: '性能监控'}
          //     ],
          //     icon_class: ''
          //   }
          // },
          {
            path: '/cluster:detail',
            name: 'MacDetail',
            component: _import('Cluster/detail/index'),
            meta: {
              header_title: "机器详情", 
              panel: "cluster", 
              breadcrumb: [
                  { name: '机器管理', path: '/cluster' },
                  { name: '机器详情'}
              ],
              icon_class: ''
            }
          },
        ]
      },
      {
        path: '/prometheus',
        name: 'Prometheus',
        component: _import('Prometheus/index'),
        meta: { title: 'prometheus', header_title: "性能监控", panel: "prometheus", icon_class: 'el-icon-s-marketing' }
      },
      {
        path: '/batch',
        component:  _import('Batch/Batch'),
        meta: { title: 'batch', header_title: "批次管理", panel: "batch", icon_class: 'el-icon-menu' },
        children:[
          {
            path: '/batch:id',
            name: 'BatchDetail',
            component: _import('Batch/detail/index'),
            meta: {
              header_title: "批次详情", 
              panel: "batch", 
              breadcrumb: [
                  { name: '批次管理', path: '/batch' },
                  { name: '批次详情'}
              ],
              icon_class: ''
            }
        },
        ]
      }, 
      {
        path: '/usermanager',
        name: 'UserManager',
        component:  _import('UserManager/UserMan'),
        meta: { title: 'usermanager', header_title: "用户管理", panel: "usermanager", icon_class: 'el-icon-user-solid' }
      },
      {
        path: '/rolemanager',
        name: 'RoleManager',
        component:  _import('RoleManager/RoleMan'),
        meta: { title: 'rolemanager', header_title: "角色管理", panel: "rolemanager", icon_class: 'el-icon-s-custom' }
      },
      /* {
        path: '/Message',
        name: 'Message',
        component:  _import('Message/message'),
        meta: { title: 'Message', header_title: "监控告警", panel: "message", icon_class: 'el-icon-s-promotion' }
      }, */
      {
        path: '/log',
        name: 'Log',
        component:  _import('Log/Log'),
        meta: { title: 'log', header_title: "日志管理", panel: "log", icon_class: 'el-icon-s-order' }
      },
      // {
      //   path: '', 
      //   redirect: '/overview'
      // },
    ]
  },
]

const router = new Router({
  mode: 'history',
  routes: [ ...routes, ...constantRouterMap],
})

export default router;
