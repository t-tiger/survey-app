import React, { ReactNode, useState } from 'react'

import { Option, Survey } from 'modules/survey/types'

type State = {
  survey: Survey
  answers: { [questionId: string]: Option['id'] }
  setAnswers: (val: State['answers']) => void
  respondent: { name: string; email: string }
  setRespondent: (val: State['respondent']) => void
}

const SurveyAnswerContext = React.createContext<State>(null!)

type Props = {
  survey: Survey
  children: ReactNode
}

export const SurveyAnswerContextProvider: React.FC<Props> = ({
  survey,
  children,
}) => {
  const [answers, setAnswers] = useState<State['answers']>({})
  const [respondent, setRespondent] = useState({ name: '', email: '' })

  return (
    <SurveyAnswerContext.Provider
      value={{ survey, answers, setAnswers, respondent, setRespondent }}
    >
      {children}
    </SurveyAnswerContext.Provider>
  )
}

export default SurveyAnswerContext
