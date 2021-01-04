import { createMuiTheme } from '@material-ui/core'

import {
  SOFT_RED,
  LIGHT_GRAYISH_BLUE,
  BLUE_RIGHT,
} from 'const/color'

export const mainTheme = createMuiTheme({
  palette: {
    type: 'light',
    background: {
      default: LIGHT_GRAYISH_BLUE,
    },
    primary: {
      main: BLUE_RIGHT,
    },
    secondary: {
      main: SOFT_RED,
    },
  },
})
