import React, { useEffect, useState } from 'react'
import NextError from 'next/error'
import { useRouter } from 'next/router'

import { Box } from '@material-ui/core'

import { fetchSurvey } from 'modules/survey/api'
import { Survey } from 'modules/survey/types'
import { SurveyAnswerContextProvider } from 'components/pages/SurveyAnswer/Context'

import DefaultTemplate from 'components/templates/DefaultTemplate'
import InitialLoading from 'components/atoms/InitialLoading'
import ContentWrapper from 'components/atoms/ContentWrapper'
import Outline from 'components/pages/SurveyAnswer/Outline'
import QuestionItem from 'components/pages/SurveyAnswer/QuestionItem'
import SubmitButton from 'components/pages/SurveyAnswer/SubmitButton'

const SurveyAnswer: React.FC = () => {
  const id = useRouter().query.id as string

  const [survey, setSurvey] = useState<Survey>()
  const [ready, setReady] = useState(false)

  const fetch = async () => {
    try {
      const { data } = await fetchSurvey(id)
      setSurvey(data)
      setReady(true)
    } catch (e) {
      setReady(true)
    }
  }

  useEffect(() => {
    fetch()
  }, [])

  if (ready && !survey) {
    return <NextError statusCode={404} />
  }
  return (
    <DefaultTemplate title={survey ? `${survey.title}` : 'Loading'}>
      <ContentWrapper>
        {!survey ? <InitialLoading /> : <Content survey={survey} />}
      </ContentWrapper>
    </DefaultTemplate>
  )
}

type ContentProps = {
  survey: Survey
}

const Content: React.FC<ContentProps> = ({ survey }) => {
  return (
    <SurveyAnswerContextProvider survey={survey}>
      <Outline />
      {survey.questions.map((q) => (
        <Box key={q.id} mt={3.5}>
          <QuestionItem question={q} />
        </Box>
      ))}
      <SubmitButton />
    </SurveyAnswerContextProvider>
  )
}

export default SurveyAnswer
