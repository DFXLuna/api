// import { useState } from 'react'
import TestButton from './test-button'

function App({ apiHost }) {
  // const [reply, setReply] = useState('reply here')

  return (
    <div className='container mx-auto min-h-screen'>
      <div className='flex min-h-screen flex-col items-center justify-center'>
        <TestButton apiHost={apiHost} />
        {/* <div>
          <h1>Button Test</h1>
        </div>
        <button className='btn btn-primary rounded-xl' onClick={() => getSlash(apiHost, setReply)}>{`Get ${apiHost}`}</button>
        <p className='reply-text'>
          {reply}
        </p> */}
      </div>
    </div>
  )
}

// async function getSlash(apiHost, set) {
//   set(`fetching ${apiHost}`)
//   await fetch(apiHost)
//     .then(resp => resp.text())
//     .then(body => set(body))
//     .catch(err => console.log(err))
// }

export default App
