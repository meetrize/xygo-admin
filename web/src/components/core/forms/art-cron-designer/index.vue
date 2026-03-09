<!-- +----------------------------------------------------------------------
  | XYGo Admin [ Vue3 + GoFrame 企业级中后台管理系统 ]
  +----------------------------------------------------------------------
  | Copyright (c) 2026 大连星韵网络科技有限公司 All rights reserved.
  +----------------------------------------------------------------------
  | Licensed ( https://opensource.org/licenses/MIT )
  +----------------------------------------------------------------------
  | Author: 喜羊羊 <751300685@qq.com>
  +---------------------------------------------------------------------- -->
<!-- Cron 表达式可视化设计器 -->
<template>
  <div class="cron-designer rounded-lg border border-g-200 bg-g-50/50 p-3">
    <!-- 快捷选择：紧凑横排 -->
    <div class="mb-3 flex flex-wrap gap-1.5">
      <button
        v-for="preset in presets"
        :key="preset.value"
        type="button"
        :class="[
          'rounded border px-2.5 py-1 text-xs leading-none transition-all',
          modelValue === preset.value
            ? 'border-theme bg-theme/10 font-medium text-theme'
            : 'border-g-200 bg-white text-g-600 hover:border-theme/40 hover:text-theme'
        ]"
        @click="selectPreset(preset.value)"
      >{{ preset.label }}</button>
    </div>

    <!-- 自定义：6列一排 -->
    <div class="mb-3 flex gap-2">
      <div v-for="f in fieldDefs" :key="f.key" class="min-w-0 flex-1">
        <div class="mb-0.5 text-center text-[10px] text-g-400">{{ f.label }}</div>
        <ElSelect v-model="fields[f.key]" size="small" @change="buildExpression">
          <ElOption v-for="opt in f.options" :key="opt.value" :label="opt.label" :value="opt.value" />
        </ElSelect>
      </div>
    </div>

    <!-- 结果预览 -->
    <div class="flex items-center gap-2 rounded bg-white px-2.5 py-1.5 ring-1 ring-g-200/60">
      <code class="min-w-0 flex-1 truncate text-xs font-semibold text-theme">{{ modelValue || '未设置' }}</code>
      <span class="shrink-0 text-[10px] text-g-400">{{ humanReadable }}</span>
      <button
        type="button"
        class="shrink-0 text-[10px] text-g-500 hover:text-theme"
        @click="showManual = !showManual"
      >{{ showManual ? '收起' : '手动' }}</button>
    </div>
    <ElInput
      v-if="showManual"
      :model-value="modelValue"
      placeholder="秒 分 时 日 月 周"
      size="small"
      class="mt-1.5"
      @input="(val: string) => emit('update:modelValue', val)"
    />
  </div>
</template>

<script setup lang="ts">
  const props = defineProps<{ modelValue: string }>()
  const emit = defineEmits<{ (e: 'update:modelValue', v: string): void }>()
  const showManual = ref(false)

  const presets = [
    { label: '5秒', value: '*/5 * * * * *' },
    { label: '30秒', value: '*/30 * * * * *' },
    { label: '1分钟', value: '0 * * * * *' },
    { label: '5分钟', value: '0 */5 * * * *' },
    { label: '10分钟', value: '0 */10 * * * *' },
    { label: '30分钟', value: '0 */30 * * * *' },
    { label: '每小时', value: '0 0 * * * *' },
    { label: '每天0点', value: '0 0 0 * * *' },
    { label: '每天6点', value: '0 0 6 * * *' },
    { label: '每天12点', value: '0 0 12 * * *' },
    { label: '周一0点', value: '0 0 0 * * 1' },
    { label: '月1号0点', value: '0 0 0 1 * *' },
  ]

  const selectPreset = (v: string) => { emit('update:modelValue', v); parseExpr(v) }

  const fields = reactive<Record<string, string>>({ second: '*', minute: '*', hour: '*', day: '*', month: '*', week: '*' })

  const mkOpts = (max: number, u: string, all: string) => {
    const o = [{ label: all, value: '*' }]
    for (const n of [2,3,5,10,15,20,30].filter(v => v <= max)) o.push({ label: `每${n}${u}`, value: `*/${n}` })
    for (let i = 0; i <= max; i++) o.push({ label: `${i}${u}`, value: String(i) })
    return o
  }

  const fieldDefs = computed(() => [
    { key: 'second', label: '秒', options: mkOpts(59, '秒', '每秒') },
    { key: 'minute', label: '分', options: mkOpts(59, '分', '每分') },
    { key: 'hour',   label: '时', options: mkOpts(23, '时', '每时') },
    { key: 'day',    label: '日', options: [{ label: '每天', value: '*' }, ...Array.from({ length: 31 }, (_, i) => ({ label: `${i+1}号`, value: String(i+1) }))] },
    { key: 'month',  label: '月', options: [{ label: '每月', value: '*' }, ...['1月','2月','3月','4月','5月','6月','7月','8月','9月','10月','11月','12月'].map((n, i) => ({ label: n, value: String(i+1) }))] },
    { key: 'week',   label: '周', options: [{ label: '不限', value: '*' }, { label: '工作日', value: '1-5' }, { label: '周末', value: '0,6' }, { label: '周一', value: '1' }, { label: '周二', value: '2' }, { label: '周三', value: '3' }, { label: '周四', value: '4' }, { label: '周五', value: '5' }, { label: '周六', value: '6' }, { label: '周日', value: '0' }] },
  ])

  const buildExpression = () => { emit('update:modelValue', `${fields.second} ${fields.minute} ${fields.hour} ${fields.day} ${fields.month} ${fields.week}`) }

  const parseExpr = (expr: string) => {
    const p = expr.trim().split(/\s+/)
    if (p.length >= 6) { fields.second = p[0]; fields.minute = p[1]; fields.hour = p[2]; fields.day = p[3]; fields.month = p[4]; fields.week = p[5] }
  }

  const humanReadable = computed(() => {
    const e = props.modelValue; if (!e) return ''
    const m = presets.find(p => p.value === e); if (m) return m.label
    const p = e.trim().split(/\s+/); if (p.length < 6) return '6位'
    const [s,mi,h,d,mo,w] = p; const r: string[] = []
    const wm: Record<string,string> = {'0':'周日','1':'周一','2':'周二','3':'周三','4':'周四','5':'周五','6':'周六','1-5':'工作日','0,6':'周末'}
    if (mo !== '*') r.push(`${mo}月`); if (d !== '*') r.push(`${d}号`); if (w !== '*') r.push(wm[w]||`周${w}`)
    if (h !== '*') r.push(h.startsWith('*/')?`每${h.slice(2)}h`:`${h}点`)
    if (mi !== '*') r.push(mi.startsWith('*/')?`每${mi.slice(2)}min`:`${mi}分`)
    if (s !== '*' && s !== '0') r.push(s.startsWith('*/')?`每${s.slice(2)}s`:`${s}秒`)
    return r.join(' ') || '每秒'
  })

  watch(() => props.modelValue, v => { if (v) parseExpr(v) }, { immediate: true })
</script>

<style scoped>
.cron-designer :deep(.el-select) { width: 100%; }
</style>
