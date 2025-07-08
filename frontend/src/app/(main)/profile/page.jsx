import ProfileScreen from '@/app/components/layout/ProfileScreen';
import React from 'react';

export default function ProfilePage() {
  return (
  <div>
    <ProfileScreen
    firstName="Jhon" 
    lastName="Doe" 
    username="Doedoe" 
    email="JhonDoe@mail.com"  
    birthdate="25/09/1970"
    followers="22" 
    following="10"
    authorAvatar = '/img/DefaultAvatar.png'
    description="Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
     />
  </div>
  );
}
