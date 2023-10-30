import { useLayoutEffect, useRef, useState } from 'react'
import { Message } from '@/components/Message'
import {
  ACTION_SEND_MASAKARI,
  ACTION_SEND_MESSAGE,
  ChatPropsType,
} from '@/utils'

interface Props {
  name: string
  messages: boolean | ChatPropsType[] | ((props: ChatPropsType) => void)
  money: boolean
  otherMoney: boolean
  sendMessage: boolean | ChatPropsType[] | ((props: ChatPropsType) => void)
  isThrowingMasakari: boolean
}

export interface Emotions {
  joy: number
  sadness: number
  anticipation: number
  surprise: number
  anger: number
  fear: number
  disgust: number
  trust: number
}

// const initState: Props = { name: '', text: '', money: false, otherMoney: false }

interface ChatProps {
  name: string
  message: string
  emotions: Emotions
}

export const Chat = ({ name, messages, sendMessage }: Props) => {
  const scrollBottomRef = useRef<HTMLDivElement>(null)
  const [text, setText] = useState('')

  const submitMessage = () => {
    if (text.length === 0) return
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    sendMessage({
      message: text,
      name: name,
      action: ACTION_SEND_MASAKARI,
    })
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    sendMessage({
      message: text,
      name: name,
      action: ACTION_SEND_MESSAGE,
    })

    setText('')
  }
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  const handleInputChange = (e) => {
    setText(e.target.value)
  }
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  const handleButtonClick = () => {
    submitMessage()
  }
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  const handleOnKeydown = (event) => {
    if (event.keyCode == 13) {
      submitMessage()
    }
  }
  useLayoutEffect(() => {
    if (scrollBottomRef && scrollBottomRef.current) {
      scrollBottomRef?.current?.scrollIntoView()
    }
  })
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  return (
    <div
      style={style.body}
      className={'d-flex flex-column justify-content-between'}
    >
      <div className={'overflow-scroll vh-100'} style={style.listBox}>
        <ul>
          {
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            //@ts-ignore
            messages.map((msg: ChatProps, idx: number) => {
              return (
                <Message
                  key={idx}
                  name={msg.name}
                  message={msg.message}
                  emotions={msg.emotions}
                />
              )
            })
          }
          <div ref={scrollBottomRef}></div>
        </ul>
      </div>
      {/*<div className={'input-group mb-3'}>*/}
      {/*    <input*/}
      {/*        type="text"*/}
      {/*        placeholder="メッセージ"*/}
      {/*        value={text}*/}
      {/*        className={'form-control'}*/}
      {/*        onChange={handleInputChange}*/}
      {/*        onKeyDown={handleOnKeydown}*/}
      {/*    />*/}
      {/*    <button disabled={!text} onClick={handleButtonClick}>*/}
      {/*        送信*/}
      {/*    </button>*/}
      {/*</div>*/}
    </div>
  )
}

const style = {
  body: {
    height: '560px',
  },
  listBox: {},
}
