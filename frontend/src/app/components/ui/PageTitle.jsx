export default function PageTitle({ children, className = '' }) {
  return (
    <h1
      className={` text-lavender-3 text-shadow-(--text-shadow) w-full text-center ${className}`}
    >
      {children}
    </h1>
  );
}
