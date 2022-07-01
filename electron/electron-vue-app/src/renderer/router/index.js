import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

import Layout from '@/layout'
/*
export default new Router({
  routes: [
    {
      path: '/',
      name: 'landing-page',
      component: require('@/components/LandingPage').default
    },
    {
      path: '*',
      redirect: '/'
    }
  ]
})
*/


export const constantRoutes = [
  /*
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    name: 'landing-page',
    component: require('@/components/LandingPage').default
  },
  */
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '概要', icon: 'dashboard' }
    }]
  },

  {
    path: '/page01',
    component: Layout,
    redirect: '/page01',
    children: [{
      path: 'page01',
      name: 'Page01',
      component: () => import('@/views/page01/index'),
      meta: { title: '页面一', icon: 'dashboard' }
    }]
  },
  {
    path: '/page02',
    component: Layout,
    redirect: '/page02',
    children: [{
      path: 'page02',
      name: 'Page02',
      component: () => import('@/views/page02/index'),
      meta: { title: '页面二', icon: 'dashboard' }
    }]
  },


  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/', hidden: true },
  //{ path: '*', redirect: '/404', hidden: true },
]



const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})



const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router


