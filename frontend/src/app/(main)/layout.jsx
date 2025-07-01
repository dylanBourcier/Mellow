import SidebarMobile from '../components/layout/SidebarMobile';

export default function MainLayout({ children }) {
  return (
    <html lang="en">
      <body>
        <div className="flex h-screen h-dvh w-screen max-w-7xl justify-center items-start">
          {children}
          <SidebarMobile />
        </div>
      </body>
    </html>
  );
}
