import React from 'react'
import { Box, Paper, TextField, Typography } from '@material-ui/core'

type Props = {
  pageTitle: string
  surveyTitle: string
  onChange: (surveyTitle: string) => void
}

const SurveyEdit: React.FC<Props> = ({ pageTitle, surveyTitle, onChange }) => {
  return (
    <Paper>
      <Box padding={3}>
        <Typography variant="h5">{pageTitle}</Typography>
        <TextField
          margin="normal"
          value={surveyTitle}
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
