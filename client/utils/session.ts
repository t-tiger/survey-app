const AUTH_SESSION_KEY = 'survey-auth-status'

type AuthStatus = { userId?: string; checked: true }

export const getAuthStatus = (): AuthStatus | undefined => {
  const authStr = sessionStorage.getItem(AUTH_SESSION_KEY)
  if (authStr) {
    return JSON.parse(authStr)
  }
}

export const setAuthStatus = (status: AuthStatus) => {
  sessionStorage.setItem(AUTH_SESSION_KEY, JSON.stringify(status))
}

export const deleteAuthSession = () =>
  sessionStorage.removeItem(AUTH_SESSION_KEY)
