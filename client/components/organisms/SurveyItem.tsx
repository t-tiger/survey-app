import React, { useContext } from 'react'
import styled from 'styled-components'

import {
  Box,
  Button,
  Divider,
  Tooltip,
  Typography,
  useTheme,
} from '@material-ui/core'
import { ContactSupport, Person } from '@material-ui/icons'

import { Survey } from 'modules/survey/types'
import { formatDate } from 'utils/date'
import { totalVoteCount } from 'modules/survey/helpers'

import AppContext from 'components/pages/AppContext'
import Link from 'components/atoms/Link'

type Props = {
  survey: Survey
  hideButton?: boolean
}

const SurveyItem: React.FC<Props> = ({ survey, hideButton = false }) => {
  const { userId } = useContext(AppContext)
  const theme = useTheme()

  return (
    <Container padding={2.5} boxShadow={2}>
      <Typography variant="h6" style={{ lineHeight: 1.4 }}>
        {survey.title}
      </Typography>
      <Box mt={1} display="flex" color="#888">
        <Box display="flex">
          <Person fontSize="small" />
          <Typography
            variant="body2"
            style={{ marginLeft: theme.spacing(0.5) }}
          >
            {survey.publisher.name}
          </Typography>
        </Box>
        <Box display="flex" ml={1.5}>
          <ContactSupport fontSize="small" />
          <Typography
            variant="body2"
            style={{ marginLeft: theme.spacing(0.5) }}
          >
            {survey.questions.length} Questions
          </Typography>
        </Box>
      </Box>
      {!hideButton && (
        <Box mt={3}>
          {survey.publisher.id === userId && <EditButton survey={survey} />}
          {survey.publisher.id === userId ? (
            <Link href={`/surveys/${survey.id}/result`} noDecoration>
              <ActionButton
                variant="contained"
                color="secondary"
                disableElevation
              >
                Check Results
              </ActionButton>
            </Link>
          ) : (
            <Link href={`/surveys/${survey.id}/answer`} noDecoration>
              <ActionButton
                variant="contained"
                color="primary"
                disableElevation
              >
                Start Survey
              </ActionButton>
            </Link>
          )}
        </Box>
      )}
      <Box marginY={2}>
        <Divider />
      </Box>
      <Box display="flex" justifyContent="space-between" color="#888">
        <Typography variant="body2">
          {formatDate(new Date(survey.created_at))}
        </Typography>
        <Typography variant="body2">{totalVoteCount(survey)} Voted</Typography>
      </Box>
    </Container>
  )
}

const EditButton: React.FC<EditProps> = ({ survey }) => {
  const alreadyVoted = totalVoteCount(survey) > 0

  const renderButton = () => (
    <ActionButton variant="contained" disabled={alreadyVoted} disableElevation>
      Edit Survey
    </ActionButton>
  )

  return (
    <Box mb={2}>
      {alreadyVoted ? (
        <Tooltip title={'You cannot edit the survey with votes.'}>
          <div>{renderButton()}</div>
        </Tooltip>
      ) : (
        <Link href={`/surveys/${survey.id}/edit`} noDecoration>
          {renderButton()}
        </Link>
      )}
    </Box>
  )
}

const Container = styled(Box)`
  border-radius: 8px;
  background: white;
`
const ActionButton = styled(Button)`
  width: 100%;
  border-radius: 20px;
`
type EditProps = {
  survey: Survey
}

export default SurveyItem
