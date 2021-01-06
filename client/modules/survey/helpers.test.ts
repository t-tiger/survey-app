import { totalVoteCount } from 'modules/survey/helpers'

describe('totalVoteCount', () => {
  it('calculated correctly', () => {
    const survey = {
      questions: [
        { options: [{ vote_count: 1 }, { vote_count: 3 }] },
        { options: [{ vote_count: 9 }] },
      ],
    }
    expect(totalVoteCount(survey)).toEqual(13)
  })
  it('returns 0 if question is empty', () => {
    const survey = { questions: [] }
    expect(totalVoteCount(survey)).toEqual(0)
  })
})
