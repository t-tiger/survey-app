import React, { useContext, useEffect, useState } from 'react'
import NextError from 'next/error'
import Router, { useRouter } from 'next/router'

import { useMessageCenter } from 'utils/messageCenter'
import { fetchSurvey, updateSurvey } from 'modules/survey/api'
import { Survey } from 'modules/survey/types'

import AppContext from 'components/pages/AppContext'
import DefaultTemplate from 'components/templates/DefaultTemplate'
import ContentWrapper from 'components/atoms/ContentWrapper'
import SurveyEditForm, {
  SurveyForEdit,
} from 'components/organisms/SurveyEditForm/Index'
import InitialLoading from 'components/atoms/InitialLoading'

const SurveyEdit: React.FC = () => {
  const id = useRouter().query.id as string
  const { userId } = useContext(AppContext)

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
    if (userId) {
      fetch()
    }
  }, [])

  if (!userId || (ready && !survey)) {
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
  const { showMessage } = useMessageCenter()

  const handleSubmit = async (updated: SurveyForEdit) => {
    try {
      await updateSurvey(survey.id, updated)
      showMessage('success', 'Survey has been updated successfully.')
      Router.push('/')
    } catch (e) {
      if (e.response?.data?.message) {
        showMessage('error', e.response.data.message)
      }
    }
  }

  return (
    <SurveyEditForm
      survey={survey}
      pageTitle="Edit survey"
      submitTitle="Update survey"
      onSubmit={handleSubmit}
    />
  )
}

export default SurveyEdit
