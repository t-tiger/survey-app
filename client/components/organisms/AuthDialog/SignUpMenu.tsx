import React, { useState } from 'react'

import {
  Button,
  DialogActions,
  DialogContent,
  TextField,
} from '@material-ui/core'

const SignUpMenu: React.FC = () => {
  const [email, setEmail] = useState('')
  const [name, setName] = useState('')
  const [password, setPassword] = useState('')

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value)
  }
  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value)
  }
  const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value)
  }

  const readyToSave =
    email.trim().length > 0 && name.trim().length > 0 && password.length > 0

  return (
    <>
      <DialogContent dividers>
        <TextField
          margin="normal"
          value={email}
          label="Email"
          onChange={handleEmailChange}
          InputLabelProps={{ shrink: true }}
          fullWidth
          required
        />
        <TextField
          margin="normal"
          value={name}
          label="User Name"
          onChange={handleNameChange}
          InputLabelProps={{ shrink: true }}
          fullWidth
          required
        />
        <TextField
          margin="normal"
          type="password"
          value={password}
          label="Password"
          onChange={handlePasswordChange}
          InputLabelProps={{ shrink: true }}
          fullWidth
          required
        />
      </DialogContent>
      <DialogActions>
        <Button color="primary" variant="contained" disabled={!readyToSave}>
          Sign up
        </Button>
      </DialogActions>
    </>
  )
}

export default SignUpMenu
