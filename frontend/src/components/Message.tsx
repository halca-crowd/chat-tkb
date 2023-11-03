import { Emotions } from '@/types.ts'
import { format } from 'date-fns'
import ja from 'date-fns/locale/ja'
import { useState } from 'react'

interface Props {
  name: string
  message: string
  prompt: string
  emotions: Emotions
}

const initState = {
  name: '',
  message: '',
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
  prompt: '',
}
export const Message = (state: Props = initState) => {
  // 状態管理のためのStateを用意
  const [isDropdownOpen, setDrowdownOpen] = useState(false)

  // プルダウンの表示状態をトグル
  const toggleDropdown = () => {
    setDrowdownOpen(!isDropdownOpen)
  }

  return (
    <div className={'d-flex flex-column m-2'}>
      <div
        style={{
          borderRadius: '10px 10px 0px 0px',
          background: state.emotions
            ? colorChanger(state.emotions)
            : '#69E3F3',
        }}
      >
        <p
          style={{
            padding: '1em 1em 0em 1em',
            textAlign: 'left',
            fontFamily: 'Kiwi Maru',
            textOverflow: 'ellipsis',
            whiteSpace: 'nowrap',
            overflow: 'hidden',
            fontSize: '.8rem',
          }}
        >
        </p>
        <p
          style={{
            padding: '0em 1em 1em 1em',
            fontSize: '1.5rem',

            marginBottom: '0',
            color: '#0a0a0a',
            fontFamily: 'Kiwi Maru',
            overflowWrap: 'break-word',
          }}
        >
          {state.prompt}
        </p>
        <p
          style={{
            padding: '0em 1em 1em 1em',
            fontSize: '1.5rem',

            marginBottom: '0',
            color: '#0a0a0a',
            fontFamily: 'Kiwi Maru',
            overflowWrap: 'break-word',
          }}
        >
        </p>    
      </div>
      <div className={'d-flex flex-column'}>
        <div style={{
          borderRadius: '0px 0px 10px 10px',
          background: state.emotions
            ? colorChanger(state.emotions)
            : '#19C37D',
        }}>
         <button style={{
          width: '100%',
          color: '#fafafa',
          border: 'none',
          borderRadius: '0px 0px 10px 10px',
          background: state.emotions
            ? colorChanger(state.emotions)
            : '#19C37D',
         }} onClick={toggleDropdown}>{isDropdownOpen ? "✖ ChatTKBの回答を閉じる" : "▼ ChatTKBの回答を聞く！"} </button>   
        {isDropdownOpen && (
            <div>
              <p style={{
                padding: '0em 1em 0em 1em',
                textOverflow: 'ellipsis',
                whiteSpace: 'nowrap',
                overflow: 'hidden',
                fontSize: '1.5rem',
                color: '#fafafa',
                fontFamily: 'Kiwi Maru',

                display: 'flex',
                flexDirection: 'column',
                textAlign: 'center',
                alignItems: 'flex-center', // テキストを左揃え
              }}>{state.message}</p>
            </div>
          )}
        </div>
      </div>

      <div
        className={'d-flex flex-row justify-content-between'}
        style={style.sub}
      >
        <p>{state.name ?? 'ChatTKB'}</p>
        <p>{format(new Date(), 'yyyy/MM/dd HH:mm', { locale: ja })}</p>
      </div>
    </div>
  )
}

const Color = {
  Joy: 'yellow',
  Sadness: '#77FFFF',
  Anticipation: 'yellow',
  Surprise: 'red',
  Anger: 'red',
  Fear: 'blue',
  Disgust: 'red',
  Trust: 'yellow',
}

function colorChanger(emotions: Emotions) {
  const arr = [
    {
      key: 'joy',
      value: emotions.joy,
    },
    {
      key: 'sadness',
      value: emotions.sadness,
    },
    {
      key: 'anticipation',
      value: emotions.anticipation,
    },
    {
      key: 'surprise',
      value: emotions.surprise,
    },
    {
      key: 'anger',
      value: emotions.anger,
    },
    {
      key: 'fear',
      value: emotions.fear,
    },
    {
      key: 'disgust',
      value: emotions.disgust,
    },
    {
      key: 'trust',
      value: emotions.trust,
    },
  ]
  const result = arr.map(function (p) {
    return p.value
  })

  let emotion = ''
  arr.forEach((a) => {
    if (a.value === Math.max.apply(null, result)) {
      if (Math.max.apply(null, result) < 0.1) {
        emotion = ''
        return
      }
      emotion = a.key
    }
  })

  switch (emotion) {
    case 'joy':
      return Color.Joy
    case 'sadness':
      return Color.Sadness
    case 'anticipation':
      return Color.Anticipation
    case 'surprise':
      return Color.Surprise
    case 'anger':
      return Color.Anger
    case 'fear':
      return Color.Fear
    case 'disgust':
      return Color.Disgust
    case 'trust':
      return Color.Trust
    default:
      return '#e0e0e0'
  }
}

const style = {
  sub: {
    fontFamily: 'Noto Sans JP',
    padding: '0 20px',
  },
}
