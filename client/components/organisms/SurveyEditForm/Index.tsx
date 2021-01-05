import React, { useState } from 'react'

import { Box, Button } from '@material-ui/core'

import { Question, Survey } from 'modules/survey/types'
import { issueId } from 'utils/id'

import SurveyEdit from 'components/organisms/SurveyEditForm/SurveyEdit'
import QuestionEdit from 'components/organisms/SurveyEditForm/QuestionEdit'

type Props = {
  survey: {
    title: Survey['title']
    questions: Array<{
      id: string
      title: Question['title']
      options: Array<{ id: string; title: string }>
    }>
  }
}

const SurveyEditForm: React.FC<Props> = ({ survey }) => {
  const [title, setTitle] = useState(survey.title || '')
  const [questions, setQuestions] = useState<Props['survey']['questions']>([
    ...survey.questions,
  ])

  const handleAddQuestion = () => {
    setQuestions([
      ...questions,
      {
        id: issueId(),
        title: '',
        options: [{ id: issueId(), title: '' }],
      },
    ])
  }
  const handleClickRemove = (i: number) => {
    setQuestions(questions.filter((_, j) => i !== j))
  }

  return (
    <>
      <SurveyEdit title={title} onChange={setTitle} />
      {questions.map((q, i) => (
        <Box mt={3} key={q.id}>
          <QuestionEdit
            question={q}
            onClickRemove={() => handleClickRemove(i)}
          />
        </Box>
      ))}
      <Box mt={3} display="flex" justifyContent="space-between">
        <Button
          color="secondary"
          size="large"
          variant="contained"
          onClick={handleAddQuestion}
        >
          Add question
        </Button>
        <Button color="primary" size="large" variant="contained">
          Create new
        </Button>
      </Box>
    </>
  )
}

export default SurveyEditForm
