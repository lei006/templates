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


  /*  
  {
    path: '/pdf2file',
    component: Layout,
    redirect: '/pdf2file',
    children: [{
      path: 'pdf2file',
      name: 'Pdf2file',
      component: () => import('@/views/pdf2file/index'),
      meta: { title: 'PDF转文件', icon: 'dashboard' }
    }]
  },
  {
    path: '/file2pdf',
    component: Layout,
    redirect: '/file2pdf',
    children: [{
      path: 'file2pdf',
      name: 'File2pdf',
      component: () => import('@/views/file2pdf/index'),
      meta: { title: 'PDF转文件', icon: 'dashboard' }
    }]
  },
  {
    path: '/pdf2tool',
    component: Layout,
    redirect: '/pdf2tool',
    children: [{
      path: 'pdf2tool',
      name: 'Pdf2tool',
      component: () => import('@/views/pdf2tool/index'),
      meta: { title: 'PDF工具', icon: 'dashboard' }
    }]
  },


  */
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


