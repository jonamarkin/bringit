export const useToast = () => {
  const message = useState<string>('bringit-toast', () => '')
  const timer = useState<ReturnType<typeof setTimeout> | null>('bringit-toast-timer', () => null)

  const showToast = (value: string) => {
    message.value = value
    if (timer.value) {
      clearTimeout(timer.value)
    }
    timer.value = setTimeout(() => {
      message.value = ''
      timer.value = null
    }, 3200)
  }

  return {
    message,
    showToast,
  }
}
