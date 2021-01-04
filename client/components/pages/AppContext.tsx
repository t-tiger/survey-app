import React, { ReactNode } from 'react'

export type AppContextState = {
  authorized: boolean
}

const AppContext = React.createContext<AppContextState>(null!)

type Props = {
  authorized: boolean
  children: ReactNode
}

export const AppContextProvider: React.FC<Props> = ({
  authorized,
  children,
}) => {
  return (
    <AppContext.Provider value={{ authorized }}>{children}</AppContext.Provider>
  )
}

export default AppContext
