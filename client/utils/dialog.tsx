import { useState, useCallback } from 'react'

/**
 * custom hook for toggling dialog.
 * increment key when enabling open value.
 */
export const useToggleDialog = (
  initial = false,
): [number, boolean, (open: boolean) => void] => {
  const [key, setKey] = useState(0)
  const [open, setOpen] = useState(initial)

  const handleOpen = useCallback(
    (value: boolean) => {
      if (value) {
        setKey(key + 1)
      }
      setOpen(value)
    },
    [key],
  )

  return [key, open, handleOpen]
}
