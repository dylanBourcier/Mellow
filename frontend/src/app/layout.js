import { Inter, Quicksand } from 'next/font/google';
import { Toaster } from 'react-hot-toast';
import './globals.css';
import { UserProvider } from './context/UserContext';

const inter = Inter({
  subsets: ['latin'],
  variable: '--font-inter',
  display: 'swap',
});

const quicksand = Quicksand({
  subsets: ['latin'],
  variable: '--font-quicksand',
  display: 'swap',
});

export const metadata = {
  title: 'Mellow - Social Network',
  description: 'Connect with friends and share your moments on Mellow',
  manifest: '/manifest.json',
  icons: {
    icon: [
      { url: '/favicon.ico', sizes: 'any' },
      { url: '/img/logo.svg', type: 'image/svg+xml' },
    ],
    apple: '/img/logo.svg',
  },
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={`${inter.variable} ${quicksand.variable} antialiased`}>
        <UserProvider>
          <Toaster
            position="top-right"
            toastOptions={{
              duration: 3000,
              style: { background: 'transparent', boxShadow: 'none' },
            }}
          />
          {children}
        </UserProvider>
      </body>
    </html>
  );
}
