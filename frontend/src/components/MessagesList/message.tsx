import { Message } from "."
import styles from './styles.module.css'

type Props = {
  message: Message
  username: string
}

const MessageComponent = ({message, username}: Props) => {
  return (
    <li  className={styles.line}>
      <span className={styles.sender}>{message.sender === username ? 'Você' : message.sender} disse</span>
      <span  className={styles.twoDots}>: </span>
      <span  className={styles.body}>{message.body}</span>
    </li>
  )
}

export default MessageComponent