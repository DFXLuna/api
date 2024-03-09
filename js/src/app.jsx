import TestButton from './pages/test-button'
import Navbar from './components/navbar'
import Login from './pages/login'
import Playlist from './components/playlist'

function App({ apiHost, c }) {
  return (
    <div>
      <Navbar />
      <div className='container mx-auto min-h-screen'>
        <div className='flex min-h-screen flex-col items-center justify-center'>
          {/* <TestButton apiHost={apiHost} /> */}
          {/* <Login c={c} /> */}
          <Playlist />
        </div>
      </div>
    </div>
  )
}

export default App
