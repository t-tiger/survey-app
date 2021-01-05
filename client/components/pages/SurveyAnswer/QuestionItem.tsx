import React, { useContext } from 'react'

import {
  Box,
  FormControlLabel,
  Paper,
  Radio,
  RadioGroup,
  Typography,
} from '@material-ui/core'

import { Question } from 'modules/survey/types'

import SurveyAnswerContext from 'components/pages/SurveyAnswer/Context'

type Props = {
  question: Question
}

const QuestionItem: React.FC<Props> = ({ question }) => {
  const { answers, setAnswers } = useContext(SurveyAnswerContext)
  const selected = answers[question.id]

  const handleSelect = (e: React.ChangeEvent<HTMLInputElement>) => {
    setAnswers({ ...answers, [question.id]: e.target.value })
  }

  return (
    <Paper>
      <Box padding={3}>
        <Typography variant="h5" style={{ lineHeight: 1.4 }}>
          Q{question.sequence}. {question.title}
        </Typography>
        <Box mt={1.5}>
          <RadioGroup value={selected || ''} onChange={handleSelect}>
            {question.options.map((o) => (
              <FormControlLabel
                key={o.id}
                value={o.id}
                control={<Radio color="primary" />}
                label={o.title}
              />
            ))}
          </RadioGroup>
        </Box>
      </Box>
    </Paper>
  )
}

export default QuestionItem
