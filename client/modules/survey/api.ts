import axios, { AxiosResponse } from 'axios'

import { API_ENDPOINT } from 'const/index'
import { Survey } from 'modules/survey/types'

type SurveyListResponse = {
  total_count: number
  items: Survey[]
}

export const fetchSurveyList = (
  page: number,
  count: number,
): Promise<AxiosResponse<SurveyListResponse>> =>
  axios.get(`${API_ENDPOINT}/surveys`, { params: { page, count } })

type SurveyResponse = Survey

export const fetchSurvey = (
  id: string,
): Promise<AxiosResponse<SurveyResponse>> =>
  axios.get(`${API_ENDPOINT}/surveys/${id}`)

type SaveSurveyParams = {
  title: string
  questions: Array<{
    title: string
    options: Array<{ title: string }>
  }>
}

export const postSurvey = (params: SaveSurveyParams) =>
  axios.post(`${API_ENDPOINT}/surveys`, params)

export const updateSurvey = (id: string, params: SaveSurveyParams) =>
  axios.put(`${API_ENDPOINT}/surveys/${id}`, params)

type PostRespondentParams = {
  surveyId: string
  email: string
  name: string
  optionIds: string[]
}

export const postRespondent = ({
  surveyId,
  email,
  name,
  optionIds,
}: PostRespondentParams) =>
  axios.post(`${API_ENDPOINT}/respondents`, {
    email,
    name,
    survey_id: surveyId,
    option_ids: optionIds,
  })
