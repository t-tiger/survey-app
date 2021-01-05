import { NextRouter, useRouter } from 'next/router'

/**
 * Returns the page query, or null if page has not yet been hydrated.
 * ref: https://github.com/vercel/next.js/issues/8259#issuecomment-650225962
 */
export const useQuery = (): NextRouter['query'] | null => {
  const router = useRouter()
  const hasQueryParam = /\[.+]/.test(router.route) || /\?./.test(router.asPath)
  const ready = !hasQueryParam || Object.keys(router.query).length > 0

  if (!ready) {
    return null
  }
  return router.query
}
