import React, { ReactElement, ReactNode } from 'react'
import styled from 'styled-components'
import Head from 'next/head'

import { Box } from '@material-ui/core'

import Header from 'components/organisms/Header'
import Link from 'components/atoms/Link'

type Props = {
  children: ReactNode
  title: string
}

const DefaultTemplate: React.FC<Props> = ({
  children,
  title,
}: Props): ReactElement => (
  <Root>
    <Head>
      <title>
        {title && `${title} | `}
        Survey app
      </title>
    </Head>
    <Box>
      <Header
        title={
          <Link href="/" color="inherit" noDecoration>
            Survey app
          </Link>
        }
      />
      <Main component="main">{children}</Main>
    </Box>
  </Root>
)

const Root = styled.div`
  min-height: 100vh;
`
const Main = styled(Box)`
  flex-grow: 1;
`

export default DefaultTemplate
