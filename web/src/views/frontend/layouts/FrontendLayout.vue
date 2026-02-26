<template>
  <div class="frontend-layout">
    <!-- 动态背景色块 -->
    <div class="fixed inset-0 pointer-events-none -z-10 overflow-hidden">
      <div class="absolute top-[-10%] left-[-10%] w-[50vw] h-[50vw] bg-blue-500/5 rounded-full blur-3xl animate-float"></div>
      <div class="absolute bottom-[-10%] right-[-10%] w-[60vw] h-[60vw] bg-cyan-500/5 rounded-full blur-3xl animate-float-delayed"></div>
      <div class="absolute top-[20%] right-[10%] w-[30vw] h-[30vw] bg-indigo-500/5 rounded-full blur-3xl animate-float"></div>
    </div>

    <!-- 导航栏 -->
    <nav
      class="fixed top-0 left-0 right-0 z-50 bg-white/70 backdrop-blur-xl border-b border-[#d1d9e6]/40 transition-all duration-300"
      :class="{ 'shadow-clay-card py-3': isScrolled, 'py-5': !isScrolled }"
    >
      <div class="max-w-[1440px] mx-auto px-8 flex items-center justify-between">
        <!-- Logo（读取站点配置） -->
        <div class="flex items-center gap-3 cursor-pointer group" @click="router.push('/')">
          <img v-if="siteStore.getLogo()" :src="siteStore.getLogo()" alt="logo" class="w-10 h-10 rounded-full shadow-clay-btn group-hover:scale-105 transition-transform object-cover" />
          <div v-else class="w-10 h-10 rounded-full bg-gradient-to-br from-blue-400 to-blue-600 shadow-clay-btn flex items-center justify-center text-white font-bold text-xl group-hover:scale-105 transition-transform">
            {{ siteName.charAt(0) }}
          </div>
          <span class="font-heading font-extrabold text-2xl text-clay-foreground tracking-tight">{{ siteNameFirst }}<span class="text-clay-accent">{{ siteNameLast }}</span></span>
        </div>

        <!-- Desktop Menu -->
        <div class="hidden md:flex items-center gap-12 w-fit mx-auto">
          <template v-for="item in navItems" :key="item.name">
            <!-- 外链：用 <a> 新标签页打开 -->
            <a
              v-if="item.isExternal"
              :href="item.url"
              target="_blank"
              rel="noopener"
              class="font-bold text-sm transition-colors relative group text-clay-muted hover:text-clay-accent"
            >
              {{ item.name }}
              <span class="absolute -bottom-1 left-0 h-0.5 bg-clay-accent transition-all duration-300 w-0 group-hover:w-full"></span>
            </a>
            <!-- 内部路由：用 RouterLink -->
            <RouterLink
              v-else
              :to="item.url"
              class="font-bold text-sm transition-colors relative group"
              :class="isActiveNav(item.url) ? 'text-clay-accent' : 'text-clay-muted hover:text-clay-accent'"
            >
              {{ item.name }}
              <span
                class="absolute -bottom-1 left-0 h-0.5 bg-clay-accent transition-all duration-300"
                :class="isActiveNav(item.url) ? 'w-full' : 'w-0 group-hover:w-full'"
              ></span>
            </RouterLink>
          </template>
        </div>

        <!-- 右侧操作区 -->
        <div class="hidden md:flex items-center gap-4">
          <!-- 语言切换 -->
          <div class="flex items-center gap-2 px-3 py-2 rounded-full bg-white/50 cursor-pointer hover:bg-white transition-all" @click="toggleLang">
            <ArtSvgIcon icon="ri:translate-2" class="text-base text-clay-muted" />
            <span class="text-xs font-bold text-clay-muted">{{ currentLang }}</span>
          </div>


          <!-- 会员中心开启时才显示登录/用户菜单 -->
          <template v-if="memberCenterOpen">
            <!-- 已登录：用户信息 -->
            <template v-if="isLoggedIn">
              <ElDropdown @command="handleUserCommand">
                <div class="flex items-center gap-2 px-4 py-2 rounded-full bg-white shadow-clay-btn cursor-pointer">
                  <ElAvatar :size="28" :src="memberInfo.avatar" class="border-2 border-white">
                    {{ memberInfo.nickname?.charAt(0) || 'U' }}
                  </ElAvatar>
                  <span class="text-sm font-bold text-clay-foreground">{{ memberInfo.nickname || memberInfo.username }}</span>
                </div>
                <template #dropdown>
                  <ElDropdownMenu>
                    <ElDropdownItem command="user">
                      <ArtSvgIcon icon="ri:user-line" class="text-base mr-2" />
                      用户中心
                    </ElDropdownItem>
                    <ElDropdownItem divided command="logout">
                      <ArtSvgIcon icon="ri:logout-box-r-line" class="text-base mr-2" />
                      退出登录
                    </ElDropdownItem>
                  </ElDropdownMenu>
                </template>
              </ElDropdown>
            </template>

            <!-- 未登录：登录按钮 -->
            <template v-else>
              <RouterLink
                to="/user/login"
                class="px-8 py-2.5 rounded-full bg-gradient-to-br from-blue-400 to-blue-600 text-white font-bold text-sm shadow-clay-btn hover:shadow-clay-btn-hover hover:-translate-y-1 active:scale-95 active:shadow-clay-pressed transition-all duration-300"
              >
                登录 / 注册
              </RouterLink>
            </template>
          </template>
        </div>

        <!-- Mobile Menu Button -->
        <button class="md:hidden text-clay-foreground p-2" @click="mobileMenuOpen = !mobileMenuOpen">
          <ArtSvgIcon icon="ri:menu-line" class="text-2xl" />
        </button>
      </div>

      <!-- Mobile Menu Panel -->
      <Transition name="fade">
        <div
          v-if="mobileMenuOpen"
          class="fixed inset-0 z-[60] bg-white/95 backdrop-blur-2xl p-8 flex flex-col items-center justify-center gap-8 md:hidden"
        >
          <button class="absolute top-8 right-8 text-clay-foreground p-2" @click="mobileMenuOpen = false">
            <ArtSvgIcon icon="ri:close-line" class="text-[28px]" />
          </button>
          <div class="flex flex-col items-center gap-6 w-full">
            <template v-for="item in navItems" :key="item.name">
              <a
                v-if="item.isExternal"
                :href="item.url"
                target="_blank"
                rel="noopener"
                class="text-2xl font-black text-clay-foreground hover:text-clay-accent transition-colors"
                @click="mobileMenuOpen = false"
              >{{ item.name }}</a>
              <RouterLink
                v-else
                :to="item.url"
                class="text-2xl font-black text-clay-foreground hover:text-clay-accent transition-colors"
                @click="mobileMenuOpen = false"
              >{{ item.name }}</RouterLink>
            </template>
          </div>
          <div v-if="memberCenterOpen" class="mt-8 w-full">
            <RouterLink
              to="/user/login"
              class="block w-full py-5 rounded-[24px] bg-gradient-to-br from-blue-400 to-blue-600 text-white text-center font-black text-lg shadow-clay-btn"
              @click="mobileMenuOpen = false"
            >
              登录 / 进入工作台
            </RouterLink>
          </div>
        </div>
      </Transition>
    </nav>

    <!-- 主内容区 -->
    <main class="frontend-main">
      <div v-if="pageLoading" class="flex items-center justify-center min-h-[60vh]">
        <div class="flex flex-col items-center gap-4">
          <div class="w-10 h-10 border-3 border-blue-400 border-t-transparent rounded-full animate-spin"></div>
          <span class="text-sm text-clay-muted font-medium">加载中...</span>
        </div>
      </div>
      <RouterView v-else />
    </main>

    <!-- 底部 -->
    <footer class="relative bg-white/70 backdrop-blur-xl pt-16 pb-10 px-6 mt-8 shadow-[0_-20px_40px_rgba(165,175,190,0.2),inset_0_10px_20px_rgba(255,255,255,0.8)] border-t border-[#d1d9e6]/40">
      <div class="max-w-7xl mx-auto">
        <div class="grid md:grid-cols-4 gap-12 mb-12">
          <div class="col-span-1 md:col-span-2">
            <div class="flex items-center gap-3 mb-6">
              <img v-if="siteStore.getLogo()" :src="siteStore.getLogo()" alt="logo" class="w-10 h-10 rounded-full shadow-clay-btn object-cover" />
              <div v-else class="w-10 h-10 rounded-full bg-gradient-to-br from-blue-400 to-blue-600 shadow-clay-btn flex items-center justify-center text-white font-bold text-xl">{{ siteName.charAt(0) }}</div>
              <span class="font-heading font-extrabold text-2xl text-clay-foreground tracking-tight">{{ siteNameFirst }}<span class="text-clay-accent">{{ siteNameLast }}</span></span>
            </div>
            <p class="text-clay-muted max-w-sm leading-relaxed">
              {{ siteStore.getSiteSubtitle() || '基于 Vue3 + GoFrame 的开源中后台管理框架，开箱即用，快速启动你的业务开发。' }}
            </p>
          </div>
          <div>
            <h4 class="font-heading font-bold text-clay-foreground mb-6">产品</h4>
            <ul class="space-y-3">
              <li><RouterLink to="/docs" class="text-clay-muted hover:text-clay-accent transition-colors">文档中心</RouterLink></li>
              <li><a href="https://www.xygoadmin.com" target="_blank" class="text-clay-muted hover:text-clay-accent transition-colors">在线演示</a></li>
            </ul>
          </div>
          <div>
            <h4 class="font-heading font-bold text-clay-foreground mb-6">资源</h4>
            <ul class="space-y-3">
              <li><a href="https://github.com/z312193608/xygo-admin" target="_blank" class="text-clay-muted hover:text-clay-accent transition-colors">GitHub</a></li>
              <li><a href="https://gitee.com/a751300685a/xygo-admin" target="_blank" class="text-clay-muted hover:text-clay-accent transition-colors">Gitee</a></li>
            </ul>
          </div>
        </div>
        <div class="border-t border-gray-100 pt-8 flex flex-col md:flex-row justify-between items-center gap-6">
          <p class="text-sm text-clay-muted font-medium">&copy; {{ new Date().getFullYear() }} {{ siteName }}. All rights reserved.</p>
          <div class="flex gap-4">
            <a href="https://github.com/z312193608/xygo-admin" target="_blank" class="w-12 h-12 rounded-full bg-white shadow-clay-btn hover:shadow-clay-btn-hover flex items-center justify-center text-clay-foreground transition-all hover:-translate-y-1">
              <ArtSvgIcon icon="ri:github-fill" class="text-[22px]" />
            </a>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { useMemberStore } from '@/store/modules/member'
