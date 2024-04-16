import { useState } from 'react'
import './App.css'
import logo from './assets/logo.svg'
import search from './assets/search.svg'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className='bg-primary min-w-screen min-h-screen bg-background-pattern bg-bottom bg-no-repeat bg-repeat-x p-8 md:p-8 lg:p-16' >
      {/* <img src={background} className='absolute bottom-0'/> */}
      <div>
        <img src={logo} className='h-4 w-auto md:h-6 lg:h-10'/>
      </div>
      <div className='py-5 md:py-6 lg:py-8 relative'>
        <input type='text' className='bg-white border-o rounded w-full h-12 md:h-12 lg:h-16 px-5 focus:outline-none' placeholder='Enter Keywords...'/>
        <div className='bg-secondary cursor-pointer h-12 md:h-12 lg:h-16 px-5 flex items-center justify-center absolute  right-0 top-5 md:top-6 lg:top-8 rounded-r'>
          <img src={search}/>
        </div>
      </div>
      <div className='text-white'>
        <p className='text-white' >Discover endless entertainment with U<span className='text-secondary'>m</span>Y - your ultimate destination for immersive audio and video streaming experiences.</p>
      </div>
    </div>
  )
}

export default App
