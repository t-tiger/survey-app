import React from 'react'

import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
} from '@material-ui/core'

type Props = {
  open: boolean
  color?: 'primary' | 'secondary'
  title: string
  description: string
  submitText: string
  onClose: () => void
  onSubmit: () => void
}

const ConfirmDialog: React.FC<Props> = ({
  open,
  color,
  title,
  description,
  submitText,
  onClose,
  onSubmit,
}: Props) => {
  return (
    <Dialog onClose={onClose} maxWidth="sm" open={open} fullWidth>
      <DialogTitle>{title}</DialogTitle>
      <DialogContent>{description}</DialogContent>
      <DialogActions>
        <Button onClick={onClose}>Cancel</Button>
        <Button
          onClick={onSubmit}
          color={color || 'primary'}
          variant="contained"
        >
          {submitText}
        </Button>
      </DialogActions>
    </Dialog>
  )
}

export default ConfirmDialog
