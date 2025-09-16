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
    <div className="flex relative min-h-[50dvh] h-[100dvh] p-1 lg:p-6  w-full max-w-7xl justify-center items-start">
      <Sidebar />
      <SidebarMobile />
      <main className="lg:ml-78 flex-1 flex flex-col px-2 lg:px-3 lg:gap-2">
        <div className=" flex justify-center items-center">
          <Link href={'/'}>
            <Image
              src="/img/logo.svg"
              alt="Mellow Logo"
              width={32}
              height={32}
              className="mx-auto"
            />
          </Link>
        </div>
        <section className="pb-24 lg:pb-0 flex-1 flex flex-col">
          {children}
        </section>
      </main>
    </div>
  );
}
