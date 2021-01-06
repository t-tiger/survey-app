import React, { useContext } from 'react'

import { Box, Paper, TextField, Typography } from '@material-ui/core'

import SurveyAnswerContext from 'components/pages/SurveyAnswer/Context'

const Outline: React.FC = () => {
  const {
    survey,
    respondent,
    respondent: { email, name },
    setRespondent,
  } = useContext(SurveyAnswerContext)

  const handleChangeEmail = (e: React.ChangeEvent<HTMLInputElement>) => {
    setRespondent({ ...respondent, email: e.target.value })
  }
  const handleChangeName = (e: React.ChangeEvent<HTMLInputElement>) => {
    setRespondent({ ...respondent, name: e.target.value })
  }

  return (
    <Paper>
      <Box padding={3}>
        <Typography variant="h5" style={{ lineHeight: 1.4 }}>
          {survey.title}
        </Typography>
        <Box mt={2}>
          <Typography variant="subtitle2">
            Please input your email and name.
          </Typography>
        </Box>
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
          label="Name"
          onChange={handleChangeName}
          InputLabelProps={{ shrink: true }}
          fullWidth
          required
        />
      </Box>
    </Paper>
  )
}

export default Outline
