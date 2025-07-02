import Image from 'next/image';
import Sidebar from '../components/layout/Sidebar';
import SidebarMobile from '../components/layout/SidebarMobile';
import Link from 'next/link';

export const metadata = {
  title: {
    template: '%s - Mellow',
    default: 'Mellow',
  },
  description:
    'Mellow is a social media platform for developers to share their projects and connect with others.',
};

export default function MainLayout({ children }) {
  return (
    <html lang="en">
      <body>
        <div className="flex relative h-full w-screen max-w-7xl justify-center items-start">
          <Sidebar />
          <main className="lg:ml-78 flex-1 flex flex-col h-full px-2 lg:px-3">
            <Link href={'/'}>
              <Image
                src="img/logo.svg"
                alt="Mellow Logo"
                width={32}
                height={32}
                className="mx-auto"
              />
            </Link>
            <section>{children}</section>
          </main>
          <SidebarMobile />
        </div>
      </body>
    </html>
  );
}
