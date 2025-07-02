import Sidebar from '../components/layout/Sidebar';
import SidebarMobile from '../components/layout/SidebarMobile';

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
        <div className="flex h-full w-screen max-w-7xl justify-center items-start">
          <Sidebar />
          <main className="flex-1">{children}</main>
          <SidebarMobile />
        </div>
      </body>
    </html>
  );
}
