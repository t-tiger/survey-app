import React from 'react'

import { Button, DialogActions, DialogContent } from '@material-ui/core'

const SignUpMenu: React.FC = () => {
  return (
    <>
      <DialogContent>sign up</DialogContent>
      <DialogActions>
        <Button color="primary" variant="contained">
          Sign up
        </Button>
      </DialogActions>
    </>
  )
}

export default SignUpMenu
