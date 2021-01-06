import React, { useMemo } from 'react'
import styled from 'styled-components'

import { Box, Paper, Typography } from '@material-ui/core'

import { Question } from 'modules/survey/types'

type Props = {
  question: Question
}

const QuestionItem: React.FC<Props> = ({ question }) => {
  const percentages = useMemo(() => {
    const totalVote = question.options.reduce(
      (count, o) => count + o.vote_count,
      0,
    )
    return question.options.map(({ vote_count }) =>
      totalVote === 0 ? 0 : Math.round((vote_count / totalVote) * 100),
    )
  }, [question.options])

  return (
    <Paper>
      <Box padding={3}>
        <Typography variant="h5" style={{ lineHeight: 1.4 }}>
          Q{question.sequence}. {question.title}
        </Typography>
        {question.options.map((o, i) => (
          <BarContainer key={o.id} mt={2} paddingX={2} paddingY={1.2}>
            <BarPercentage style={{ width: `${percentages[i]}%` }} />
            <BarContent>
              <span>{o.title}</span>
              <Typography variant="caption">
                {o.vote_count} ({percentages[i]}%)
              </Typography>
            </BarContent>
          </BarContainer>
        ))}
      </Box>
    </Paper>
  )
}

const BarContainer = styled(Box)`
  position: relative;
  border-radius: 12px;
  background: #eee;
  overflow: hidden;
`
const BarPercentage = styled(Box)`
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  background: #abc9f5;
`
const BarContent = styled(Box)`
  position: relative;
  z-index: 1;
  display: flex;
  justify-content: space-between;
`

export default QuestionItem
