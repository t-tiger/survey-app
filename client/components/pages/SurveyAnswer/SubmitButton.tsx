import React, { useContext, useEffect, useState } from 'react'
import styled from 'styled-components'
import Router from 'next/router'

import { Box, Button } from '@material-ui/core'

import { postRespondent } from 'modules/respondent/api'
import { useMessageCenter } from 'utils/messageCenter'
import { saveRespondentUser } from "modules/survey/helpers";

import SurveyAnswerContext from 'components/pages/SurveyAnswer/Context'
import MultiLineToolTip from 'components/atoms/MultiLineTooltip'

const SubmitButton: React.FC = () => {
  const [sending, setSending] = useState(false)
  const [validationErrs, setValidationErrs] = useState<string[]>([])

  const { survey, respondent, answers } = useContext(SurveyAnswerContext)
  const { showMessage } = useMessageCenter()

  const validate = () => {
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
    return errs
  }

  useEffect(() => {
    setValidationErrs(validate())
  }, [survey, respondent, answers])

  const handleSubmit = async () => {
    try {
      setSending(true)
      saveRespondentUser(respondent)

      const { name, email } = respondent
      await postRespondent({
        name,
        email,
        surveyId: survey.id,
        optionIds: Object.values(answers),
      })
      showMessage('success', 'Submission has been received successfully.')
      Router.replace(`/surveys/${survey.id}/result`)
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
      <MultiLineToolTip titles={validationErrs}>
        <StyledButton
          color="primary"
          variant="contained"
          disabled={sending || validationErrs.length > 0}
          onClick={handleSubmit}
        >
          Complete survey
        </StyledButton>
      </MultiLineToolTip>
    </Box>
  )
}

const StyledButton = styled(Button)`
  font-size: 18px;
  padding: 8px 20px;
`

export default SubmitButton
