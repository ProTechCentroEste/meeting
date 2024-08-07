import React from 'react';
import Head from 'next/head';
import Header from '@/components/Header';
import './globals.css';

const RootLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <html lang="en">
      <Head>
        <title>Meeting</title>
        <meta name="description" content="Real-time chat and video call application" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <body>
        <div>
          <Header />
          <main className="container mx-auto p-4">{children}</main>
          <footer className="bg-gray-800 text-white p-4 text-center">
            <div className="container mx-auto">
              <p>© 2024 Meeting. All rights reserved.</p>
            </div>
          </footer>
        </div>
      </body>
    </html>
  );
};

export default RootLayout;
