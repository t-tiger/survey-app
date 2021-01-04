export type Survey = {
  id: string
  title: string
  questions: Question[]
  created_at: '2021-01-03T02:32:33.049872Z'
}

export type Question = {
  id: string
  title: string
  sequence: number
  options: Option[]
}

export type Option = {
  id: string
  title: string
  sequence: number
  vote_count: number
}