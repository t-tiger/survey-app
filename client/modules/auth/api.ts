import axios, { AxiosResponse } from 'axios'

import { API_ENDPOINT } from 'const/index'

type AuthState = {
  authorized: boolean
}

export const fetchAuthState = (): Promise<AxiosResponse<AuthState>> =>
  axios.get(`${API_ENDPOINT}/check_auth`)
