import Image from 'next/image';
import Sidebar from '../components/layout/Sidebar';
import SidebarMobile from '../components/layout/SidebarMobile';
import Link from 'next/link';
import { icons } from '@/app/lib/icons';

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
    <div className="flex relative h-full w-screen max-w-7xl justify-center items-start">
      <Sidebar />
      <SidebarMobile />
      <main className="lg:ml-78 flex-1 flex flex-col h-full px-2 lg:px-3">
        <div className=" flex relative justify-center items-center">
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
        <section className="pb-24 lg:pb-0">{children}</section>
      </main>
    </div>
  );
}
