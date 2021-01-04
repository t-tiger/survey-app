import React, { ReactNode } from 'react'

import { AppBar, Toolbar, Typography } from '@material-ui/core'

type Props = {
  title: ReactNode
}

const Header: React.FC<Props> = ({ title }) => {
  return (
    <AppBar position="sticky" color="primary">
      <Toolbar>
        <Typography variant="h6" noWrap>
          {title}
        </Typography>
      </Toolbar>
    </AppBar>
  )
}

export default Header
