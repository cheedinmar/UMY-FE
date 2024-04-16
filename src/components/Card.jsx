import { Children, useState } from 'react'


function App( props) {
  const [count, setCount] = useState(0)

  return (
    <div className='h-48 md:h-52 relative w-16 bg-gradient-to-b flex items-center cursor-pointer justify-center hover:-translate-y-4 hover:scale-105 transition ease-in-out from-cardColor1 to-cardColor2 rounded shadow-2xl w-2/5 md:w-1/3 lg:w-1/5 rounded-3xl '>
        {props.children}
    </div>
  )
}

export default App
