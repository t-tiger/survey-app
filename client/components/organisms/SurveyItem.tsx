import React, { useContext } from 'react'
import styled from 'styled-components'

import {
  Box,
  Button,
  Divider,
  Typography,
  useTheme,
} from '@material-ui/core'
import { ContactSupport, Person } from '@material-ui/icons'

import { Survey } from 'modules/survey/types'
import { formatDate } from 'utils/date'
import { totalVoteCount } from 'modules/survey/helpers'

import AppContext from 'components/pages/AppContext'
import Link from "components/atoms/Link";

type Props = {
  survey: Survey
}

const SurveyItem: React.FC<Props> = ({ survey }) => {
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
      <Box mt={3} mb={2}>
        {survey.publisher.id === userId ? (
          <StartButton variant="contained" color="secondary" disableElevation>
            Watch Results
          </StartButton>
        ) : (
          <Link href={`/surveys/${survey.id}/answer`} noDecoration>
            <StartButton variant="contained" color="primary" disableElevation>
              Start Survey
            </StartButton>
          </Link>
        )}
      </Box>
      <Divider />
      <Box mt={2} display="flex" justifyContent="space-between" color="#888">
        <Typography variant="body2">
          {formatDate(new Date(survey.created_at))}
        </Typography>
        <Typography variant="body2">{totalVoteCount(survey)} Voted</Typography>
      </Box>
    </Container>
  )
}

const Container = styled(Box)`
  border-radius: 8px;
  background: white;
`
const StartButton = styled(Button)`
  width: 100%;
  border-radius: 20px;
`

export default SurveyItem
