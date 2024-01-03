import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'

console.log(`Using VITE_API_HOST=${import.meta.env.VITE_API_HOST}`)
console.log(`Mode: ${import.meta.env.MODE}`)

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App apiHost={import.meta.env.VITE_API_HOST} />
  </React.StrictMode>,
)
