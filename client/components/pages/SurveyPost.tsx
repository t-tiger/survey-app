import React, { useContext, useMemo } from 'react'
import NextError from 'next/error'
import Router from 'next/router'

import { issueId } from 'utils/id'
import { useMessageCenter } from 'utils/messageCenter'
import { postSurvey } from 'modules/survey/api'

import AppContext from 'components/pages/AppContext'
import DefaultTemplate from 'components/templates/DefaultTemplate'
import ContentWrapper from 'components/atoms/ContentWrapper'
import SurveyEditForm, {
  SurveyForEdit,
} from 'components/organisms/SurveyEditForm/Index'

const SurveyPost: React.FC = () => {
  const { userId } = useContext(AppContext)

  if (!userId) {
    return <NextError statusCode={404} />
  }
  return <Content />
}

const Content: React.FC = () => {
  const { showMessage } = useMessageCenter()

  const initialSurvey = useMemo(
    () => ({
      title: '',
      questions: [
        {
          id: issueId(),
          title: '',
          options: [{ id: issueId(), title: '' }],
        },
      ],
    }),
    [],
  )

  const handleSubmit = async (survey: SurveyForEdit) => {
    try {
      await postSurvey(survey)
      showMessage('success', 'Survey has been created successfully.')
      Router.push('/')
    } catch (e) {
      if (e.response?.data?.message) {
        showMessage('error', e.response.data.message)
      }
    }
  }

  return (
    <DefaultTemplate title="Post survey">
      <ContentWrapper>
        <SurveyEditForm
          survey={initialSurvey}
          submitTitle="Create new"
          onSubmit={handleSubmit}
        />
      </ContentWrapper>
    </DefaultTemplate>
  )
}

export default SurveyPost
