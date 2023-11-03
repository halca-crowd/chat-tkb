import { BottomObj } from '@/views/BottomObj.tsx'
import { NavBar } from '@/views/NavBar.tsx'
import { Stream } from '@/views/stream.tsx'
import './App.css'
import backStyle from './App.module.css'

function App() {
  return (
    <div className={'App'}>
      <NavBar />
      <div className={backStyle.back}>
        <Stream />
      </div>
      <BottomObj />
    </div>
  )
}

export default App
