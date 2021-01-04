import axios, { AxiosResponse } from 'axios'

import { API_ENDPOINT } from 'const/index'

type SurveyList = {
  authorized: boolean
}

export const fetchAuthState = (): Promise<AxiosResponse<SurveyList>> =>
  axios.get(`${API_ENDPOINT}/check_auth`)
