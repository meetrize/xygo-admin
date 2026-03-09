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
    type="fail"
    title="提交失败"
    message="请核对并修改以下信息后，再重新提交。"
    iconCode="ri:close-fill"
  >
    <template #content>
      <p>您提交的内容有如下错误：</p>
      <p>
        <ArtSvgIcon icon="ri:close-circle-line" class="text-red-500 mr-1" />
        <span>您的账户已被冻结</span>
      </p>
      <p>
        <ArtSvgIcon icon="ri:close-circle-line" class="text-red-500 mr-1" />
        <span>您的账户还不具备申请资格</span>
      </p>
    </template>
    <template #buttons>
      <ElButton type="primary" v-ripple>返回修改</ElButton>
      <ElButton v-ripple>查看</ElButton>
    </template>
  </ArtResultPage>
</template>

<script setup lang="ts">
  defineOptions({ name: 'ResultFail' })
</script>
