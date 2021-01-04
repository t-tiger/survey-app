/* eslint-disable react/jsx-props-no-spreading */
import React, { useEffect, useState } from 'react'
import { AppProps } from 'next/app'

import { CssBaseline, ThemeProvider } from '@material-ui/core'

import { mainTheme } from 'const/theme'
import { fetchAuthState } from 'modules/auth/api'

import { MessageCenterProvider, useMessageCenter } from 'utils/messageCenter'
import { AppContextProvider } from 'components/pages/AppContext'
import InitialLoading from 'components/atoms/InitialLoading'
import DefaultTemplate from 'components/templates/DefaultTemplate'

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
  const [authorized, setAuthorized] = useState<boolean>()
  const { showMessage } = useMessageCenter()

  useEffect(() => {
    // check authorized status
    const checkAuth = async () => {
      try {
        const {
          data: { authorized },
        } = await fetchAuthState()
        setAuthorized(authorized)
      } catch (e) {
        console.error(e)
        showMessage('error', 'failed to fetch authorization status')
      }
    }
    checkAuth()
  }, [])

  if (authorized === undefined) {
    return (
      <DefaultTemplate title="Surveys" {...pageProps}>
        <InitialLoading />
      </DefaultTemplate>
    )
  }

  return (
    <AppContextProvider authorized={authorized}>
      <Component {...pageProps} />
    </AppContextProvider>
  )
}

export default App
