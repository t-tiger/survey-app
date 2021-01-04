import React, { ReactNode } from 'react'

export type AppContextState = {
  userId?: string
}

const AppContext = React.createContext<AppContextState>(null!)

type Props = {
  userId?: string
  children: ReactNode
}

export const AppContextProvider: React.FC<Props> = ({ userId, children }) => (
  <AppContext.Provider value={{ userId }}>{children}</AppContext.Provider>
)

export default AppContext
