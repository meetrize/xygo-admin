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
    <div class="text-2xl font-medium mt-5 max-sm:text-2xl max-sm:mt-3">图标选择器</div>
    <div class="text-g-800">
      ArtIconSelector 是一个可视化的图标选择组件，支持快速选择常用图标或自定义输入任意 Iconify 图标名称。
    </div>

    <!-- 基础示例 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">基础示例</div>
      <div class="space-y-4">
        <div class="flex items-center gap-4">
          <span class="text-sm text-g-700 w-32">选择图标：</span>
          <ArtIconSelector v-model="selectedIcon1" style="width: 300px;" />
        </div>
        <div v-if="selectedIcon1" class="flex items-center gap-3 px-4 py-3 bg-g-100 dark:bg-g-800 rounded-lg">
          <ArtSvgIcon :icon="selectedIcon1" :size="32" class="text-theme" />
          <div class="text-sm">
            <div class="font-medium text-g-800">{{ selectedIcon1 }}</div>
            <div class="text-xs text-g-500">当前选中的图标</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 多个示例 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">多个图标选择</div>
      <div class="space-y-4">
        <div class="flex items-center gap-4">
          <span class="text-sm text-g-700 w-32">菜单图标：</span>
          <ArtIconSelector v-model="selectedIcon2" style="width: 300px;" />
        </div>
        <div class="flex items-center gap-4">
          <span class="text-sm text-g-700 w-32">按钮图标：</span>
          <ArtIconSelector v-model="selectedIcon3" style="width: 300px;" />
        </div>
        <div class="flex items-center gap-4">
          <span class="text-sm text-g-700 w-32">状态图标：</span>
          <ArtIconSelector v-model="selectedIcon4" style="width: 300px;" />
        </div>
      </div>
    </div>

    <!-- 图标展示 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">选中的图标展示</div>
      <div class="flex items-center gap-6">
        <div v-if="selectedIcon1" class="flex flex-col items-center gap-2">
          <ArtSvgIcon :icon="selectedIcon1" :size="48" class="text-theme" />
          <span class="text-xs text-g-600">大尺寸</span>
        </div>
        <div v-if="selectedIcon2" class="flex flex-col items-center gap-2">
          <ArtSvgIcon :icon="selectedIcon2" :size="32" class="text-success" />
          <span class="text-xs text-g-600">中尺寸</span>
        </div>
        <div v-if="selectedIcon3" class="flex flex-col items-center gap-2">
          <ArtSvgIcon :icon="selectedIcon3" :size="24" class="text-warning" />
          <span class="text-xs text-g-600">小尺寸</span>
        </div>
        <div v-if="selectedIcon4" class="flex flex-col items-center gap-2">
          <ArtSvgIcon :icon="selectedIcon4" :size="16" class="text-danger" />
          <span class="text-xs text-g-600">迷你尺寸</span>
        </div>
      </div>
    </div>

    <!-- 代码示例 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">代码示例</div>
      <div class="space-y-4">
        <div>
          <div class="text-sm text-g-600 mb-2">基础使用</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;template&gt;
  &lt;ArtIconSelector v-model="iconValue" /&gt;
&lt;/template&gt;

&lt;script setup lang="ts"&gt;
const iconValue = ref('ri:home-line')
&lt;/script&gt;</code></pre>
          </div>
        </div>

        <div>
          <div class="text-sm text-g-600 mb-2">在表单中使用</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;ElFormItem label="图标"&gt;
  &lt;ArtIconSelector v-model="form.icon" /&gt;
&lt;/ElFormItem&gt;</code></pre>
          </div>
        </div>

        <div>
          <div class="text-sm text-g-600 mb-2">禁用清空按钮</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;ArtIconSelector 
  v-model="iconValue" 
  :clearable="false" 
/&gt;</code></pre>
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
          <div class="text-sm text-g-700">可视化图标选择界面</div>
        </div>
        <div class="flex items-start gap-2">
          <ArtSvgIcon icon="ri:checkbox-circle-fill" class="text-success mt-0.5" />
          <div class="text-sm text-g-700">支持图标搜索功能</div>
        </div>
        <div class="flex items-start gap-2">
          <ArtSvgIcon icon="ri:checkbox-circle-fill" class="text-success mt-0.5" />
          <div class="text-sm text-g-700">内置 200+ 常用图标</div>
        </div>
        <div class="flex items-start gap-2">
          <ArtSvgIcon icon="ri:checkbox-circle-fill" class="text-success mt-0.5" />
          <div class="text-sm text-g-700">支持自定义输入任意图标</div>
        </div>
        <div class="flex items-start gap-2">
          <ArtSvgIcon icon="ri:checkbox-circle-fill" class="text-success mt-0.5" />
          <div class="text-sm text-g-700">实时预览图标效果</div>
        </div>
        <div class="flex items-start gap-2">
          <ArtSvgIcon icon="ri:checkbox-circle-fill" class="text-success mt-0.5" />
          <div class="text-sm text-g-700">支持所有 Iconify 图标库</div>
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
              <td class="px-4 py-2">当前选中的图标值</td>
              <td class="px-4 py-2">string</td>
              <td class="px-4 py-2">''</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">clearable</td>
              <td class="px-4 py-2">是否显示清空按钮</td>
              <td class="px-4 py-2">boolean</td>
              <td class="px-4 py-2">true</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineOptions({ name: 'IconSelectorPage' })

const selectedIcon1 = ref('ri:home-line')
const selectedIcon2 = ref('ri:menu-line')
const selectedIcon3 = ref('ri:add-line')
const selectedIcon4 = ref('ri:check-line')
</script>
