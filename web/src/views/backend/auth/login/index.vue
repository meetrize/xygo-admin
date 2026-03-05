<!-- 登录页面 -->
<template>
  <div class="flex w-full h-screen">
    <LoginLeftView />

    <div class="relative flex-1">
      <AuthTopBar />

      <div class="auth-right-wrap">
        <div class="form">
          <h3 class="title">{{ $t('login.title') }}</h3>
          <p class="sub-title">{{ $t('login.subTitle') }}</p>
          <ElForm
            ref="formRef"
            :model="formData"
            :rules="rules"
            :key="formKey"
            @keyup.enter="handleSubmit"
            style="margin-top: 25px"
          >
            <ElFormItem prop="username">
              <ElInput
                class="custom-height"
                :placeholder="$t('login.placeholder.username')"
                v-model.trim="formData.username"
              />
            </ElFormItem>
            <ElFormItem prop="password">
              <ElInput
                class="custom-height"
                :placeholder="$t('login.placeholder.password')"
                v-model.trim="formData.password"
                type="password"
                autocomplete="off"
                show-password
              />
            </ElFormItem>

            <!-- 点选验证码 -->
            <ArtClickCaptcha ref="clickCaptchaRef" :on-success="onCaptchaSuccess" />

            <div class="flex-cb mt-2 text-sm">
              <ElCheckbox v-model="formData.rememberPassword">{{
                $t('login.rememberPwd')
              }}</ElCheckbox>
              <RouterLink class="text-theme" :to="{ name: 'ForgetPassword' }">{{
                $t('login.forgetPwd')
              }}</RouterLink>
            </div>

            <div style="margin-top: 30px">
              <ElButton
                class="w-full custom-height"
                type="primary"
                @click="handleSubmit"
                :loading="loading"
                v-ripple
              >
                {{ $t('login.btnText') }}
              </ElButton>
            </div>

            <div class="mt-5 text-sm text-gray-600">
              <span>{{ $t('login.noAccount') }}</span>
              <RouterLink class="text-theme" :to="{ name: 'Register' }">{{
                $t('login.register')
              }}</RouterLink>
            </div>
          </ElForm>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { useUserStore } from '@/store/modules/user'
  import { useSiteStore } from '@/store/modules/site'
  import { useI18n } from 'vue-i18n'
  import { HttpError } from '@/utils/http/error'
  import { fetchLogin } from '@/api/backend/auth'
  import { ADMIN_BASE_PATH } from '@/router/routesAlias'
  import { ElNotification, type FormInstance, type FormRules } from 'element-plus'
  import ArtClickCaptcha from '@/components/core/forms/art-click-captcha/index.vue'

  defineOptions({ name: 'Login' })

  const { t, locale } = useI18n()
  const formKey = ref(0)
  watch(locale, () => { formKey.value++ })

  const clickCaptchaRef = ref<InstanceType<typeof ArtClickCaptcha>>()
  const userStore = useUserStore()
  const router = useRouter()
  const route = useRoute()
  const formRef = ref<FormInstance>()

  const siteStore = useSiteStore()
  const systemName = computed(() => siteStore.getSiteName())

  const formData = reactive({
    username: 'admin',
    password: '123456',
    rememberPassword: true
  })

  const rules = computed<FormRules>(() => ({
    username: [{ required: true, message: t('login.placeholder.username'), trigger: 'blur' }],
    password: [{ required: true, message: t('login.placeholder.password'), trigger: 'blur' }]
  }))

  const loading = ref(false)

  const handleSubmit = async () => {
    if (!formRef.value) return
    const valid = await formRef.value.validate().catch(() => false)
    if (!valid) return
    clickCaptchaRef.value?.open()
  }

  const onCaptchaSuccess = async (captchaId: string, captchaInfo: string) => {
    loading.value = true
    try {
      const { username, password } = formData
      const { accessToken, token, refreshToken } = await fetchLogin({
        userName: username,
        password,
        captchaId,
        captchaInfo,
      })

      const finalToken = accessToken || token
      if (!finalToken) throw new Error('Login failed - no token received')

      userStore.setToken(finalToken, refreshToken)
      userStore.setLoginStatus(true)

      const redirect = route.query.redirect as string
      let targetPath = `${ADMIN_BASE_PATH}/dashboard/console`
      if (redirect && !redirect.startsWith('/user') && redirect !== '/') {
        targetPath = redirect
      }

      router.push(targetPath)

      const unregister = router.afterEach(() => {
        unregister()
        loading.value = false
        showLoginSuccessNotice()
      })
    } catch (error) {
      loading.value = false
      if (!(error instanceof HttpError)) {
        console.error('[Login] Unexpected error:', error)
      }
    }
  }

  const showLoginSuccessNotice = () => {
    setTimeout(() => {
      const userInfo = userStore.getUserInfo
      const displayName = userInfo.nickname || userInfo.username || systemName.value
      ElNotification({
        title: t('login.success.title'),
        type: 'success',
        duration: 2500,
        zIndex: 10000,
        message: `${t('login.success.message')}, ${displayName}!`
      })
    }, 1000)
  }
</script>

<style scoped>
  @import './style.css';
</style>
