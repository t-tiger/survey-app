import { Survey } from 'modules/survey/types'

export const validateRespondent = (
  respondent: { name: string; email: string },
  answers: { [questionId: string]: string },
  survey: Pick<Survey, 'questions'>,
): string[] => {
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
