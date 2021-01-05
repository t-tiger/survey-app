import React from 'react'
import { Box, DialogContent, Paper, TextField, Typography } from '@material-ui/core'

type Props = {
  title: string
  onChange: (title: string) => void
}

const SurveyEdit: React.FC<Props> = ({ title, onChange }) => {
  return (
    <Paper>
      <Box padding={3}>
        <Typography variant="h5">
          Post new survey
        </Typography>
        <TextField
          margin="normal"
          value={title}
          label="Input survey title"
          onChange={(e) => onChange(e.target.value)}
          InputLabelProps={{ shrink: true }}
          fullWidth
          required
        />
      </Box>
    </Paper>
  )
}

export default SurveyEdit