import { useMemberMenuStore } from '@/store/modules/memberMenu'
import { useSiteStore } from '@/store/modules/site'
import { useSettingStore } from '@/store/modules/setting'
import { useI18n } from 'vue-i18n'
import { ElMessageBox } from 'element-plus'

defineOptions({ name: 'FrontendLayout' })

const router = useRouter()
const route = useRoute()
const { locale } = useI18n()

const memberStore = useMemberStore()
const memberMenuStore = useMemberMenuStore()
const siteStore = useSiteStore()
const settingStore = useSettingStore()

const siteName = computed(() => siteStore.getSiteName())
// 站点名称拆分：前半深色 + 后半蓝色（保持渐变品牌效果）
const siteNameFirst = computed(() => {
  const name = siteName.value
  const spaceIdx = name.indexOf(' ')
  if (spaceIdx > 0) return name.substring(0, spaceIdx)
  return name.substring(0, Math.ceil(name.length / 2))
})
const siteNameLast = computed(() => {
  const name = siteName.value
  const spaceIdx = name.indexOf(' ')
  if (spaceIdx > 0) return name.substring(spaceIdx)
  return name.substring(Math.ceil(name.length / 2))
})
const currentLang = computed(() => locale.value === 'zh-CN' ? '中文' : 'EN')
const isDark = computed(() => settingStore.isDark)
const memberInfo = computed(() => memberStore.getMemberInfo)
const isLoggedIn = computed(() => memberStore.isLogin)
const memberCenterOpen = computed(() => siteStore.isUserCenterEnabled())

