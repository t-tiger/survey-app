import React, { ReactNode } from 'react'
import { Box, Tooltip } from '@material-ui/core'

type Props = {
  titles: string[]
  children: ReactNode
}

const MultiLineToolTip: React.FC<Props> = ({ titles, children }) => (
  <Tooltip
    title={
      titles.length > 0 ? (
        <span style={{ whiteSpace: 'pre-line' }}>{titles.join('\n')}</span>
      ) : (
        ''
      )
    }
  >
    {/* If disabled element is included, tooltip will be also disabled, so div element is located here to prevent */}
    <div>{children}</div>
  </Tooltip>
)

export default MultiLineToolTip
