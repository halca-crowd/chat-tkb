import {useEffect, useRef, useState} from 'react'
import {
    ACTION_RECV_MASAKARI,
    ACTION_RECV_STATUS,
    ACTION_SEND_MESSAGE,
    randomStr,
} from '@/utils/constants'

interface Props {
    name: string
    message: string
    action: string
}

export type ChatPropsType = Props

export const ChatService = (props: Props) => {
    const SOCKET_URL =
        import.meta.env.VITE_WS_URL || 'ws://localhost:8081/ws/masakari'
    const [messages, setMessages] = useState([props])
    const [status, setStatus] = useState(0.0)
    const socketRef = useRef(null)
    const [isPaused, setPause] = useState(false)
    const [gptMessage, setGptMessage] = useState(
        '進捗どうですか？進捗どうですか？進捗どうですか？進捗どうですか？進捗どうですか？進捗どうですか？進捗どうですか？進捗どうですか？進捗どうですか？進捗どうですか？進捗どうですか？進捗どうですか？'
    )
    const [isThrowingMasakari, setThrowingMasakari] = useState(false)
    const [isBreakingWinddow, setBreakingWindow] = useState(false)


    useEffect(() => {
        //console.log('Connectinng..')
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        socketRef.current = new WebSocket(SOCKET_URL)
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        socketRef.current.onopen = () => console.log('socketRef opened')
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        socketRef.current.onclose = () => console.log('socketRef closed')

        const socketRefCurrent = socketRef.current

        return () => {
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore
            socketRefCurrent.close()
        }
    }, [])

    useEffect(() => {
        if (!socketRef.current) return

        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        socketRef.current.onmessage = (e) => {
            if (isPaused) return
            //console.log(e.data)
            const message = JSON.parse(e.data)
            if (message.action == ACTION_SEND_MESSAGE) {
                setMessages((prevMessages) => [...prevMessages, message])
            } else if (message.action == ACTION_RECV_MASAKARI) {
                setThrowingMasakari(false)
                setBreakingWindow(false)
                setGptMessage(message.message)
            } else if (message.action == ACTION_RECV_STATUS) {
                console.log(message.status)
                setStatus(message.status.cpuutilization)
            }
        }
    }, [isPaused])

    const sendMessage = (props: Props) => {
        setThrowingMasakari(true)
        const aMessage = {
            // name: props.name,
            name: randomStr + 'さん',
            message: props.message,
            action: props.action,
            user_id: randomStr,
            emotions: {
                joy: 0,
                sadness: 0,
                anticipation: 0,
                surprise: 0,
                anger: 0,
                fear: 0,
                disgust: 0,
                trust: 0,
            },
        }
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        socketRef.current.send(JSON.stringify(aMessage))
        // setMessages((prevMessages) => [...prevMessages, aMessage])
    }

    // マサカリを投げた後、
    useEffect(() => {
        if (isThrowingMasakari) {
            setBreakingWindow(true)
        }

        return () => {
            setBreakingWindow(false)
        }
    }, [isThrowingMasakari])

    return [
        messages,
        sendMessage,

        otherMoney,
        gptMessage,
        status,
        isThrowingMasakari,
        isBreakingWinddow,
    ]
}