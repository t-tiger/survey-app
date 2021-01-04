import axios, { AxiosResponse } from 'axios'

import { API_ENDPOINT } from 'const/index'
import { Survey } from 'modules/survey/types'

type SurveyList = {
  total_count: number
  items: Survey[]
}

export const fetchSurveyList = (
  page: number,
  count: number,
): Promise<AxiosResponse<SurveyList>> =>
  axios.get(`${API_ENDPOINT}/surveys`, { params: { page, count } })
