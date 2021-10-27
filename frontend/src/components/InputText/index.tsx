import { ChangeEvent } from "react"

import styles from './styles.module.css'

type Props = {
  type: string
  placeholder: string
  onChange: (e: ChangeEvent<HTMLInputElement>) => void
  value: string
}

const InputText = (props: Props) => {
  
  return (
    <input className={styles.container}  {...props} />
  )
}

export default InputText