import React, { useState, } from 'react'

import DefaultTemplate from 'components/templates/DefaultTemplate'
import InitialLoading from "components/atoms/InitialLoading";

const Index: React.FC = () => {
  const [loading, setLoading] = useState(true)

  return (
    <DefaultTemplate title={'Surveys'}>
      {loading ? (
        <InitialLoading />
      ) : (
        <>test</>
      )}
    </DefaultTemplate>
  )
}

export default Index
