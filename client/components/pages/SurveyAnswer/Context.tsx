import React, { ReactNode, useState } from 'react'

import { Option, Survey } from 'modules/survey/types'
import { getRespondentUser } from 'modules/survey/helpers'

export type SurveyAnswerState = {
  survey: Survey
  answers: { [questionId: string]: Option['id'] }
  setAnswers: (val: SurveyAnswerState['answers']) => void
  respondent: { name: string; email: string }
  setRespondent: (val: SurveyAnswerState['respondent']) => void
}

const SurveyAnswerContext = React.createContext<SurveyAnswerState>(null!)

type Props = {
  survey: Survey
  children: ReactNode
}

export const SurveyAnswerContextProvider: React.FC<Props> = ({
  survey,
  children,
}) => {
  const [answers, setAnswers] = useState<SurveyAnswerState['answers']>({})
  const [respondent, setRespondent] = useState(
    getRespondentUser() || { name: '', email: '' },
  )
  return (
    <SurveyAnswerContext.Provider
      value={{ survey, answers, setAnswers, respondent, setRespondent }}
    >
      {children}
    </SurveyAnswerContext.Provider>
  )
}

export default SurveyAnswerContext
