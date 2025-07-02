import Image from 'next/image';
import Link from 'next/link';
import { icons } from '@/app/lib/icons';

export default function Navlink({ icon, href, children, isActive = false }) {
  return (
    <Link
      href={href}
      className={`
        ${isActive ? 'bg-lavender-6 text-lavender-3' : 'text-dark-grey'}
      } hover:text-lavender-3 transition-colors duration-200 flex items-center justify-center lg:justify-start rounded-md h-12 px-3 gap-2 w-fit lg:w-full `}
    >
      <span className="flex-shrink-0">{icons[icon]}</span>
      {children && <span className='font-heading text-2xl tracking-tighter'>{children}</span> }
      
    </Link>
  );
}
