'use client';

import React, { use } from 'react';
import PageTitle from '../ui/PageTitle';
import NotificationsCard from '../ui/NotificationsCard';
import { useState, useEffect } from 'react';

function NotificationScreen() {
  const [notifications, setNotifications] = useState([]);
  useEffect(() => {
    // Fetch notifications from the API
    const fetchNotifications = async () => {
      try {
        const res = await fetch('/api/notifications', {
          credentials: 'include',
        });
        const data = await res.json();
        if (data.status !== 'success') {
          throw new Error(data.message || 'Failed to fetch notifications');
        }
        setNotifications(data.data);
        // Handle notifications data
      } catch (err) {
        toast.custom((t) => (
          <CustomToast
            t={t}
            type="error"
            message={'Error fetching notifications! ' + err.message}
          />
        ));
      }
    };

    fetchNotifications();
  }, []);
  const handleAccept = (notificationId) => {
    console.log(`Accepted notification with ID: ${notificationId}`);
  };
  const handleDecline = (notificationId) => {
    console.log(`Declined notification with ID: ${notificationId}`);
  };

  // Sort notifications by date (assuming notifications have a 'date' property)

  if (!notifications || notifications.length === 0) {
    return (
      <div className="flex items-center justify-center h-full">
        <span className="text-gray-500">No notifications available</span>
      </div>
    );
  }
  const sortedNotifications = [...notifications].sort(
    (a, b) => new Date(b.creation_date) - new Date(a.creation_date)
  );

  return (
    <div className="flex flex-col gap-2">
      <PageTitle>Notifications</PageTitle>
      <div className="flex flex-col gap-2">
        {sortedNotifications.map((notification) => (
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
