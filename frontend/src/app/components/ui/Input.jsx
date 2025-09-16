export default function Input({
  id,
  name,
  type = 'text',
  className = '',
  readOnly = false,
  ...props
}) {
  const readOnlyStyle = readOnly
    ? 'bg-gray-100 text-gray-500 cursor-not-allowed focus:outline-none focus:ring-0 focus:border-gray-200'
    : 'focus:outline-none focus:ring-2 focus:ring-lavender-4 focus:border-lavender-4';

  if (type === 'textarea') {
    return (
      <textarea
        id={id}
        name={name || id}
        className={`px-2 py-1.5 w-full placeholder:italic bg-white rounded-lg shadow-(--box-shadow) ${readOnlyStyle} ${className}`}
        {...props}
      />
    );
  }
  return (
    <input
      id={id}
      name={name || id}
      type={type}
      className={`px-2 py-1.5 w-full placeholder:italic placeholder:font-light bg-white rounded-lg shadow-(--box-shadow)  ${readOnlyStyle} ${className}`}
      readOnly={readOnly}
      {...props}
    />
  );
}
