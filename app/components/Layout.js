import React from 'react'
import Head from 'next/head'
import Script from 'next/script'
import Footer from './Footer'

export default function Layout({children}) {
  return (
    <>
		<Head>
			<meta charSet="UTF-8"/>
			<meta httpEquiv="X-UA-Compatible" content="IE=edge"/>
			<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
			
			<title>Linky - Link Shortener</title>
		</Head>
		
		<section className="flex justify-center">
			<div className="flex flex-col mt-24 max-w-7xl h-screen justify-between">
				<main>
					{children}
				</main>

				<div className="flex justify-center w-full">
					<Footer/>
				</div>
			</div>

			
		</section>

		<Script src="https://kit.fontawesome.com/4cf9184a24.js" crossOrigin="anonymous"></Script>
    </>
  )
}