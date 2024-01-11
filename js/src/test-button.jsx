import { useState } from 'react'

function TestButton({ apiHost }) {
    const [reply, setReply] = useState('reply here')

    return (
        <section className='card w-1/3 bg-base-200 shadow-m'>
            <div className='card-body'>
                <h2 className='card-title font-serif font-light'>Button Test</h2>
                <p className='reply-text font-serif font-light'>
                    {reply}
                </p>
                <div className='card-actions justify-end'>
                    <button className='btn btn-primary rounded-xl' onClick={() => getSlash(apiHost, setReply)}>{`Get ${apiHost}`}</button>
                </div>
            </div>
        </section>
    )
}

async function getSlash(apiHost, set) {
    set(`fetching ${apiHost}`)
    await fetch(apiHost)
        .then(resp => resp.text())
        .then(body => set(body))
        .catch(err => console.log(err))
}

export default TestButton