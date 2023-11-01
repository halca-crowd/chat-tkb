import { Message } from '@/components/Message'
import { ChatPropsType } from '@/utils'

interface Props {
  messages: boolean | ChatPropsType[] | ((props: ChatPropsType) => void)
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
  prompt: string
}

export const Chat = ({ messages }: Props) => {
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  return (
    <div
      style={style.body}
      className={'d-flex flex-column justify-content-between'}
    >
      <div className={'vh-100'} style={style.listBox}>
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
                  prompt={msg.prompt}
                  emotions={msg.emotions}
                />
              )
            })
          }
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
    height: '90vh',
  },
  listBox: {},
}
