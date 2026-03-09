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
  <div class="space-y-5 mb-5">
    <div class="text-2xl font-medium mt-5 max-sm:text-2xl max-sm:mt-3">图片上传</div>
    <div class="text-g-800">
      ArtImageUpload 是一个功能强大的图片上传组件，支持单图和多图上传，提供拖拽排序、预览、删除等功能。
    </div>

    <!-- 单图上传 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">单图上传</div>
      <div class="space-y-4">
        <div>
          <div class="text-sm text-g-700 mb-3">上传头像或LOGO：</div>
          <ArtImageUpload v-model="singleImage" :max-size="5" />
        </div>
        <div v-if="singleImage" class="text-sm text-g-600">
          图片地址：<span class="text-theme font-mono">{{ singleImage }}</span>
        </div>
      </div>
    </div>

    <!-- 多图上传 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">多图上传</div>
      <div class="space-y-4">
        <div>
          <div class="text-sm text-g-700 mb-3">上传轮播图或相册（支持拖拽排序）：</div>
          <ArtImageUpload v-model="multipleImages" multiple :limit="9" :max-size="5" />
        </div>
        <div v-if="multipleImages" class="text-sm text-g-600">
          已上传 {{ multipleImages.split(',').filter(url => url).length }} 张图片
        </div>
      </div>
    </div>

    <!-- 自定义限制 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">自定义限制</div>
      <div class="space-y-4">
        <div>
          <div class="text-sm text-g-700 mb-3">最多3张，每张最大2MB：</div>
          <ArtImageUpload v-model="limitedImages" multiple :limit="3" :max-size="2" />
        </div>
      </div>
    </div>

    <!-- 代码示例 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">代码示例</div>
      <div class="space-y-4">
        <div>
          <div class="text-sm text-g-600 mb-2">单图上传</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;template&gt;
  &lt;ArtImageUpload v-model="imageUrl" :max-size="5" /&gt;
&lt;/template&gt;

&lt;script setup lang="ts"&gt;
const imageUrl = ref('')
&lt;/script&gt;</code></pre>
          </div>
        </div>

        <div>
          <div class="text-sm text-g-600 mb-2">多图上传</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;template&gt;
  &lt;ArtImageUpload 
    v-model="images" 
    multiple 
    :limit="9"
    :max-size="5"
  /&gt;
&lt;/template&gt;

&lt;script setup lang="ts"&gt;
const images = ref('')
&lt;/script&gt;</code></pre>
          </div>
        </div>

        <div>
          <div class="text-sm text-g-600 mb-2">在表单中使用</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;ElFormItem label="商品图片"&gt;
  &lt;ArtImageUpload 
    v-model="form.images" 
    multiple 
    :limit="5"
  /&gt;
&lt;/ElFormItem&gt;</code></pre>
          </div>
        </div>
      </div>
    </div>

    <!-- 功能特性 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">功能特性</div>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
        <div class="flex items-start gap-2">
          <ArtSvgIcon icon="ri:checkbox-circle-fill" class="text-success mt-0.5" />
          <div class="text-sm text-g-700">单图/多图上传模式</div>
        </div>
        <div class="flex items-start gap-2">
          <ArtSvgIcon icon="ri:checkbox-circle-fill" class="text-success mt-0.5" />
          <div class="text-sm text-g-700">图片预览（点击放大）</div>
        </div>
        <div class="flex items-start gap-2">
          <ArtSvgIcon icon="ri:checkbox-circle-fill" class="text-success mt-0.5" />
          <div class="text-sm text-g-700">拖拽排序（多图模式）</div>
        </div>
        <div class="flex items-start gap-2">
          <ArtSvgIcon icon="ri:checkbox-circle-fill" class="text-success mt-0.5" />
          <div class="text-sm text-g-700">文件大小限制</div>
        </div>
        <div class="flex items-start gap-2">
          <ArtSvgIcon icon="ri:checkbox-circle-fill" class="text-success mt-0.5" />
          <div class="text-sm text-g-700">支持JPG/PNG/GIF/WebP</div>
        </div>
        <div class="flex items-start gap-2">
          <ArtSvgIcon icon="ri:checkbox-circle-fill" class="text-success mt-0.5" />
          <div class="text-sm text-g-700">悬停显示操作按钮</div>
        </div>
      </div>
    </div>

    <!-- API -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">API</div>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead class="bg-g-100 dark:bg-g-800">
            <tr>
              <th class="px-4 py-2 text-left">参数</th>
              <th class="px-4 py-2 text-left">说明</th>
              <th class="px-4 py-2 text-left">类型</th>
              <th class="px-4 py-2 text-left">默认值</th>
            </tr>
          </thead>
          <tbody>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">modelValue</td>
              <td class="px-4 py-2">图片URL（多图用逗号分隔）</td>
              <td class="px-4 py-2">string | string[]</td>
              <td class="px-4 py-2">''</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">multiple</td>
              <td class="px-4 py-2">是否多图模式</td>
              <td class="px-4 py-2">boolean</td>
              <td class="px-4 py-2">false</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">limit</td>
              <td class="px-4 py-2">最大上传数量</td>
              <td class="px-4 py-2">number</td>
              <td class="px-4 py-2">9</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">maxSize</td>
              <td class="px-4 py-2">最大文件大小（MB）</td>
              <td class="px-4 py-2">number</td>
              <td class="px-4 py-2">5</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">accept</td>
              <td class="px-4 py-2">接受的文件类型</td>
              <td class="px-4 py-2">string</td>
              <td class="px-4 py-2">'image/*'</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineOptions({ name: 'ImageUploadPage' })

const singleImage = ref('')
const multipleImages = ref('')
const limitedImages = ref('')
</script>
