import React from "react";
import { useRouter } from "next/router";

const SurveyAnswer: React.FC = () => {
  const id = useRouter().query.id
  console.log(id)

  return <>answer</>
}

export default SurveyAnswer