// 导航菜单：优先从后端获取（type=nav），降级为硬编码首页
interface NavItem { name: string; url: string; isExternal: boolean }
const defaultNavItems: NavItem[] = [{ name: '首页', url: '/', isExternal: false }]

const navItems = computed<NavItem[]>(() => {
  const dynamicNav = memberMenuStore.getNavMenus
  // 始终保留首页
  const items: NavItem[] = [{ name: '首页', url: '/', isExternal: false }]
  if (dynamicNav.length > 0) {
    dynamicNav.forEach((m) => {
      const isLink = m.menuType === 'link'
      const url = isLink ? m.url : (m.path || `/${m.name}`)
      const isExternal = isLink || url.startsWith('http')
      items.push({ name: m.title, url, isExternal })
    })
  }
  return items
})

const isActiveNav = (url: string) => {
  if (url === '/') return route.path === '/'
  return route.path.startsWith(url)
}

// 页面加载状态
const pageLoading = ref(true)

// Layout 挂载时加载站点信息 + 前台菜单
onMounted(async () => {
  try {
    // 加载站点配置（名称、logo 等）
    if (!siteStore.loaded) {
      await siteStore.loadSiteInfo()
    }
    // 设置页面标题
    document.title = siteStore.getSiteName()
    // 加载前台菜单
    if (!memberMenuStore.isLoaded) {
      await memberMenuStore.fetchMenus()
    }
  } finally {
    pageLoading.value = false
  }
})

