export default function PageTitle({ children, className = 'text-center' }) {
  return (
    <h1
      className={` text-lavender-3 text-shadow-(--text-shadow) w-full  break-all text-2xl lg:text-[2rem] ${className}`}
    >
      {children}
    </h1>
  );
}
