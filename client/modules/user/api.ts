import axios, { AxiosResponse } from 'axios'

import { API_ENDPOINT } from 'const/index'
import { User } from 'modules/user/types'

type AuthState = {
  user: User | null
}

export const fetchAuthState = (): Promise<AxiosResponse<AuthState>> =>
  axios.get(`${API_ENDPOINT}/check_auth`)
