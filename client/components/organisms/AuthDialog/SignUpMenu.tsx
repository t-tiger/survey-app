import React, { useContext, useState } from 'react'

import {
  Button,
  DialogActions,
  DialogContent,
  TextField,
} from '@material-ui/core'

import { signUp } from 'modules/user/api'
import { useMessageCenter } from 'utils/messageCenter'

import AppContext from 'components/pages/AppContext'

type Props = {
  onFinish: () => void
}

const SignUpMenu: React.FC<Props> = ({ onFinish }) => {
  const [email, setEmail] = useState('')
  const [name, setName] = useState('')
  const [password, setPassword] = useState('')

  const { setUserId } = useContext(AppContext)
  const { showMessage } = useMessageCenter()

  const handleChangeEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value)
  }
  const handleChangeName = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value)
  }
  const handleChangePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value)
  }
  const handleClickSubmit = async () => {
    try {
      const {
        data: { user },
      } = await signUp({ email, name, password })
      setUserId(user.id)
      showMessage('success', 'Signed up successfully.')
      onFinish()
    } catch (e) {
      if (e.response?.data?.message) {
        showMessage('error', e.response.data.message)
      }
    }
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
          onChange={handleChangeEmail}
          InputLabelProps={{ shrink: true }}
          fullWidth
          required
        />
        <TextField
          margin="normal"
          value={name}
          label="User Name"
          onChange={handleChangeName}
          InputLabelProps={{ shrink: true }}
          fullWidth
          required
        />
        <TextField
          margin="normal"
          type="password"
          value={password}
          label="Password"
          onChange={handleChangePassword}
          InputLabelProps={{ shrink: true }}
          fullWidth
          required
        />
      </DialogContent>
      <DialogActions>
        <Button variant="text" onClick={onFinish}>
          Close
        </Button>
        <Button
          color="primary"
          variant="contained"
          disabled={!readyToSave}
          onClick={handleClickSubmit}
        >
          Sign up
        </Button>
      </DialogActions>
    </>
  )
}

export default SignUpMenu
