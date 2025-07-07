export default function Label({ htmlFor, children, className = '' }) {
  return (
    <label htmlFor={htmlFor} className={` font-heading text-xl ${className}`}>
      {children}
    </label>
  );
}
