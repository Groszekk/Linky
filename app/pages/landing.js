import React, {useState} from 'react'

function Landing() {
    const [link, setLink] = useState("")
    const [shorter, setShorter] = useState("")
    const [copyInfo, setCopyInfo] = useState("click to copy")
    const [processing, setProcessing] = useState(false)

    function ShortLink() {
        setProcessing(true)
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
            .then(result => {
                setShorter(result)
                setCopyInfo("click to copy")
                setProcessing(false)
            })
            .catch(error => console.log('error', error))
    }

    function Copy() {
        navigator.clipboard.writeText(shorter)
        setCopyInfo("copied! ✅")
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

                        {processing ?
                            <button type="button" class="inline-flex items-center py-2 px-3
                            font-semibold text-sm shadow rounded-md text-white 
                            bg-teal-500 hover:bg-teal-700 border-teal-500
                            transition ease-in-out duration-150 cursor-not-allowed">
                            <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                            </svg>
                                Processing...
                            </button>
                            :
                            <button className="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500
                            hover:border-teal-700 text-sm border-4 text-white py-1
                            px-2 rounded hover:scale-125 ease-out duration-300"
                            type="button" onClick={ShortLink}>
                                Short! 📈
                            </button>
                        }

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
                            <p className="text-gray-500 cursor-pointer" onClick={Copy}>
                                {copyInfo}
                            </p>
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