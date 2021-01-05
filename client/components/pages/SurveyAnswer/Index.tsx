import React, { useEffect, useState } from 'react'
import NextError from 'next/error'

import { useRouter } from 'next/router'
import { fetchSurvey } from 'modules/survey/api'
import { Survey } from 'modules/survey/types'
import DefaultTemplate from 'components/templates/DefaultTemplate'
import InitialLoading from 'components/atoms/InitialLoading'
import ContentWrapper from 'components/atoms/ContentWrapper'
import { SurveyAnswerContextProvider } from 'components/pages/SurveyAnswer/Context'
import Outline from 'components/pages/SurveyAnswer/Outline'

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
    </SurveyAnswerContextProvider>
  )
}

export default SurveyAnswer
