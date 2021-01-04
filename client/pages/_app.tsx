import React, { useEffect, useState } from 'react'
import { AppProps } from 'next/app'

import { CssBaseline, ThemeProvider } from '@material-ui/core'

import { mainTheme } from 'const/theme'
import { getAuthStatus, setAuthStatus } from 'utils/session'
import { fetchAuthState } from 'modules/user/api'

import { MessageCenterProvider } from 'utils/messageCenter'
import { AppContextProvider } from 'components/pages/AppContext'

const App: React.FC<AppProps> = ({ Component, pageProps }) => {
  useEffect(() => {
    // Remove server-side injecting CSS.
    const jssStyles = document.querySelector('#jss-server-side')
    jssStyles?.parentElement?.removeChild(jssStyles)
  }, [])

  return (
    <ThemeProvider theme={mainTheme}>
      <MessageCenterProvider>
        <CssBaseline />
        <Root Component={Component} pageProps={pageProps} />
      </MessageCenterProvider>
    </ThemeProvider>
  )
}

type RootProps = Pick<AppProps, 'Component' | 'pageProps'>

const Root: React.FC<RootProps> = ({ Component, pageProps }) => {
  const [ready, setReady] = useState(false)
  const [userId, setUserId] = useState<string>()

  useEffect(() => {
    const checkAuth = async () => {
      // check state from sessionStorage
      const authStatus = getAuthStatus()
      if (authStatus) {
        setUserId(authStatus.userId)
        setReady(true)
      }

      // check state from api
      try {
        const {
          data: { user },
        } = await fetchAuthState()
        setAuthStatus({ checked: true, userId: user?.id })
      } finally {
        setReady(true)
      }
    }
    checkAuth()
  }, [])

  if (!ready) {
    return null
  }
  return (
    <AppContextProvider userId={userId}>
      <Component {...pageProps} />
    </AppContextProvider>
  )
}

export default App
