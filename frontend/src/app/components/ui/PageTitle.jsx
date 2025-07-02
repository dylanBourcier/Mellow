export default function PageTitle({ children, className = '' }) {
  return (
    <h1
      className={`tracking-tighter text-lavender-3 text-shadow-xs w-full text-center ${className}`}
    >
      {children}
    </h1>
  );
}
