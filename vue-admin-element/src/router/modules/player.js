/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const playerRouter = {
  path: '/player',
  component: Layout,
  redirect: '/system/menu1/menu1-1',
  name: 'Player',
  hidden: true,
  meta: {
    title: '播放器',
    icon: 'nested'
  },
  children: [
    {
      path: 'webrtc',
      name: 'Menu21',
      component: () => import('@/views/player/webrtc/index'),
      meta: { title: 'Webrtc', icon: 'el-icon-camera' }
    },
    {
      path: 'flvjs',
      name: 'Table4',
      component: () => import('@/views/player/flvjs/index'),
      meta: { title: 'FlvJs', icon: 'el-icon-video-camera'   }
    },
    
    /*
    {
      path: 'playbox',
      name: 'PlayBox',
      component: () => import('@/views/player/playbox/index'),
      meta: { title: '播放器', icon: 'el-icon-video-camera'   }
    }
    */
  ]
}



export default playerRouter
