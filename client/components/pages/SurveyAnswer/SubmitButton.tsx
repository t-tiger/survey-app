import React, { useContext, useEffect, useState } from 'react'
import styled from 'styled-components'

import { Box, Button, Tooltip } from '@material-ui/core'

import { postRespondent } from 'modules/survey/api'
import { useMessageCenter } from 'utils/messageCenter'

import SurveyAnswerContext from 'components/pages/SurveyAnswer/Context'

const SubmitButton: React.FC = () => {
  const [sending, setSending] = useState(false)
  const [validationErrs, setValidationErrs] = useState<string[]>([])

  const { survey, respondent, answers } = useContext(SurveyAnswerContext)
  const { showMessage } = useMessageCenter()

  useEffect(() => {
    const errs = []
    if (respondent.name.trim().length === 0) {
      errs.push('Please input your name.')
    }
    if (respondent.email.trim().length === 0) {
      errs.push('Please input your email.')
    }
    if (Object.keys(answers).length < survey.questions.length) {
      errs.push('Please answer all questions.')
    }

    setValidationErrs(errs)
  }, [survey, respondent, answers])

  const handleSubmit = async () => {
    try {
      setSending(true)

      const { name, email } = respondent
      await postRespondent({
        name,
        email,
        surveyId: survey.id,
        optionIds: Object.values(answers),
      })
      showMessage('success', 'Submission has been received successfully')
    } catch (e) {
      if (e.response?.data?.message) {
        showMessage('error', e.response.data.message)
      }
    } finally {
      setSending(false)
    }
  }

  return (
    <Box mt={3.5} textAlign="center">
      <Tooltip
        title={
          validationErrs.length > 0 ? (
            <span style={{ whiteSpace: 'pre-line' }}>
              {validationErrs.join('\n')}
            </span>
          ) : (
            ''
          )
        }
      >
        {/* insert div element because disabled button also makes tooltip disabled */}
        <div>
          <StyledButton
            color="primary"
            variant="contained"
            disabled={sending || validationErrs.length > 0}
            onClick={handleSubmit}
          >
            Complete survey
          </StyledButton>
        </div>
      </Tooltip>
    </Box>
  )
}

const StyledButton = styled(Button)`
  font-size: 20px;
  padding: 10px 30px;
`

export default SubmitButton
