import React, { useState } from 'react'

import { Box, Button, Divider, Paper } from '@material-ui/core'

import { issueId } from 'utils/id'
import { Question } from 'modules/survey/types'

type Props = {
  question: {
    title: Question['title']
    options: Array<{ id: string; title: string }>
  }
  onClickRemove: () => void
}

const QuestionEdit: React.FC<Props> = ({ question, onClickRemove }) => {
  const [options, setOptions] = useState([...question.options])

  const handleAddOption = () =>
    setOptions([...options, { id: issueId(), title: '' }])

  return (
    <Paper>
      <Box padding={3}>
        {options.map((o) => (
          <Box key={o.id} mb={2}>
            {o.title}
          </Box>
        ))}
        <Divider />
        <Box mt={2} display="flex" justifyContent="space-between">
          <Button
            color="secondary"
            variant="contained"
            onClick={handleAddOption}
          >
            Add option
          </Button>
          <Button variant="text" onClick={onClickRemove}>
            Remove
          </Button>
        </Box>
      </Box>
    </Paper>
  )
}

export default QuestionEdit
