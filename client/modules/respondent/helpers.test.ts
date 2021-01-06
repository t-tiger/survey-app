import { validateRespondent } from 'modules/respondent/helpers'
import { Question } from "modules/survey/types";

describe('validateRespondent', () => {
  it('name is blank', () => {
    const validate = (name: string) =>
      validateRespondent(
        { name, email: 'test@dummy.com' },
        {},
        { questions: [] },
      )
    expect(validate('')).toContain('Please input your name.')
    expect(validate(' ')).toContain('Please input your name.')
  })
  it('email is blank', () => {
    const validate = (email: string) =>
      validateRespondent({ name: 'test', email }, {}, { questions: [] })
    expect(validate('')).toContain('Please input your email.')
    expect(validate(' ')).toContain('Please input your email.')
  })
  it('some of the answers are missing', () => {
    const actual = validateRespondent(
      { name: 'test', email: "test@dummy.com" },
      {},
      { questions: [{} as Question] },
    )
    expect(actual).toContain('Please answer all questions.')
  })
  it('every field is filled', () => {
    const actual = validateRespondent(
      { name: 'test', email: "test@dummy.com" },
      {"1": "2"},
      { questions: [{options: [{id: "2"}]} as Question] },
    )
    expect(actual).toEqual([])
  })
})
