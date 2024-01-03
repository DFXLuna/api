import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App({ apiHost }) {
  const [reply, setReply] = useState("reply here")

  return (
    <>
      <div>
        <h1>Button Test</h1>
      </div>
      <button onClick={() => getSlash(apiHost, setReply)}>{`Get ${apiHost}`}</button>
      <p className="reply-text">
        {reply}
      </p>
    </>
  )
}

async function getSlash(apiHost, set) {
  console.log("getSlash")
  await fetch(apiHost)
    .then(resp => resp.text())
    .then(body => set(body))
    .catch(err => console.log(err))
}

export default App
