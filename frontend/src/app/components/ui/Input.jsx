export default function Input({
  id,
  name,
  type = 'text',
  className = '',
  ...props
}) {
  return (
    <input
      id={id}
      name={name || id}
      type={type}
      className={`px-2 py-1.5 w-full placeholder:italic bg-white rounded-lg shadow-(--box-shadow) focus:outline-none focus:ring-2 focus:ring-lavender-4 focus:border-lavender-4 ${className}`}
      {...props}
    />
  );
}
