/* eslint-disable */
// @ts-nocheck
import {useState} from 'react'
import {ACTION_RECV_MESSAGE, ChatService} from '@/utils'
import {css} from '@emotion/react'
import {Chat} from '@/containers/Chat.tsx'

export function Stream() {
    const [name, setName] = useState('anonymous')
    const [show, setShow] = useState(false)
    const [event, setEvent] = useState(true)
    const [messages, sendMessage, status, isThrowingMasakari] = ChatService({
        name: 'ChatTKB',
        message: `ようこそ、${name}さん`,
        action: ACTION_RECV_MESSAGE,
    })

    if (event && !show && Number(status) / 100.0 > 0.8) {
        setShow(true)
        //一度だけイベントを発火させる
        setEvent(false)
    }

    return (
        <>
            {/*<Modal setShow={setShow} show={show} />*/}

            <div
                className={''}
                style={style.back}
            >
                <Chat
                    name={name}
                    messages={messages}
                    money={false}
                    otherMoney={false}
                    sendMessage={sendMessage}
                    isThrowingMasakari={isThrowingMasakari}
                />
            </div>
        </>
    )
}

export default Stream
const style = {
    header: {
        flex: '1',
        fontSize: '3rem',
        fontFamily: 'DotGothic16',
    },
    title: {
        fontSize: '30px',
        fontWeight: 'bold',
    },
    userIcon: {
        height: '50px',
        borderRadius: '50%',
    },
    back: {
        backgroundColor: '#f0f0f0b0',
        padding: "1em"
    },
}

const viewerStyle = css`
  //margin: 2em 0;
  position: relative;
  padding: 0.5em 1.5em;
  border-top: solid 2px black;
  border-bottom: solid 2px black;

  :before,
  :after {
    content: '';
    position: absolute;
    top: -10px;
    width: 2px;
    height: -webkit-calc(100% + 20px);
    height: calc(100% + 20px);
    background-color: black;
  }

  :before {
    left: 10px;
  }

  :after {
    right: 10px;
  }
`
