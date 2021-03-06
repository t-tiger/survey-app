import React, { useContext, useEffect, useState } from 'react'

import { Box, Fab, Grid, useTheme } from '@material-ui/core'
import { Add, PersonAdd } from '@material-ui/icons'

import { fetchSurveyList } from 'modules/survey/api'
import { useMessageCenter } from 'utils/messageCenter'
import { useToggleDialog } from 'utils/dialog'
import { Survey } from 'modules/survey/types'
import { fetchRespondentList } from 'modules/respondent/api'
import { getRespondentUser } from 'modules/survey/helpers'

import Link from 'components/atoms/Link'
import DefaultTemplate from 'components/templates/DefaultTemplate'
import InitialLoading from 'components/atoms/InitialLoading'
import AppContext from 'components/pages/AppContext'
import ContentWrapper from 'components/atoms/ContentWrapper'
import SurveyItem from 'components/organisms/SurveyItem'
import AuthDialog from 'components/organisms/AuthDialog/Index'

type Answers = { [surveyID: string]: true }

const Index: React.FC = () => {
  const [ready, setReady] = useState(false)
  const [surveys, setSurveys] = useState<Survey[]>([])
  const [answers, setAnswers] = useState<Answers>({})

  const { showMessage } = useMessageCenter()

  const fetch = async (page = 1) => {
    try {
      const {
        data: { items },
      } = await fetchSurveyList(page, 30)
      setSurveys(items)

      // fetch respondent list to detect whether each survey has been answered
      const respondent = getRespondentUser()
      if (respondent && items.length > 0) {
        const { data: rs } = await fetchRespondentList({
          ...respondent,
          surveyIds: items.map((s) => s.id),
        })
        setAnswers(rs.reduce((a, r) => ({ ...a, [r.survey_id]: true }), {}))
      }
    } catch {
      showMessage('error', 'Failed to fetch survey list.')
    } finally {
      setReady(true)
    }
  }

  useEffect(() => {
    fetch()
  }, [])

  return (
    <DefaultTemplate title="Surveys">
      {!ready ? (
        <InitialLoading />
      ) : (
        <Content surveys={surveys} answers={answers} refresh={fetch} />
      )}
    </DefaultTemplate>
  )
}

type ContentProps = {
  surveys: Survey[]
  answers: Answers
  refresh: () => void
}

const Content: React.FC<ContentProps> = ({ surveys, answers, refresh }) => {
  const theme = useTheme()

  const { userId } = useContext(AppContext)
  const [authDialogKey, isOpenAuthDialog, setOpenAuthDialog] = useToggleDialog()

  return (
    <ContentWrapper>
      <Box mb={3} textAlign="center">
        {userId ? (
          <Link href="/surveys/new" noDecoration>
            <Fab
              variant="extended"
              color="secondary"
              size="large"
              style={{ flexShrink: 0 }}
            >
              <Add />
              <span style={{ marginLeft: theme.spacing(1) }}>Post survey</span>
            </Fab>
          </Link>
        ) : (
          <Fab
            variant="extended"
            color="secondary"
            size="large"
            style={{ flexShrink: 0 }}
            onClick={() => setOpenAuthDialog(true)}
          >
            <PersonAdd />
            <span style={{ marginLeft: theme.spacing(1) }}>
              Sign Up to post survey
            </span>
          </Fab>
        )}
      </Box>
      <Grid container spacing={3}>
        {surveys.map((s) => (
          <Grid key={s.id} item xs={12} sm={6} md={4}>
            <SurveyItem
              survey={s}
              answered={s.id in answers}
              onDelete={refresh}
            />
          </Grid>
        ))}
      </Grid>
      <AuthDialog
        key={authDialogKey}
        open={isOpenAuthDialog}
        onClose={() => setOpenAuthDialog(false)}
      />
    </ContentWrapper>
  )
}

export default Index
