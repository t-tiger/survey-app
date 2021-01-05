import React from 'react'

import {
  Box,
  Button,
  Divider,
  Paper,
  TextField,
  Typography,
} from '@material-ui/core'

import { issueId } from 'utils/id'
import { Option, Question } from 'modules/survey/types'

type Props = {
  question: Pick<Question, 'id' | 'title' | 'sequence'> & {
    options: Array<Pick<Option, 'id' | 'title'>>
  }
  onChange: (question: Props['question']) => void
  onRemove: () => void
}

const QuestionEdit: React.FC<Props> = ({ question, onChange, onRemove }) => {
  const { title, options } = question

  const handleChangeQuestionTitle = (
    e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>,
  ) => {
    onChange({ ...question, title: e.target.value })
  }
  const handleChangeOptionTitle = (
    i: number,
    e: React.ChangeEvent<HTMLTextAreaElement | HTMLInputElement>,
  ) => {
    onChange({
      ...question,
      options: options.map((o, j) =>
        i === j ? { ...o, title: e.target.value } : o,
      ),
    })
  }
  const handleAddOption = () => {
    onChange({
      ...question,
      options: [...options, { id: issueId(), title: '' }],
    })
  }

  return (
    <Paper>
      <Box padding={3}>
        <Typography variant="h5">QuestionItem {question.sequence}</Typography>
        <TextField
          margin="normal"
          value={title}
          label="QuestionItem title"
          onChange={handleChangeQuestionTitle}
          InputLabelProps={{ shrink: true }}
          fullWidth
          required
        />
        <Box marginY={2}>
          <Divider />
        </Box>
        <Typography variant="h6">Options</Typography>
        {question.options.map((o, i) => (
          <TextField
            key={o.id}
            margin="normal"
            value={o.title}
            label={`Title of option ${i + 1}`}
            onChange={(e) => handleChangeOptionTitle(i, e)}
            InputLabelProps={{ shrink: true }}
            fullWidth
            required
          />
        ))}
        <Box marginY={2}>
          <Divider />
        </Box>
        <Box display="flex" justifyContent="space-between">
          <Button
            color="secondary"
            variant="contained"
            onClick={handleAddOption}
          >
            Add option
          </Button>
          <Button variant="text" onClick={onRemove}>
            Remove
          </Button>
        </Box>
      </Box>
    </Paper>
  )
}

export default QuestionEdit
