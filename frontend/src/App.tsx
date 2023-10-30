import './App.css'
import backStyle from './App.module.css'
import { Stream } from '@/views/stream.tsx'

function App() {
  return (
    <div className={'App'}>
      <div className={backStyle.back}>
        <Stream />
      </div>
    </div>
  )
}

export default App
