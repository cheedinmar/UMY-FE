import { useState } from 'react'
import './App.css'
import logo from './assets/logo.svg'
import search from './assets/search.svg'
import Card from './components/Card'
import musicIcon from './assets/musicIcon.svg'
import videoIcon from './assets/videoIcon.svg'
import circleIcon from './assets/circle.svg'
import yellowcircle from './assets/yellowcircle.svg'
import Multimedia from './components/Multimedia'
import music from './assets/images/music.png'
import wonka from './assets/images/wonka.png'
import shogun from './assets/images/shogun.png'

function App() {
  const [count, setCount] = useState(0)
  const Multimedias = [
    {
      image: wonka,
      title:'Willy Wonka',
      author:'Hulu Films'
    },
    {
      image:music,
      title:'Johnny Drille',
      author:'Stay with me'
    },
    {
      image:shogun,
      title:'Shogun',
      author:'Paramount'
    },
  ]

  return (
    <div className='bg-primary min-w-screen min-h-screen bg-background-pattern bg-bottom bg-no-repeat bg-repeat-x p-8 md:p-8 lg:p-16' >
      {/* <img src={background} className='absolute bottom-0'/> */}
      <div>
        <img src={logo} className='h-4 w-auto md:h-6 lg:h-10'/>
      </div>
      <div className='pt-5 md:pt-6 lg:pt-8 relative pb-4'>
        <input type='text' className='bg-white border-o rounded w-full h-12 md:h-12 lg:h-16 px-5 focus:outline-none' placeholder='Enter Keywords...'/>
        <div className='bg-secondary cursor-pointer h-12 md:h-12 lg:h-16 px-5 flex items-center justify-center absolute  right-0 top-5 md:top-6 lg:top-8 rounded-r'>
          <img src={search}/>
        </div>
      </div>
      <div >
        <p className='text-white' >Discover endless entertainment with U<span className='text-secondary'>m</span>Y - your ultimate destination for immersive audio and video streaming experiences.</p>
      </div>
      <div className='flex items-center justify-between py-8 md:py-10 lg:py-12 md:justify-evenly'>
        <Card>
          <img src={circleIcon} className='absolute left-3  top-5'/>
          <img src={circleIcon} className='absolute right-8 bottom-2'/>
          <img src={yellowcircle} className='absolute top-1/3 right-2'/>
          <img src={videoIcon}/>
        </Card>
        <Card>
        <img src={circleIcon} className='absolute left-3  top-5'/>
          <img src={circleIcon} className='absolute right-8 bottom-2'/>
          <img src={yellowcircle} className='absolute bottom-1/3 left-2'/>
        <img src={musicIcon}/>
        </Card>
      </div>
      <div className='mt-4'>
        <p className='text-white'>Trending Activities</p>
        <div className='flex items-center justify-between'>
        {Multimedias.map((item, index) => (
        <Multimedia image={item.image} title={item.title} author={item.author}/>
      ))}
          
        </div>

      </div>
    </div>
  )
}

export default App
