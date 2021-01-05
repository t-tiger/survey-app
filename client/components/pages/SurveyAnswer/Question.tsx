import React from 'react'
import { Question } from 'modules/survey/types'

type Props = {
  question: Question
}

const Question: React.FC<Props> = ({ question }) => {
  return <>{question.title}</>
}

export default Question
