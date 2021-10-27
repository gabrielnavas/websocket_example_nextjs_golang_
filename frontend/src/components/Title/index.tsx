import styles from './styles.module.css'

type Props = {
  title: string
}

const Title = ({title}: Props) => {
  return (
    <div className={styles.container}>
     <span>Seu nome na sala é</span>
     <span>:</span>
     <span className={styles.title}>{title}</span>
    </div>
  )
}

export default Title