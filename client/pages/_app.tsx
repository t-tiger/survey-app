import React, { ReactElement } from 'react'
import { AppProps } from 'next/app'

import { CssBaseline, ThemeProvider } from '@material-ui/core'

import { mainTheme } from 'const/theme'
import { MessageCenterProvider } from 'utils/messageCenter'

const App = ({ Component, pageProps }: AppProps): ReactElement => {
  React.useEffect(() => {
    // Remove server-side injecting CSS.
    const jssStyles = document.querySelector('#jss-server-side')
    jssStyles?.parentElement?.removeChild(jssStyles)
  }, [])

  return (
    <ThemeProvider theme={mainTheme}>
      <MessageCenterProvider>
        <CssBaseline />
        <Component {...pageProps} />
      </MessageCenterProvider>
    </ThemeProvider>
  )
}

export default App
