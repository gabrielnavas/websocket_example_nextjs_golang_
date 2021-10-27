import { useCallback, useState } from "react"

import styles from './styles.module.css'

import InputText from "../../components/InputText"
import MessagesList, { Message } from "../../components/MessagesList"
import Title from "../../components/Title"

const baseurl = 'ws://localhost:8000/chat'

const ChatPage = () => {
  const [ws, setWs] = useState<WebSocket | undefined>(undefined)
  const [connected, setConnected] = useState(false)
  const [username, setUsername] = useState('')
  const [message, setMessage] = useState('')
  const [messages, setMessages] = useState([] as Message[])

  const handlerEnterChat = useCallback(() => {
    const ws = new WebSocket(`${baseurl}?username=${username}`)
    ws.onopen = event => {
      console.log("Websocket opened", {event})
    }

    ws.close = event => {
      setWs(undefined)
      setConnected(false)
    }

    ws.onmessage = event => {
      const message = JSON.parse(event.data) as Message
      setMessages(old => [message, ...old])
      console.log("Websocket message!", {event})
    }

    ws.onerror = error => {
      console.log("Websocket error!", {error})
    }

    setWs(ws)
    setConnected(true)
  }, [username])

  const handleSendMessage = useCallback(() => {
    if (message.length > 0) {
      ws?.send(message)
      setMessage('')
    }
  }, [message, ws])


  return (
    <div className={styles.container}>
      <Title title={connected ? username : 'Anônimo'} />
      {
        connected && <MessagesList username={username} messages={messages}  />
      }
      <form className={styles.form}>
        <InputText 
          placeholder={connected ? 'Escreva uma mensagem...' : 'Entre com seu nome...'}
          onChange={e => connected ? setMessage(e.target.value) : setUsername(e.target.value)}
          type='text'
          value={ws? message : username}
        />
        <button
          className={styles.buttonSubmit}
          type='button'
          onClick={connected ? handleSendMessage : handlerEnterChat}>
          {connected ? 'Enviar':'Entrar'}
        </button>
      </form>
    </div>
  )
}

export default ChatPage