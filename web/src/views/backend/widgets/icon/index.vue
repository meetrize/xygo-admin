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
    <div class="text-2xl font-medium mt-5 max-sm:text-2xl max-sm:mt-3">图标</div>
    <div class="text-g-800">
      v.3.0 版本图标库升级为 iconify，可在
      <a href="https://icones.js.org/" target="_blank" class="text-theme hover:underline"
        >Iconify</a
      >
      中查找，支持多种图标库，如 Remix Icon, Solar, Tabler Icons 等。
    </div>
    <div class="text-g-800">
      为确保系统图标风格统一，项目全部采用 Remix Icon 图标库，可在
      <a
        href="https://icones.js.org/collection/ri"
        target="_blank"
        class="text-theme hover:underline"
        >Iconify</a
      >
      或<a href="https://remixicon.com/" target="_blank" class="text-theme hover:underline">
        Remix Icon 官网 </a
      >搜索使用。
    </div>

    <!-- Iconify 图标 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">Iconify</div>
      <div class="flex items-center gap-6">
        <ArtSvgIcon icon="ri:github-fill" class="text-2xl" />
        <ArtSvgIcon icon="ri:copilot-line" class="text-2xl text-theme" />
        <ArtSvgIcon icon="ri:edge-line" class="text-2xl text-secondary" />
        <ArtSvgIcon icon="ri:planet-line" class="text-2xl text-warning" />
        <ArtSvgIcon icon="ri:windows-line" class="text-2xl text-info" />
        <ArtSvgIcon icon="ri:thumb-up-line" class="text-2xl text-danger" />
        <ArtSvgIcon icon="ri:gift-2-line" class="text-2xl text-success" />
        <ArtSvgIcon icon="ri:apple-line" class="text-2xl text-secondary" />
      </div>
    </div>

    <!-- Svg Icons -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">Svg Icons</div>
      <div class="flex items-center gap-6">
        <ArtSvgIcon icon="svg-spinners:3-dots-fade" class="text-2xl text-red-400" />
        <ArtSvgIcon icon="svg-spinners:3-dots-bounce" class="text-2xl text-blue-400" />
        <ArtSvgIcon icon="svg-spinners:3-dots-move" class="text-2xl text-orange-400" />
        <ArtSvgIcon icon="svg-spinners:3-dots-rotate" class="text-2xl text-purple-400" />
        <ArtSvgIcon icon="svg-spinners:blocks-shuffle-2" class="text-2xl text-pink-300" />
        <ArtSvgIcon icon="svg-spinners:clock" class="text-2xl text-yellow-500" />
        <ArtSvgIcon icon="svg-spinners:tadpole" class="text-2xl text-orange-500" />
        <ArtSvgIcon icon="svg-spinners:blocks-wave" class="text-2xl text-blue-500" />
      </div>
    </div>

    <!-- Svg Icons -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">Material Line Icons</div>
      <div class="flex items-center gap-6">
        <ArtSvgIcon icon="line-md:phone-call-twotone-loop" class="text-2xl text-blue-500" />

        <ArtSvgIcon icon="line-md:switch-off" class="text-2xl text-green-500" />
        <ArtSvgIcon icon="line-md:sun-rising-filled-loop" class="text-2xl text-yellow-400" />
        <ArtSvgIcon icon="line-md:volume-high-filled" class="text-2xl text-purple-500" />
        <ArtSvgIcon icon="line-md:github-twotone" class="text-2xl text-gray-700" />
        <ArtSvgIcon icon="line-md:telegram" class="text-2xl text-sky-500" />
        <ArtSvgIcon icon="line-md:reddit-loop" class="text-2xl text-orange-400" />
        <ArtSvgIcon
          icon="line-md:coffee-half-empty-filled-loop"
          class="text-2xl text-emerald-500"
        />
      </div>
    </div>

    <!-- 使用示例 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">使用示例</div>
      <div class="space-y-4">
        <div>
          <div class="text-sm text-g-600 mb-2">基础使用</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded font-mono text-sm text-g-800">
            &lt;ArtSvgIcon icon="ri:home-line" /&gt;
          </div>
        </div>
        <div>
          <div class="text-sm text-g-600 mb-2">自定义大小</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded font-mono text-sm text-g-800">
            &lt;ArtSvgIcon icon="ri:user-line" class="text-2xl" /&gt;
          </div>
        </div>
        <div>
          <div class="text-sm text-g-600 mb-2">自定义颜色</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded font-mono text-sm text-g-800">
            &lt;ArtSvgIcon icon="ri:heart-fill" class="text-red-500" /&gt;
          </div>
        </div>
        <div>
          <div class="text-sm text-g-600 mb-2">组合使用</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded font-mono text-sm text-g-800">
            &lt;ArtSvgIcon icon="ri:star-fill" class="text-4xl text-yellow-500" /&gt;
          </div>
        </div>
      </div>
    </div>

    <!-- 图标选择器组件 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">图标选择器组件</div>
      <div class="text-sm text-g-600 mb-4">
        提供可视化的图标选择界面，支持快速选择常用图标或自定义输入任意 Iconify 图标名称
      </div>
      
      <div class="space-y-6">
        <!-- 实际示例 -->
        <div>
          <div class="text-sm font-medium text-g-700 mb-3">在线示例</div>
          <div class="flex items-center gap-4 flex-wrap">
            <div class="flex items-center gap-2">
              <span class="text-sm text-g-600">选择的图标：</span>
              <ArtIconSelector v-model="selectedIcon" style="width: 300px;" />
            </div>
            <div v-if="selectedIcon" class="flex items-center gap-3 px-4 py-2 bg-g-100 dark:bg-g-800 rounded">
              <ArtSvgIcon :icon="selectedIcon" :size="32" class="text-theme" />
              <div class="text-sm">
                <div class="font-medium text-g-800">{{ selectedIcon }}</div>
                <div class="text-xs text-g-500">当前选中</div>
              </div>
            </div>
          </div>
        </div>

        <!-- 代码示例 -->
        <div>
          <div class="text-sm font-medium text-g-700 mb-3">基础使用</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;template&gt;
  &lt;ArtIconSelector v-model="iconValue" /&gt;
