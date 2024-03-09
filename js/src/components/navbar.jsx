function Navbar() {
    return (
        <nav className='navbar sticky bg-base-300 rounded-xl'>
            <a href='/' className='btn btn-ghost flex-none text-xl font-serif font-light'>something</a>
            <div className='flex-1'></div>
            <a href='/login' role='button' className='btn btn-neutral flex-none font-serif font-light'>login</a>
        </nav>
    )
}

export default Navbar