import React, {useState} from 'react'

function Landing() {
    const [link, setLink] = useState("")
    const [shorter, setShorter] = useState("")

    function ShortLink() {
        const headers = new Headers()
        headers.append('Content-Type', 'application/json')

        const body = {
            link: link
        }

        const requestOptions = {
            method: 'POST',
            headers: headers,
            body: JSON.stringify(body),
            redirect: 'follow'
        }

        fetch('http://localhost/api/short', requestOptions)
            .then(response => response.text())
            .then(result => setShorter(result))
            .catch(error => console.log('error', error))
    }

    function Copy() {
        navigator.clipboard.writeText(shorter)
    }
    
    return (
        <div>
            <h1 className="text-center text-4xl md:text-5xl">
                Linky - Link Shortener
            </h1>

            <div className="flex flex-col w-full mt-20">

                <div className="shadow-xl">
                    <div className="flex items-center border-b border-teal-500 py-2">
                        <input className="appearance-none bg-transparent border-none w-full
                            text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none"
                            type="text" placeholder="https://google.com/"
                            value={link} onChange={(e) => setLink(e.target.value)}/>

                        <button className="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500
                            hover:border-teal-700 text-sm border-4 text-white py-1
                            px-2 rounded hover:scale-125 ease-out duration-300"
                            type="button" onClick={ShortLink}>
                            Short! 📈
                        </button>
                    </div>
                </div>

                <div className="text-md text-center mt-8">
                    <p>Linky is a free tool to reduce line size.</p>
                    <p>Just paste your link and click Short button!</p>
                </div>

                <div className="text-center mt-12">
                    {
                        shorter.length > 0 ? 
                        <>
                            <p className="text-4xl cursor-pointer" onClick={Copy}>
                                {shorter}
                            </p>
                            <p className="text-gray-500">click to copy</p>
                        </>
                    :
                        null
                    }
                </div>
                
            </div>
        </div>
    )
}

export default Landing