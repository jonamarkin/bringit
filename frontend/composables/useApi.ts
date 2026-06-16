export const useApi = () => {
  const config = useRuntimeConfig()
  const reqHeaders = useRequestHeaders(['cookie'])

  const apiFetch = <T = any>(request: string, opts: any = {}) => {
    return $fetch<T>(request, {
      baseURL: config.public.apiBase as string,
      credentials: 'include',
      ...opts,
      headers: {
        ...(reqHeaders || {}),
        ...(opts.headers || {}),
      },
    })
  }

  return apiFetch
}
