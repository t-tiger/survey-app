import React, { ReactNode, useCallback, useContext, useRef } from 'react'

import { SnackbarProvider, useSnackbar } from 'notistack'
import { IconButton } from '@material-ui/core'
import { Close } from '@material-ui/icons'

const MAX_SNACK = 5
const AUTO_HIDE_DURATION = 5000

type MessageVariant = 'success' | 'error'
type MessageCenter = {
  showMessage: (variant: MessageVariant, message: string) => void
}

// Context
const Context = React.createContext<MessageCenter>(null!)
export const useMessageCenter = (): MessageCenter => useContext(Context)

// Provider
type Props = { children: ReactNode }
export const MessageCenterProvider: React.FC<Props> = ({ children }) => {
  return (
    <SnackbarProvider
      maxSnack={MAX_SNACK}
      anchorOrigin={{
        vertical: 'bottom',
        horizontal: 'right',
      }}
    >
      <Content>{children}</Content>
    </SnackbarProvider>
  )
}

const Content: React.FC<Props> = ({ children }) => {
  const { enqueueSnackbar, closeSnackbar } = useSnackbar()
  const { current: state } = useRef({ key: 0 })

  const showMessage = useCallback(
    (variant: MessageVariant, message: string) => {
      const key = state.key + 1
      state.key = key

      enqueueSnackbar(message, {
        variant,
        key,
        autoHideDuration: AUTO_HIDE_DURATION,
        action: (
          <IconButton
            color="inherit"
            size="small"
            onClick={() => closeSnackbar(key)}
          >
            <Close />
          </IconButton>
        ),
        style: { maxWidth: '650px' },
      })
    },
    [enqueueSnackbar, closeSnackbar],
  )
  return <Context.Provider value={{ showMessage }}>{children}</Context.Provider>
}
