import Image from 'next/image';
import Link from 'next/link';
import { icons } from '@/app/lib/icons';

export default function UserInfo({
    userName = 'johndoe',
    authorAvatar = '/img/DefaultAvatar.png',
    groupId, 
    userId,
    groupName,

    }) {

    return (
        <div className="flex items-center gap-2">
        <Image
            src={authorAvatar}
            width={32}
            height={32}
            alt="Author Avatar"
            className="w-8 h-8 rounded-full">
        </Image>
        <span className='flex gap-1'>
           <Link className="hover:underline" href={`/user/${userId}`}>{userName}</Link> {groupId && groupName ? (
               <>
                   · <Link className='font-semibold flex items-center gap-1 hover:underline' href={`/groups/${groupId}`}>{icons["groups"]}{groupName}</Link>
               </>
           ) : ""}
        </span>
        </div>
    );
    }