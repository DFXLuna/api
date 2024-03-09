import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import App from './app.jsx'
import PocketBase from 'pocketbase';
import Container from './container.js';

console.log(`Mode: ${import.meta.env.MODE}`)
console.log(`Using VITE_API_HOST=${import.meta.env.VITE_API_HOST}`)
console.log(`Using VITE_PB_HOST=${import.meta.env.VITE_PB_HOST}`)

const pb = new PocketBase(import.meta.env.VITE_PB_HOST);
const c = new Container(pb)


ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App apiHost={import.meta.env.VITE_API_HOST} c={c} />
  </React.StrictMode>
)
