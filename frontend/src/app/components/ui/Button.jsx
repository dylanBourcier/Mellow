import Link from 'next/link';
import { icons } from '@/app/lib/icons';

export default function Button({
  href,
  children,
  onClick,
  className = '',
  type = 'button',
  disabled = false,
  isSecondary = false,
  icon,
  childrenClassName = '',
}) {
  
  if (!className.includes('bg-')) {
    if (isSecondary) {
      className +=
        ' text-lavender-5 bg-transparent border border-lavender-5 hover:bg-white hover:bg-dark-grey disabled:text-white disabled:border-transparent';
    } else {
      className += ' text-white bg-lavender-3 hover:bg-lavender-5';
    }
  }
  className += icon ? ' py-1.5' : ' py-2';
  
  // if (isSecondary) {
  //   className +=
  //     ' text-lavender-5 bg-transparent border border-lavender-5 hover:bg-white hover:bg-dark-grey disabled:text-white disabled:border-transparent';
  // } else {
  //   className += ' text-white bg-lavender-3 hover:bg-lavender-5';
  // }
  // className += icon ? ' py-1.5' : ' py-2';

  const button = (
    <button
      type={type}
      onClick={onClick}
      disabled={disabled}
      className={`flex align-middle gap-1 cursor-pointer justify-center rounded-2xl px-4 transition-colors duration-200 shadow-(--box-shadow) disabled:bg-dark-grey-lighter ${className}`}
    >
      <span className={childrenClassName}>{children}</span>
      {icon && icons[icon] && <span className="w-6 h-6">{icons[icon]}</span>}
    </button>
  );
  if (href) {
    return (
      <Link href={href} className="flex w-full">
        {button}
      </Link>
    );
  }
  return button;
}
