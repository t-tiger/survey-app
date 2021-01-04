import axios, { AxiosResponse } from 'axios'

import { API_ENDPOINT } from 'const/index'
import { User } from 'modules/user/types'

axios.defaults.withCredentials = true

type AuthStateResponse = {
  user: User | null
}

export const fetchAuthState = (): Promise<AxiosResponse<AuthStateResponse>> =>
  axios.get(`${API_ENDPOINT}/check_auth`)

type SignUpParams = {
  email: string
  name: string
  password: string
}

type SignUpResponse = {
  user: User
}

export const signUp = (
  params: SignUpParams,
): Promise<AxiosResponse<SignUpResponse>> =>
  axios.post(`${API_ENDPOINT}/users`, params)

type SignInParams = {
  email: string
  password: string
}

type SignInResponse = {
  user: User
}

export const signIn = (
  params: SignInParams,
): Promise<AxiosResponse<SignInResponse>> =>
  axios.post(`${API_ENDPOINT}/login`, params)
