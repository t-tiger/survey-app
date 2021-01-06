import { Option } from 'modules/survey/types'

type RespondentUser = { name: string; email: string }

const RESPONDENT_STORAGE_KEY = 'respondent_user'

export const totalVoteCount = (survey: {
  questions: Array<{ options: Array<Pick<Option, 'vote_count'>> }>
}) =>
  survey.questions.reduce(
    (n1, q) => n1 + q.options.reduce((n2, o) => n2 + o.vote_count, 0),
    0,
  )

export const saveRespondentUser = (user: RespondentUser) =>
  localStorage.setItem(RESPONDENT_STORAGE_KEY, JSON.stringify(user))

export const getRespondentUser = (): RespondentUser | null => {
  const val = localStorage.getItem(RESPONDENT_STORAGE_KEY)
  if (val) {
    const parsed = JSON.parse(val)
    if (
      'name' in parsed &&
      typeof parsed['name'] == 'string' &&
      'email' in parsed &&
      typeof parsed['email'] == 'string'
    ) {
      return parsed
    }
  }
  return null
}
