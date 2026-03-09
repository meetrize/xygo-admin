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
    <div class="text-2xl font-medium mt-5 max-sm:text-2xl max-sm:mt-3">文件上传</div>
    <div class="text-g-800">
      ArtFileUpload 是一个通用的文件上传组件，支持各种文件类型，提供智能文件图标识别和下载功能。
    </div>

    <!-- 单文件上传 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">单文件上传</div>
      <div class="space-y-4">
        <div>
          <div class="text-sm text-g-700 mb-3">上传文档或文件：</div>
          <ArtFileUpload v-model="singleFile" :max-size="10" />
        </div>
        <div v-if="singleFile" class="text-sm text-g-600">
          文件地址：<span class="text-theme font-mono">{{ singleFile }}</span>
        </div>
      </div>
    </div>

    <!-- 多文件上传 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">多文件上传</div>
      <div class="space-y-4">
        <div>
          <div class="text-sm text-g-700 mb-3">上传附件或资料：</div>
          <ArtFileUpload v-model="multipleFiles" multiple :limit="10" :max-size="10" />
        </div>
        <div v-if="multipleFiles" class="text-sm text-g-600">
          已上传 {{ multipleFiles.split(',').filter(url => url).length }} 个文件
        </div>
      </div>
    </div>

    <!-- 限制文件类型 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">限制文件类型</div>
      <div class="space-y-4">
        <div>
          <div class="text-sm text-g-700 mb-3">只允许上传 PDF 和 Word 文档：</div>
          <ArtFileUpload 
            v-model="pdfFile" 
            accept=".pdf,.doc,.docx" 
            :max-size="20" 
          />
        </div>
      </div>
    </div>

    <!-- 代码示例 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">代码示例</div>
      <div class="space-y-4">
        <div>
          <div class="text-sm text-g-600 mb-2">单文件上传</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;template&gt;
  &lt;ArtFileUpload v-model="fileUrl" :max-size="10" /&gt;
&lt;/template&gt;

&lt;script setup lang="ts"&gt;
const fileUrl = ref('')
&lt;/script&gt;</code></pre>
          </div>
        </div>

        <div>
          <div class="text-sm text-g-600 mb-2">多文件上传</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;template&gt;
  &lt;ArtFileUpload 
    v-model="files" 
    multiple 
    :limit="10"
    :max-size="10"
  /&gt;
&lt;/template&gt;

&lt;script setup lang="ts"&gt;
const files = ref('')
&lt;/script&gt;</code></pre>
          </div>
        </div>

        <div>
          <div class="text-sm text-g-600 mb-2">限制文件类型</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;ArtFileUpload 
  v-model="file" 
  accept=".pdf,.doc,.docx,.xls,.xlsx"
  :max-size="20"
/&gt;</code></pre>
          </div>
        </div>
      </div>
    </div>

    <!-- 支持的文件类型 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">支持的文件图标</div>
      <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4">
        <div class="flex flex-col items-center gap-2 p-3 bg-g-50 dark:bg-g-800 rounded">
          <ArtSvgIcon icon="ri:file-pdf-line" :size="32" class="text-red-500" />
          <span class="text-xs text-g-600">PDF</span>
        </div>
        <div class="flex flex-col items-center gap-2 p-3 bg-g-50 dark:bg-g-800 rounded">
          <ArtSvgIcon icon="ri:file-word-line" :size="32" class="text-blue-500" />
          <span class="text-xs text-g-600">Word</span>
        </div>
        <div class="flex flex-col items-center gap-2 p-3 bg-g-50 dark:bg-g-800 rounded">
          <ArtSvgIcon icon="ri:file-excel-line" :size="32" class="text-green-500" />
          <span class="text-xs text-g-600">Excel</span>
        </div>
        <div class="flex flex-col items-center gap-2 p-3 bg-g-50 dark:bg-g-800 rounded">
          <ArtSvgIcon icon="ri:file-ppt-line" :size="32" class="text-orange-500" />
          <span class="text-xs text-g-600">PPT</span>
        </div>
        <div class="flex flex-col items-center gap-2 p-3 bg-g-50 dark:bg-g-800 rounded">
          <ArtSvgIcon icon="ri:file-zip-line" :size="32" class="text-purple-500" />
          <span class="text-xs text-g-600">压缩包</span>
        </div>
        <div class="flex flex-col items-center gap-2 p-3 bg-g-50 dark:bg-g-800 rounded">
          <ArtSvgIcon icon="ri:file-code-line" :size="32" class="text-cyan-500" />
          <span class="text-xs text-g-600">代码</span>
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
              <td class="px-4 py-2">文件URL</td>
              <td class="px-4 py-2">string | string[]</td>
              <td class="px-4 py-2">''</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">multiple</td>
              <td class="px-4 py-2">是否多文件模式</td>
              <td class="px-4 py-2">boolean</td>
              <td class="px-4 py-2">false</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">limit</td>
              <td class="px-4 py-2">最大上传数量</td>
              <td class="px-4 py-2">number</td>
              <td class="px-4 py-2">10</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">maxSize</td>
              <td class="px-4 py-2">最大文件大小（MB）</td>
              <td class="px-4 py-2">number</td>
              <td class="px-4 py-2">10</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">accept</td>
              <td class="px-4 py-2">接受的文件类型</td>
              <td class="px-4 py-2">string</td>
              <td class="px-4 py-2">'*'</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineOptions({ name: 'FileUploadPage' })

const singleFile = ref('')
const multipleFiles = ref('')
const pdfFile = ref('')
</script>
