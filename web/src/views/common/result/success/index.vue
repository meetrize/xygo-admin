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
  <ArtResultPage
    type="success"
    title="提交成功"
    message="提交结果页用于反馈一系列操作任务的处理结果，如果仅是简单操作，使用 Message 全局提示反馈即可。灰色区域可以显示一些补充的信息。"
    iconCode="ri:check-fill"
  >
    <template #content>
      <p>已提交申请，等待部门审核。</p>
    </template>
    <template #buttons>
      <ElButton type="primary" v-ripple>返回修改</ElButton>
      <ElButton v-ripple>查看</ElButton>
      <ElButton v-ripple>打印</ElButton>
    </template>
  </ArtResultPage>
</template>

<script setup lang="ts">
  defineOptions({ name: 'ResultSuccess' })
</script>
