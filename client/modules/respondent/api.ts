import axios, { AxiosResponse } from 'axios'

import { API_ENDPOINT } from 'const/index'

type RespondentListParams = {
  email: string
  name: string
  surveyIds: string[]
}

type RespondentListResponse = Array<{ id: string; survey_id: string }>

export const fetchRespondentList = ({
  email,
  name,
  surveyIds,
}: RespondentListParams): Promise<AxiosResponse<RespondentListResponse>> =>
  axios.get(`${API_ENDPOINT}/respondents`, {
    params: {
      email,
      name,
      surveyIds: surveyIds.join(','),
    },
  })

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
