import { Survey } from 'modules/survey/types'

export const totalVoteCount = (survey: Pick<Survey, 'questions'>) =>
  survey.questions.reduce(
    (n1, q) => n1 + q.options.reduce((n2, o) => n2 + o.vote_count, 0),
    0,
  )
