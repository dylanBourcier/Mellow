'use client';

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
  wFull = true,
}) {
  if (!className.includes('bg-')) {
    if (isSecondary) {
      className +=
        ' text-lavender-5 bg-transparent border border-lavender-5 hover:bg-lavender-6 disabled:text-white disabled:border-transparent';
    } else {
      className +=
        ' text-white bg-lavender-3 hover:bg-lavender-5 shadow-(--box-shadow)';
    }
  }
  className += icon ? ' py-1.5' : ' py-2';

  const button = (
    <button
      type={type}
      onClick={onClick}
      disabled={disabled}
      className={`flex align-middle gap-1 h-fit cursor-pointer justify-center rounded-2xl px-4 transition-colors duration-200  disabled:bg-dark-grey-lighter ${className}`}
    >
      <span className={childrenClassName}>{children}</span>
      {icon && icons[icon] && <span className="w-6 h-6">{icons[icon]}</span>}
    </button>
  );
  if (href) {
    return (
      <Link href={href} className={`flex` + (wFull ? ' w-full' : '')}>
        {button}
      </Link>
    );
  }
  return button;
}