&lt;/template&gt;

&lt;script setup lang="ts"&gt;
const iconValue = ref('ri:home-line')
&lt;/script&gt;</code></pre>
          </div>
        </div>

        <!-- 功能特性 -->
        <div>
          <div class="text-sm font-medium text-g-700 mb-3">功能特性</div>
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

        <!-- 使用说明 -->
        <div>
          <div class="text-sm font-medium text-g-700 mb-3">使用说明</div>
          <div class="space-y-2 text-sm text-g-600">
            <div class="flex items-start gap-2">
              <span class="font-medium text-theme">快速选择：</span>
              <span>打开弹窗 → 搜索/浏览内置图标 → 点击选择 → 确定</span>
            </div>
            <div class="flex items-start gap-2">
              <span class="font-medium text-theme">自定义输入：</span>
              <span>打开弹窗 → 点击"自定义输入" → 输入图标名称（如 mdi:github）→ 预览 → 确定</span>
            </div>
            <div class="flex items-start gap-2">
              <span class="font-medium text-theme">支持图标库：</span>
              <span>Remix Icon (ri:)、Material Design (mdi:)、Font Awesome (fa:)、Heroicons (heroicons:) 等所有 Iconify 图标库</span>
            </div>
          </div>
        </div>

        <!-- 更多示例 -->
        <div>
          <div class="text-sm font-medium text-g-700 mb-3">更多示例</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;!-- 在表单中使用 --&gt;
&lt;ElFormItem label="图标"&gt;
  &lt;ArtIconSelector v-model="form.icon" /&gt;
&lt;/ElFormItem&gt;

&lt;!-- 禁用清空按钮 --&gt;
&lt;ArtIconSelector v-model="iconValue" :clearable="false" /&gt;

&lt;!-- 监听变化事件 --&gt;
&lt;ArtIconSelector 
  v-model="iconValue" 
  @change="handleIconChange"
/&gt;</code></pre>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ArtSvgIcon from '@/components/core/base/art-svg-icon/index.vue'

defineOptions({ name: 'IconPage' })

// 图标选择器示例
const selectedIcon = ref('ri:home-line')
</script>
