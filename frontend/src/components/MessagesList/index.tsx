import MessageComponent from './message'

import styles from './styles.module.css'

export type Message = {
  id: number
  sender: string
  body: string
}

type Props = {
  username: string
  messages: Message[]
}

const MessagesList = (props: Props) => {
  return (
    <ul className={styles.container}>
      {
        props.messages
          .map(message => 
            <MessageComponent 
              key={message.id} 
              username={props.username} 
              message={message}
            />)
      }
    </ul>
  )
}

export default MessagesList