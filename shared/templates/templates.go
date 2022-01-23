package templates

var ComponentsJSX string=`
import React from 'react'

type {{ComponentsName}}Props = {

}

const {{ComponentsName}}:React.FC<{{ComponentsName}}Props> = (props) =>{
  return (
    <div>
     {/* This is a {{ComponentsName}} Components */}
    </div>
  )
}

export default {{ComponentsName}}
`

var IndexJSX =`
import {{ComponentsName}} from './{{ComponentsName}}'

export default {{ComponentsName}}
` 