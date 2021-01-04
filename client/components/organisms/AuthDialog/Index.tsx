import React, { useState } from 'react'

import { Dialog, Tab, Tabs } from '@material-ui/core'

import SignUpMenu from 'components/organisms/AuthDialog/SignUpMenu'
import SignInMenu from 'components/organisms/AuthDialog/SignInMenu'

type Props = {
  open: boolean
  onClose: () => void
}

const AuthDialog: React.FC<Props> = ({ open, onClose }) => {
  const [menu, setMenu] = useState(0)

  const handleChangeMenu = (_: React.ChangeEvent<{}>, value: number) => {
    setMenu(value)
  }

  return (
    <Dialog onClose={onClose} open={open} maxWidth="sm" fullWidth>
      <Tabs
        value={menu}
        onChange={handleChangeMenu}
        variant="fullWidth"
        indicatorColor="primary"
        textColor="primary"
      >
        <Tab label="Sign Up" />
        <Tab label="Sign In" />
      </Tabs>
      {menu === 0 ? (
        <SignUpMenu onFinish={onClose} />
      ) : (
        <SignInMenu onFinish={onClose} />
      )}
    </Dialog>
  )
}

export default AuthDialog
