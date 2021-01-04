import React, { ReactElement, ReactNode } from 'react'
import styled from 'styled-components'
import Head from 'next/head'

import { Box } from '@material-ui/core'

import Header from 'components/organisms/Header'

type Props = {
  children: ReactNode
  title: string
  headerTitle?: ReactNode
}

const DefaultTemplate: React.FC<Props> = ({
  children,
  title,
  headerTitle,
}: Props): ReactElement => {
  return (
    <Root>
      <Head>
        <title>
          {title && `${title} | `}
          Survey app
        </title>
      </Head>
      <Box>
        <Header title={headerTitle || title} />
        <Main component="main">{children}</Main>
      </Box>
    </Root>
  )
}

const Root = styled.div`
  min-height: 100vh;
`
const Main = styled(Box)`
  flex-grow: 1;
`

export default DefaultTemplate
