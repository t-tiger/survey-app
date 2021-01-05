import React, { useContext, useMemo } from 'react'
import NextError from 'next/error'

import { issueId } from 'utils/id'

import AppContext from 'components/pages/AppContext'
import DefaultTemplate from 'components/templates/DefaultTemplate'
import ContentWrapper from 'components/atoms/ContentWrapper'
import SurveyEditForm from 'components/organisms/SurveyEditForm/Index'

const SurveyPost: React.FC = () => {
  const { userId } = useContext(AppContext)

  if (!userId) {
    return <NextError statusCode={404} />
  }
  return <Content />
}

const Content: React.FC = () => {
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

  return (
    <DefaultTemplate title="Post survey">
      <ContentWrapper>
        <SurveyEditForm survey={initialSurvey} />
      </ContentWrapper>
    </DefaultTemplate>
  )
}

export default SurveyPost
