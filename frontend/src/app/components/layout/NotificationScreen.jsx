"use client";

import React from 'react';
import PageTitle from '../ui/PageTitle';
import NotificationsCard from '../ui/NotificationsCard';

function NotificationScreen() {
    const notifications=[
        {
            id:'1',
            type:'followed',
            username:'Doedoe',
            avatarUrl:'/img/lion.png',
            timestamp:'21:45'
        },
        {
            id:'2',
            type:'follow_request',
            username:'Hunk',
            avatarUrl:'/img/DefaultAvatar.png',
            timestamp:'16:30'
        },
        {
            id:'3',
            type:'followed',
            username:'JaneDoe',
            avatarUrl:'/img/DefaultAvatar.png',
            timestamp:'03:54'
        },
        {
            id:'4',
            type:'follow_request',
            username:'JohnSmith',
            avatarUrl:'/img/lion.png',
            timestamp:'16:50'
        },
        {
            id:'5',
            type:'followed',
            username:'Alice',
            avatarUrl:'/img/lion.png',
            timestamp:'12:20'
        },
        {
            id:'6',
            type:'follow_request',
            username:'Bob',
            avatarUrl:'/img/lion.png',
            timestamp:'10:15'
        },
        {
            id:'7',
            type:'followed',
            username:'Charlie',
            avatarUrl:'/img/lion.png',
            timestamp:'08:00'
        }
    ]
    const handleAccept = (notificationId) => {
        console.log(`Accepted notification with ID: ${notificationId}`);
    }
    const handleDecline = (notificationId) => {
        console.log(`Declined notification with ID: ${notificationId}`);
    }

    return (

        <div className='flex flex-col gap-2'>
            <PageTitle>Notifications</PageTitle>
            <div className='flex flex-col gap-2'>
                {notifications.map((notification) => (
                    <NotificationsCard
                        key={notification.id}
                        notification={notification}
                        onAccept={() => handleAccept(notification.id)}
                        onDecline={() => handleDecline(notification.id)}
                    />
                ))}
            </div>
        </div>
    );
}

export default NotificationScreen;