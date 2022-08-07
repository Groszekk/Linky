import React from 'react'
import Link from 'next/link'

function Footer() {
    return (
        <footer className="text-center lg:text-left bg-gray-100 text-gray-600 w-full max-w-6xl">
            <div className="flex justify-center items-center lg:justify-between p-6 border-b border-gray-300">
            </div>
            <div className="mx-6 py-10 text-center md:text-left">
                <div className="grid grid-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
                    <div>
                        <h6 className="uppercase font-semibold mb-4 flex items-center justify-center md:justify-start">
                            Linky
                        </h6>
                        <p>
                            Link Shortener
                        </p>
                    </div>

                    <div>
                        <h6 className="uppercase font-semibold mb-4 flex justify-center md:justify-start">
                        Contact
                        </h6>
                        <p className="flex items-center justify-center md:justify-start mb-4">
                            <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="envelope"
                                className="w-4 mr-4" role="img" xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 512 512">
                                <path fill="currentColor"
                                d="M502.3 190.8c3.9-3.1 9.7-.2 9.7 4.7V400c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V195.6c0-5 5.7-7.8 9.7-4.7 22.4 17.4 52.1 39.5 154.1 113.6 21.1 15.4 56.7 47.8 92.2 47.6 35.7.3 72-32.8 92.3-47.6 102-74.1 131.6-96.3 154-113.7zM256 320c23.2.4 56.6-29.2 73.4-41.4 132.7-96.3 142.8-104.7 173.4-128.7 5.8-4.5 9.2-11.5 9.2-18.9v-19c0-26.5-21.5-48-48-48H48C21.5 64 0 85.5 0 112v19c0 7.4 3.4 14.3 9.2 18.9 30.6 23.9 40.7 32.4 173.4 128.7 16.8 12.2 50.2 41.8 73.4 41.4z">
                                </path>
                            </svg>
                                paul@gross.codes
                        </p>

                    </div>
                </div>
            </div>
            <div className="text-center p-6 bg-gray-200">
                <Link href="/" className="text-gray-600 font-semibold">
                    Linky
                </Link>
            </div>
            </footer>
    )
}

export default Footer