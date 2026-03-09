<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<template>
  <div class="page-content">
    <ArtCutterImg
      v-model:imgUrl="imageUrl"
      :boxWidth="530"
      :boxHeight="300"
      :cutWidth="360"
      :cutHeight="200"
      :quality="1"
      :tool="true"
      :watermarkText="'My Watermark'"
      watermarkColor="#ff0000"
      :showPreview="true"
      :originalGraph="false"
      :title="'图片裁剪'"
      :previewTitle="'预览效果'"
      @error="handleError"
      @imageLoadComplete="handleLoadComplete"
      @imageLoadError="handleLoadError"
    />
  </div>
</template>

<script setup lang="ts">
  import lockImg from '@imgs/lock/bg_dark.webp'

  defineOptions({ name: 'WidgetsImageCrop' })

  /**
   * 图片 URL
   */
  const imageUrl = ref(lockImg)

  /**
   * 处理裁剪错误
   * @param error 错误对象
   */
  const handleError = (error: Error) => {
    console.error('裁剪错误:', error)
    ElMessage.error('图片裁剪失败')
  }

  /**
   * 处理图片加载完成
   * @param result 加载结果
   */
  const handleLoadComplete = (result: { url: string; width: number; height: number }) => {
    console.log('图片加载完成:', result)
  }

  /**
   * 处理图片加载错误
   * @param error 错误对象
   */
  const handleLoadError = (error: Error) => {
    console.error('图片加载失败:', error)
    ElMessage.error('图片加载失败')
  }
</script>
