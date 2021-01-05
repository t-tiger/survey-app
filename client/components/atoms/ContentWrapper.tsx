import React, { ReactNode } from 'react'
import styled from 'styled-components'

import { CONTENT_MAX_WIDTH } from 'const/size'

type Props = {
  children: ReactNode
}

const ContentWrapper: React.FC<Props> = ({ children }) => (
  <Container>{children}</Container>
)

const Container = styled.div`
  max-width: ${CONTENT_MAX_WIDTH + (12 * 2)}px;
  margin: auto;
  padding: 25px 12px 0;
`

export default ContentWrapper
