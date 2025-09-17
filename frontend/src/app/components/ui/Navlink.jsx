import Link from 'next/link';
import { icons } from '@/app/lib/icons';
import Image from 'next/image';

export default function Navlink({
  icon,
  img,
  href,
  children,
  isActive = false,
  disabled = false,
}) {
  if (disabled) {
    return (
      <div
        className={
          'text-dark-grey-lighter flex items-center justify-center lg:justify-start rounded-md h-12 px-3 gap-2 w-fit lg:w-full '
        }
      >
        <span className="flex-shrink-0">{icons[icon]}</span>
        {children && (
          <span className="font-heading text-2xl tracking-tighter">
            {children}
          </span>
        )}
      </div>
    );
  }
  return (
    <Link
      href={href}
      className={`
        ${isActive ? 'bg-lavender-6 text-lavender-3' : 'text-dark-grey'}
      } relative hover:text-lavender-3 transition-colors duration-200 flex items-center justify-center lg:justify-start rounded-md h-12 px-3 gap-2 w-fit lg:w-full `}
    >
      {icon && <span className="flex-shrink-0">{icons[icon]}</span>}
      {img && (
        <Image
          src={img}
          alt="User Avatar"
          className="w-6 h-6 rounded-full"
          width={24}
          height={24}
        />
      )}
      {children && (
        <span className="font-heading text-2xl tracking-tighter">
          {children}
        </span>
      )}
    </Link>
  );
}
