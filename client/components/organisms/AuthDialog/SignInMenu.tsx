import React from 'react'

import { Button, DialogActions, DialogContent } from '@material-ui/core'

const SignInMenu: React.FC = () => {
  return (
    <>
      <DialogContent>sign in</DialogContent>
      <DialogActions>
        <Button color="primary" variant="contained">
          Sign in
        </Button>
      </DialogActions>
    </>
  )
}

export default SignInMenu
