import React, { ReactNode, useCallback, useState } from 'react'

export type AppContextState = {
  userId?: string
  setUserId: (id: string) => void
  clearUserId: () => void
}

const AppContext = React.createContext<AppContextState>(null!)

type Props = {
  userId?: string
  children: ReactNode
}

export const AppContextProvider: React.FC<Props> = ({
  userId: propUserId,
  children,
}) => {
  const [userId, setUserId] = useState(propUserId)
  const clearUserId = useCallback(() => {
    setUserId(undefined)
  }, [])

  return (
    <AppContext.Provider value={{ userId, setUserId, clearUserId }}>
      {children}
    </AppContext.Provider>
  )
}

export default AppContext
