export default function MainLayout({ children }) {
  return (
    <html lang="en">
      <body >
        <div className="flex h-screen w-screen max-w-7xl justify-center items-start">
        {children}
        </div>
      </body>
    </html>
  );
}
