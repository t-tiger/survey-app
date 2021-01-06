import { totalVoteCount, validateSurvey } from 'modules/survey/helpers'

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

describe('validateSurvey', () => {
  it('title is blank', () => {
    const validate = (title: string) => validateSurvey(title, [])
    expect(validate('')).toContain('Please input survey title.')
    expect(validate(' ')).toContain('Please input survey title.')
  })
  it('question is blank', () => {
    expect(validateSurvey('test', [])).toContain(
      'At least one question is required.',
    )
  })
  it('some of the question titles are blank', () => {
    const actual = validateSurvey('test', [
      { title: 'q1', options: [] },
      { title: ' ', options: [] },
    ])
    expect(actual).toContain('Please input question title.')
  })
  it('some of the options are blank', () => {
    const actual = validateSurvey('test', [
      { title: 'q1', options: [{ title: 'o1' }] },
      { title: 'q2', options: [{ title: '' }] },
    ])
    expect(actual).toContain('At least one option is required for every question.')
  })
  it('every field is filled', () => {
    const actual = validateSurvey('test', [
      { title: 'q1', options: [{ title: 'o1' }, { title: 'o2' }] },
      { title: 'q2', options: [{ title: 'o3' }] },
    ])
    expect(actual).toEqual([])
  })
})
