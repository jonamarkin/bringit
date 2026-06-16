import type { User } from '~/types/bringit'

export const useAuth = () => {
  const api = useApi()
  const user = useState<User | null>('bringit-user', () => null)

  const loadUser = async () => {
    try {
      const res = await api<{ user: User }>('/api/v1/auth/me')
      user.value = res.user
      return res.user
    } catch {
      user.value = null
      return null
    }
  }

  const requestOTP = (email: string, displayName: string) => {
    return api<{ message: string, dev_code?: string }>('/api/v1/auth/otp/request', {
      method: 'POST',
      body: {
        email,
        display_name: displayName,
      },
    })
  }

  const verifyOTP = async (email: string, displayName: string, code: string) => {
    const res = await api<{ user: User, token: string }>('/api/v1/auth/otp/verify', {
      method: 'POST',
      body: {
        email,
        display_name: displayName,
        code,
      },
    })
    user.value = res.user
    return res
  }

  const logout = async () => {
    await api('/api/v1/auth/logout', { method: 'POST' })
    user.value = null
    await navigateTo('/host')
  }

  return {
    user,
    loadUser,
    requestOTP,
    verifyOTP,
    logout,
  }
}
