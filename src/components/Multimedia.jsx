import { Children, useState } from 'react'


function App( {image , title, author}) {
  const [count, setCount] = useState(0)

  return (
    <div className='flex flex-col items-center justify-center' >
        <img className='w-[87px] h-[87px] rounded-full ' src={image}/>
        <p className='text-white'>{title}</p>
        <p className='text-secondary text-sm'>{author}</p>
        
    </div>
  )
}

export default App
