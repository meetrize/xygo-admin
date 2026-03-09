// +----------------------------------------------------------------------
// | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
// +----------------------------------------------------------------------
// | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( https://opensource.org/licenses/MIT )
// +----------------------------------------------------------------------
// | Author: 喜羊羊 <751300685@qq.com>
// +----------------------------------------------------------------------

import { AppRouteRecord } from '@/types/router'

export const widgetsRoutes: AppRouteRecord = {
  path: '/widgets',
  name: 'Widgets',
  component: '/index/index',
  meta: {
    title: 'menus.widgets.title',
    icon: 'ri:apps-2-add-line'
  },
  children: [
    {
      path: 'icon',
      name: 'Icon',
      component: '/widgets/icon',
      meta: {
        title: 'menus.widgets.icon',
        icon: 'ri:palette-line',
        keepAlive: true
      }
    },
    {
      path: 'image-crop',
      name: 'ImageCrop',
      component: '/widgets/image-crop',
      meta: {
        title: 'menus.widgets.imageCrop',
        icon: 'ri:screenshot-line',
        keepAlive: true
      }
    },
    {
      path: 'excel',
      name: 'Excel',
      component: '/widgets/excel',
      meta: {
        title: 'menus.widgets.excel',
        icon: 'ri:download-2-line',
        keepAlive: true
      }
    },
    {
      path: 'video',
      name: 'Video',
      component: '/widgets/video',
      meta: {
        title: 'menus.widgets.video',
        icon: 'ri:vidicon-line',
        keepAlive: true
      }
    },
    {
      path: 'count-to',
      name: 'CountTo',
      component: '/widgets/count-to',
      meta: {
        title: 'menus.widgets.countTo',
        icon: 'ri:anthropic-line',
        keepAlive: false
      }
    },
    {
      path: 'wang-editor',
      name: 'WangEditor',
      component: '/widgets/wang-editor',
      meta: {
        title: 'menus.widgets.wangEditor',
        icon: 'ri:t-box-line',
        keepAlive: true
      }
    },
    {
      path: 'watermark',
      name: 'Watermark',
      component: '/widgets/watermark',
      meta: {
        title: 'menus.widgets.watermark',
        icon: 'ri:water-flash-line',
        keepAlive: true
      }
    },
    {
      path: 'context-menu',
      name: 'ContextMenu',
      component: '/widgets/context-menu',
      meta: {
        title: 'menus.widgets.contextMenu',
        icon: 'ri:menu-2-line',
        keepAlive: true
      }
    },
    {
      path: 'qrcode',
      name: 'Qrcode',
      component: '/widgets/qrcode',
      meta: {
        title: 'menus.widgets.qrcode',
        icon: 'ri:qr-code-line',
        keepAlive: true
      }
    },
    {
      path: 'drag',
      name: 'Drag',
      component: '/widgets/drag',
      meta: {
        title: 'menus.widgets.drag',
        icon: 'ri:drag-move-fill',
        keepAlive: true
      }
    },
    {
      path: 'text-scroll',
      name: 'TextScroll',
      component: '/widgets/text-scroll',
      meta: {
        title: 'menus.widgets.textScroll',
        icon: 'ri:input-method-line',
        keepAlive: true
      }
    },
    {
      path: 'fireworks',
      name: 'Fireworks',
      component: '/widgets/fireworks',
      meta: {
        title: 'menus.widgets.fireworks',
        icon: 'ri:magic-line',
        keepAlive: true,
        showTextBadge: 'Hot'
      }
    },
    {
      path: 'icon-selector',
      name: 'IconSelector',
      component: '/widgets/icon-selector',
      meta: {
        title: 'menus.widgets.iconSelector',
        icon: 'ri:palette-line',
        keepAlive: true
      }
    },
    {
      path: 'color-picker',
      name: 'ColorPicker',
      component: '/widgets/color-picker',
      meta: {
        title: 'menus.widgets.colorPicker',
        icon: 'ri:palette-fill',
        keepAlive: true
      }
    },
    {
      path: 'image-upload',
      name: 'ImageUpload',
      component: '/widgets/image-upload',
      meta: {
        title: 'menus.widgets.imageUpload',
        icon: 'ri:image-2-line',
        keepAlive: true
      }
    },
    {
      path: 'file-upload',
      name: 'FileUpload',
      component: '/widgets/file-upload',
      meta: {
        title: 'menus.widgets.fileUpload',
        icon: 'ri:file-upload-line',
        keepAlive: true
      }
    },
    {
      path: 'array-editor',
      name: 'ArrayEditor',
      component: '/widgets/array-editor',
      meta: {
        title: 'menus.widgets.arrayEditor',
        icon: 'ri:list-settings-line',
        keepAlive: true
      }
    },
    {
      path: '/outside/iframe/elementui',
      name: 'ElementUI',
      component: '',
      meta: {
        title: 'menus.widgets.elementUI',
        icon: 'ri:apps-2-line',
        keepAlive: false,
        link: 'https://element-plus.org/zh-CN/component/overview.html',
        isIframe: true
      }
    }
  ]
}
