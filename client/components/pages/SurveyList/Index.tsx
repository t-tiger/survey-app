import React, { useContext, useEffect, useState } from 'react'

import { fetchSurveyList } from 'modules/survey/api'
import { useMessageCenter } from 'utils/messageCenter'
import { Survey } from 'modules/survey/types'

import DefaultTemplate from 'components/templates/DefaultTemplate'
import InitialLoading from 'components/atoms/InitialLoading'
import AppContext from 'components/pages/AppContext'

const Index: React.FC = () => {
  const [ready, setReady] = useState(true)
  const [surveys, setSurveys] = useState<Survey[]>([])
  const { showMessage } = useMessageCenter()

  const fetch = async (page = 1) => {
    try {
      const {
        data: { items },
      } = await fetchSurveyList(page, 30)
      setSurveys(items)
    } catch {
      showMessage('error', 'failed to fetch survey list')
    } finally {
      setReady(true)
    }
  }

  useEffect(() => {
    fetch()
  }, [])

  return (
    <DefaultTemplate title="Surveys">
      {!ready ? <InitialLoading /> : <Content surveys={surveys} />}
    </DefaultTemplate>
  )
}

type ContentProps = {
  surveys: Survey[]
}

const Content: React.FC<ContentProps> = ({ surveys }) => {
  const { userId } = useContext(AppContext)

  return <>{userId && <span>Post</span>}</>
}

export default Index
