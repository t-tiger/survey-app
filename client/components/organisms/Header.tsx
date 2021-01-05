import React, { ReactNode, useContext } from 'react'
import styled from 'styled-components'

import { AppBar, Button, Toolbar, Typography } from '@material-ui/core'

import { CONTENT_MAX_WIDTH } from 'const/size'
import { logout } from 'modules/user/api'
import { useMessageCenter } from 'utils/messageCenter'
import { useToggleDialog } from 'utils/dialog'

import AppContext from 'components/pages/AppContext'
import AuthDialog from 'components/organisms/AuthDialog/Index'

type Props = {
  title: ReactNode
}

const Header: React.FC<Props> = ({ title }) => {
  const { userId, clearUserId } = useContext(AppContext)
  const { showMessage } = useMessageCenter()
  const [authDialogKey, isOpenAuthDialog, setOpenAuthDialog] = useToggleDialog()

  const handleLogOutClick = async () => {
    try {
      await logout()
      clearUserId()
      showMessage('success', 'Logged out successfully.')
    } catch (e) {
      if (e.response?.data?.message) {
        showMessage('error', e.response.data.message)
      }
    }
  }
  const handleSignUpClick = () => {
    setOpenAuthDialog(true)
  }

  return (
    <>
      <AppBar position="sticky" color="primary">
        <Toolbar>
          <ToolbarInner>
            <Typography variant="h5" style={{ fontWeight: 'bold' }} noWrap>
              {title}
            </Typography>
            <MenuContainer>
              {userId ? (
                <Button
                  color="inherit"
                  variant="outlined"
                  onClick={handleLogOutClick}
                >
                  Log out
                </Button>
              ) : (
                <Button
                  color="inherit"
                  variant="outlined"
                  onClick={handleSignUpClick}
                >
                  Sign Up
                </Button>
              )}
            </MenuContainer>
          </ToolbarInner>
        </Toolbar>
      </AppBar>
      <AuthDialog
        key={authDialogKey}
        open={isOpenAuthDialog}
        onClose={() => setOpenAuthDialog(false)}
      />
    </>
  )
}

const ToolbarInner = styled.div`
  display: flex;
  flex-grow: 1;
  max-width: ${CONTENT_MAX_WIDTH}px;
  margin: auto;
`

const MenuContainer = styled.div`
  flex-grow: 1;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  margin-left: 15px;
`

export default Header