// 监听登录状态变化
watch(() => memberStore.isLogin, async (newVal) => {
  if (newVal) {
    await memberMenuStore.fetchMenus()
  } else {
    // 登出后重新拉取公开菜单（no_login_valid=1 的菜单未登录也应显示）
    await memberMenuStore.fetchMenus()
  }
})

// 滚动状态
const isScrolled = ref(false)
const mobileMenuOpen = ref(false)

const handleScroll = () => {
  isScrolled.value = window.scrollY > 20
}

onMounted(() => window.addEventListener('scroll', handleScroll))
onUnmounted(() => window.removeEventListener('scroll', handleScroll))

// 切换语言
const toggleLang = () => {
  const newLang = locale.value === 'zh-CN' ? 'en' : 'zh-CN'
  locale.value = newLang
  localStorage.setItem('language', newLang)
}


// 用户操作
const handleUserCommand = async (command: string) => {
  switch (command) {
    case 'user':
      router.push('/user')
      break
    case 'logout':
      try {
        await ElMessageBox.confirm('确定要退出登录吗？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
        memberStore.logOut({ redirect: false })
        router.push('/')
      } catch {
        // 取消
      }
      break
  }
}
</script>

<style lang="scss" scoped>
/* ===== 黏土拟态全局变量 ===== */
/*
 * 对齐 homesite 的背景渲染方式：
 * homesite 在 body 上设 background-color: #f2f4f7，
 * 动态色块 fixed -z-10 渲染在 body 背景色之上、内容之下。
 *
 * 这里不能在 .frontend-layout 上设实色 background-color，
 * 否则会遮住 -z-10 的动态色块。改用 ::before 伪元素做底色，
 * z-index 设为 -20，低于色块的 -10。
 */
.frontend-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  color: #32325d;
  -webkit-font-smoothing: antialiased;
  position: relative;
}

.frontend-layout::before {
  content: '';
  position: fixed;
  inset: 0;
  z-index: -20;
  background-color: #f2f4f7;
}

/* 暗色模式适配（前后台主题同步） */
:global(html.dark) .frontend-layout::before {
  background-color: #1a1c23;
}
:global(html.dark) .frontend-layout {
  color: #e0e0e0;
}

.frontend-main {
  flex: 1;
  margin-top: 76px; // 导航栏高度（py-5=20px*2 + 内容约36px）
  padding-top: 12px; // 确保内容卡片上边框可见
}

/* ===== 黏土拟态阴影 ===== */
:deep(.shadow-clay-deep) {
  box-shadow:
    30px 30px 60px #d1d9e6,
    -30px -30px 60px #ffffff,
    inset 10px 10px 20px rgba(90, 141, 238, 0.05),
    inset -10px -10px 20px rgba(255, 255, 255, 0.8);
}

.shadow-clay-card {
  box-shadow:
    16px 16px 32px rgba(165, 175, 190, 0.3),
    -10px -10px 24px rgba(255, 255, 255, 0.9),
    inset 6px 6px 12px rgba(90, 141, 238, 0.03),
    inset -6px -6px 12px rgba(255, 255, 255, 1);
}

.shadow-clay-btn {
  box-shadow:
    12px 12px 24px rgba(90, 141, 238, 0.3),
    -8px -8px 16px rgba(255, 255, 255, 0.4),
    inset 4px 4px 8px rgba(255, 255, 255, 0.4),
    inset -4px -4px 8px rgba(0, 0, 0, 0.05);
}

.shadow-clay-btn-hover {
  box-shadow:
    16px 16px 32px rgba(90, 141, 238, 0.4),
    -10px -10px 20px rgba(255, 255, 255, 0.5),
    inset 4px 4px 8px rgba(255, 255, 255, 0.4),
    inset -4px -4px 8px rgba(0, 0, 0, 0.05);
}

.shadow-clay-pressed {
  box-shadow:
    inset 10px 10px 20px #e0e5ec,
    inset -10px -10px 20px #ffffff;
}

/* ===== 黏土色彩 ===== */
.text-clay-foreground { color: #32325d; }
.text-clay-muted { color: #8898aa; }
.text-clay-accent { color: #5a8dee; }
.bg-clay-bg { background-color: #f2f4f7; }

/* ===== 动画 ===== */
@keyframes float {
  0%, 100% { transform: translateY(0); }
  50% { transform: translateY(-20px); }
}

.animate-float { animation: float 8s ease-in-out infinite; }
.animate-float-delayed { animation: float 8s ease-in-out 4s infinite; }

/* ===== 字体 ===== */
.font-heading { font-family: 'Nunito', 'PingFang SC', sans-serif; }

/* ===== 过渡 ===== */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
