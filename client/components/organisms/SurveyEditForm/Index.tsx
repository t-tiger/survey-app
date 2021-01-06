import React, { useEffect, useState } from 'react'

import { Box, Button } from '@material-ui/core'

import { Option, Question, Survey } from 'modules/survey/types'
import { issueId } from 'utils/id'
import { validateSurvey } from 'modules/survey/helpers'

import QuestionEdit from 'components/organisms/SurveyEditForm/QuestionEdit'
import MultiLineToolTip from 'components/atoms/MultiLineTooltip'
import SurveyEdit from 'components/organisms/SurveyEditForm/SurveyEdit'

export type SurveyForEdit = {
  title: Survey['title']
  questions: Array<
    Pick<Question, 'id' | 'title'> & {
      options: Array<Pick<Option, 'id' | 'title'>>
    }
  >
}

type Props = {
  survey: SurveyForEdit
  pageTitle: string
  submitTitle: string
  onSubmit: (survey: SurveyForEdit) => Promise<void>
}

const SurveyEditForm: React.FC<Props> = ({
  survey,
  pageTitle,
  submitTitle,
  onSubmit,
}) => {
  const [sending, setSending] = useState(false)
  const [title, setTitle] = useState(survey.title || '')
  const [questions, setQuestions] = useState<SurveyForEdit['questions']>([
    ...survey.questions,
  ])
  const [validationErrs, setValidationErrs] = useState<string[]>([])

  useEffect(() => {
    setValidationErrs(validateSurvey(title, questions))
  }, [title, questions])

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
  const handleChangeQuestion = (
    idx: number,
    question: typeof questions[number],
  ) => {
    setQuestions(questions.map((q, i) => (idx === i ? question : q)))
  }
  const handleRemoveQuestion = (idx: number) => {
    setQuestions(questions.filter((_, i) => idx !== i))
  }
  const handleSubmit = async () => {
    setSending(true)
    await onSubmit({
      title,
      questions: questions.map((q) => ({
        ...q,
        options: q.options.filter((o) => o.title.trim().length > 0),
      })),
    })
    setSending(false)
  }

  return (
    <>
      <SurveyEdit
        pageTitle={pageTitle}
        surveyTitle={title}
        onChange={setTitle}
      />
      {questions.map((q, i) => (
        <Box mt={3} key={q.id}>
          <QuestionEdit
            question={{ ...q, sequence: i + 1 }}
            onChange={(q) => handleChangeQuestion(i, q)}
            onRemove={() => handleRemoveQuestion(i)}
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
        <MultiLineToolTip titles={validationErrs}>
          <Button
            color="primary"
            size="large"
            variant="contained"
            disabled={sending || validationErrs.length > 0}
            onClick={handleSubmit}
          >
            {submitTitle}
          </Button>
        </MultiLineToolTip>
      </Box>
    </>
  )
}

export default SurveyEditForm
