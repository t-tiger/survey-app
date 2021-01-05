import React, { useEffect, useState } from 'react'
import NextError from 'next/error'
import Router, { useRouter } from 'next/router'

import { Box } from '@material-ui/core'

import { fetchSurvey } from 'modules/survey/api'
import { Survey } from 'modules/survey/types'

import DefaultTemplate from 'components/templates/DefaultTemplate'
import ContentWrapper from 'components/atoms/ContentWrapper'
import InitialLoading from 'components/atoms/InitialLoading'
import SurveyItem from 'components/organisms/SurveyItem'
import QuestionItem from 'components/pages/SurveyResult/QuestionItem'

const SurveyResult: React.FC = () => {
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
    <DefaultTemplate title="Edit survey">
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
  const handleDelete = () => {
    Router.replace(`/`)
  }
  return (
    <>
      <SurveyItem survey={survey} onDelete={handleDelete} hideButton />
      {survey.questions.map((q) => (
        <Box mt={3} key={q.id}>
          <QuestionItem question={q} />
        </Box>
      ))}
    </>
  )
}

export default SurveyResult
