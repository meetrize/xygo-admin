<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- 代码预览面板 -->
<template>
  <div class="preview-panel">
    <!-- 操作栏 -->
    <div class="preview-actions">
      <ElButton type="primary" @click="handlePreview" :loading="previewLoading">
        <ArtSvgIcon icon="ri:eye-line" class="text-sm mr-1.5" />
        预览代码
      </ElButton>
      <ElButton type="success" @click="handleBuild" :loading="buildLoading" :disabled="!previewFiles.length">
        <ArtSvgIcon icon="ri:rocket-line" class="text-sm mr-1.5" />
        {{ buildStatusText || '执行生成' }}
      </ElButton>
    </div>

    <!-- 预览内容 -->
    <div v-if="previewFiles.length" class="preview-content">
      <ElTabs v-model="activeTab" type="border-card" class="preview-tabs">
        <ElTabPane
          v-for="file in previewFiles"
          :key="file.path"
          :label="getFileName(file.path)"
          :name="file.path"
        >
          <div class="file-path-bar">
            <ArtSvgIcon icon="ri:file-code-line" class="text-sm mr-1.5" />
            <span>{{ file.path }}</span>
            <ElButton size="small" text class="ml-auto" @click="copyCode(file.content)">
              <ArtSvgIcon icon="ri:file-copy-line" class="text-sm mr-1" />
              复制
            </ElButton>
          </div>
          <div class="code-block">
            <pre><code>{{ file.content }}</code></pre>
          </div>
        </ElTabPane>
      </ElTabs>
    </div>

    <!-- 空状态 -->
    <div v-else class="preview-empty">
      <ArtSvgIcon icon="ri:code-s-slash-line" class="text-5xl" style="color: var(--el-text-color-placeholder)" />
      <p>点击"预览代码"按钮查看将要生成的文件</p>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { fetchGenCodesPreview, fetchGenCodesBuild, fetchGenCodesEdit, fetchGenCodesPublishFrontend } from '@/api/backend/develop/genCodes'
  import { ElNotification } from 'element-plus'
  import ArtSvgIcon from '@/components/core/base/art-svg-icon/index.vue'
  import { useClipboard } from '@vueuse/core'

  const props = defineProps<{ formData: any }>()
  const emit = defineEmits<{ (e: 'generated'): void }>()

  const previewLoading = ref(false)
  const buildLoading = ref(false)
  const previewFiles = ref<any[]>([])
  const activeTab = ref('')

  /** 暴露给父组件的刷新方法 */
  const refresh = () => handlePreview()
  defineExpose({ refresh })

  const { copy } = useClipboard()

  const getFileName = (path: string) => path.split('/').pop() || path

  const copyCode = async (content: string) => {
    try {
      await copy(content)
      ElMessage.success('已复制到剪贴板')
    } catch {
      ElMessage.error('复制失败')
    }
  }

  const handlePreview = async () => {
    previewLoading.value = true
    try {
      const res = await fetchGenCodesPreview(props.formData)
      previewFiles.value = res.files || []
      if (previewFiles.value.length) {
        activeTab.value = previewFiles.value[0].path
      }
    } catch (e) {
      console.error('预览失败:', e)
    } finally {
      previewLoading.value = false
    }
  }

  const buildStatusText = ref('')

  // 轮询后端是否就绪（两阶段：先等旧进程断开，再等新进程启动）
  const waitForBackend = async (): Promise<boolean> => {
    const probe = async (timeout = 5000) => {
      try {
        const res = await fetch('/site/index', { method: 'GET', signal: AbortSignal.timeout(timeout) })
        if (!res.ok) return false
        const data = await res.json().catch(() => null)
        return data && data.code === 0
      } catch { return false }
    }

    // 阶段1：等旧进程断开（短超时快速探测，最多 10 秒）
    buildStatusText.value = '等待后端编译中...'
    for (let i = 0; i < 10; i++) {
      await new Promise(r => setTimeout(r, 1000))
      buildStatusText.value = `等待后端编译中... ${i + 1}s`
      if (!(await probe(1500))) break // 短超时，快速检测断开
    }

    // 阶段2：等新进程启动（长超时耐心等响应，最多 40 秒）
    for (let i = 0; i < 20; i++) {
      buildStatusText.value = `等待后端启动... ${20 - i * 2}s`
      if (await probe(8000)) return true // 长超时，等后端冷启动完成
      await new Promise(r => setTimeout(r, 2000))
    }
    return false
  }

  const handleBuild = async () => {
    buildLoading.value = true
    try {
      // 1. 保存配置
      buildStatusText.value = '保存配置...'
      await fetchGenCodesEdit(props.formData)

      // 2. 生成代码（后端写项目目录，前端写临时目录，不触发 Vite HMR）
      buildStatusText.value = '生成后端代码...'
      await fetchGenCodesBuild(props.formData).catch(() => {})

      // 3. 轮询等后端重启就绪
      buildStatusText.value = '等待后端重启...'
      const ready = await waitForBackend()

      // 4. 后端就绪后发布前端文件（一次性写入，触发 Vite HMR，此时后端已正常）
      buildStatusText.value = '发布前端文件...'
      let published = false
      try {
        const pubRes = await fetchGenCodesPublishFrontend()
        published = !!pubRes
      } catch { /* ignore */ }

      if (ready) {
        ElNotification({ title: '生成成功', message: '代码已生成，页面即将刷新...', type: 'success', duration: 2000 })
      } else {
        ElNotification({ title: '生成完成', message: '页面即将刷新...', type: 'warning', duration: 2000 })
      }

      buildStatusText.value = ''
      emit('generated')

      // 5. 等 Vite 处理文件变更后强制刷新（解决二次生成无 HMR 和代理中断后页面 500 的问题）
      await new Promise(r => setTimeout(r, 1500))
      window.location.reload()
    } catch (e) {
      console.error('生成失败:', e)
      ElMessage.error('生成失败')
      buildStatusText.value = ''
    } finally {
      buildLoading.value = false
    }
  }
</script> 

<style scoped>
  @reference '@styles/core/tailwind.css';

  .preview-actions {
    display: flex;
    gap: 12px;
    margin-bottom: 16px;
  }

  .preview-tabs {
    border-radius: calc(var(--custom-radius) / 2 + 2px);
    overflow: hidden;
  }

  .file-path-bar {
    display: flex;
    align-items: center;
    padding: 8px 14px;
    font-size: 12px;
    color: var(--el-text-color-secondary);
    background: var(--el-fill-color-lighter);
    border-radius: calc(var(--custom-radius) / 2) calc(var(--custom-radius) / 2) 0 0;
    margin-bottom: 0;
  }

  .code-block {
    background: #1e1e2e;
    border-radius: 0 0 calc(var(--custom-radius) / 2) calc(var(--custom-radius) / 2);
    overflow: auto;
    max-height: 460px;
  }

  .code-block pre {
    margin: 0;
    padding: 16px 20px;
    font-size: 13px;
    line-height: 1.6;
    color: #cdd6f4;
    font-family: 'JetBrains Mono', 'Fira Code', 'Menlo', 'Monaco', 'Consolas', monospace;
    white-space: pre;
    tab-size: 4;
  }

  .preview-empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 80px 0;
    color: var(--el-text-color-placeholder);
  }

  .preview-empty p {
    margin-top: 16px;
    font-size: 14px;
  }
</style>
