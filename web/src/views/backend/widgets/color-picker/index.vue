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
    <div class="text-2xl font-medium mt-5 max-sm:text-2xl max-sm:mt-3">颜色选择器</div>
    <div class="text-g-800">
      ArtColorPicker 是一个美观实用的颜色选择组件，支持可视化颜色选择和手动输入颜色值。
    </div>

    <!-- 基础示例 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">基础示例</div>
      <div class="space-y-4">
        <div class="flex items-center gap-4">
          <span class="text-sm text-g-700 w-32">选择颜色：</span>
          <ArtColorPicker v-model="color1" />
        </div>
        <div v-if="color1" class="flex items-center gap-3">
          <div class="w-32 h-32 rounded-lg border-2 border-g-300" :style="{ background: color1 }"></div>
          <div class="text-sm">
            <div class="font-medium text-g-800">{{ color1 }}</div>
            <div class="text-xs text-g-500">当前选中的颜色</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 多个颜色选择 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">多个颜色选择</div>
      <div class="space-y-4">
        <div class="flex items-center gap-4">
          <span class="text-sm text-g-700 w-32">主题色：</span>
          <ArtColorPicker v-model="color2" placeholder="请选择主题色" />
        </div>
        <div class="flex items-center gap-4">
          <span class="text-sm text-g-700 w-32">成功色：</span>
          <ArtColorPicker v-model="color3" placeholder="请选择成功色" />
        </div>
        <div class="flex items-center gap-4">
          <span class="text-sm text-g-700 w-32">警告色：</span>
          <ArtColorPicker v-model="color4" placeholder="请选择警告色" />
        </div>
        <div class="flex items-center gap-4">
          <span class="text-sm text-g-700 w-32">危险色：</span>
          <ArtColorPicker v-model="color5" placeholder="请选择危险色" />
        </div>
      </div>
    </div>

    <!-- 颜色展示 -->
    <div class="art-card-sm p-5">
      <div class="text-lg font-semibold mb-4">颜色方案预览</div>
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div v-if="color2" class="flex flex-col items-center gap-3">
          <div class="w-full h-24 rounded-lg" :style="{ background: color2 }"></div>
          <div class="text-sm text-center">
            <div class="font-medium text-g-800">主题色</div>
            <div class="text-xs text-g-500">{{ color2 }}</div>
          </div>
        </div>
        <div v-if="color3" class="flex flex-col items-center gap-3">
          <div class="w-full h-24 rounded-lg" :style="{ background: color3 }"></div>
          <div class="text-sm text-center">
            <div class="font-medium text-g-800">成功色</div>
            <div class="text-xs text-g-500">{{ color3 }}</div>
          </div>
        </div>
        <div v-if="color4" class="flex flex-col items-center gap-3">
          <div class="w-full h-24 rounded-lg" :style="{ background: color4 }"></div>
          <div class="text-sm text-center">
            <div class="font-medium text-g-800">警告色</div>
            <div class="text-xs text-g-500">{{ color4 }}</div>
          </div>
        </div>
        <div v-if="color5" class="flex flex-col items-center gap-3">
          <div class="w-full h-24 rounded-lg" :style="{ background: color5 }"></div>
          <div class="text-sm text-center">
            <div class="font-medium text-g-800">危险色</div>
            <div class="text-xs text-g-500">{{ color5 }}</div>
          </div>
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
  &lt;ArtColorPicker v-model="color" /&gt;
&lt;/template&gt;

&lt;script setup lang="ts"&gt;
const color = ref('#409EFF')
&lt;/script&gt;</code></pre>
          </div>
        </div>

        <div>
          <div class="text-sm text-g-600 mb-2">在表单中使用</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;ElFormItem label="主题颜色"&gt;
  &lt;ArtColorPicker v-model="form.themeColor" /&gt;
&lt;/ElFormItem&gt;</code></pre>
          </div>
        </div>

        <div>
          <div class="text-sm text-g-600 mb-2">自定义占位符</div>
          <div class="bg-g-200 dark:bg-g-300/30 p-4 rounded">
            <pre class="font-mono text-sm text-g-800 whitespace-pre-wrap"><code>&lt;ArtColorPicker 
  v-model="color" 
  placeholder="请选择您喜欢的颜色"
/&gt;</code></pre>
          </div>
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
              <td class="px-4 py-2">当前颜色值</td>
              <td class="px-4 py-2">string</td>
              <td class="px-4 py-2">''</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">placeholder</td>
              <td class="px-4 py-2">占位文本</td>
              <td class="px-4 py-2">string</td>
              <td class="px-4 py-2">'请选择颜色'</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">showAlpha</td>
              <td class="px-4 py-2">是否支持透明度</td>
              <td class="px-4 py-2">boolean</td>
              <td class="px-4 py-2">false</td>
            </tr>
            <tr class="border-b border-g-200 dark:border-g-700">
              <td class="px-4 py-2 font-mono text-theme">colorFormat</td>
              <td class="px-4 py-2">颜色格式</td>
              <td class="px-4 py-2">'hex' | 'rgb' | 'hsl' | 'hsv'</td>
              <td class="px-4 py-2">'hex'</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineOptions({ name: 'ColorPickerPage' })

const color1 = ref('#409EFF')
const color2 = ref('#409EFF')
const color3 = ref('#67C23A')
const color4 = ref('#E6A23C')
const color5 = ref('#F56C6C')
</script>
