'use client';

import React, { use } from 'react';
import PageTitle from '../ui/PageTitle';
import NotificationsCard from '../ui/NotificationsCard';
import { useState, useEffect } from 'react';
import CustomToast from '../ui/CustomToast';
import { toast } from 'react-hot-toast';

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
          console.log(data);

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
  const handleAccept = (notification) => {
    switch (notification.type) {
      case 'follow_request':
        // Logic to accept follow request
        try {
          //Need to mark the notification as read and accept the request, so two fetches Promise.all
          Promise.all([
            fetch(`/api/notifications/read/${notification.notification_id}`, {
              method: 'PATCH',
              credentials: 'include',
            }),
            fetch(
              `/api/users/follow/request/${notification.request_id}?action=accept`,
              {
                method: 'POST',
                credentials: 'include',
              }
            ),
          ]).then((responses) => {
            if (!responses[0].ok || !responses[1].ok) {
              throw new Error('Failed to accept follow request');
            }
            // Update the notifications state or refetch notifications
            setNotifications((prev) =>
              prev.filter(
                (n) => n.notification_id !== notification.notification_id
              )
            );
          });
        } catch (error) {
          console.error('Error accepting follow request:', error);
          toast.custom((t) => (
            <CustomToast
              t={t}
              type="error"
              message={'Error accepting follow request! ' + error.message}
            />
          ));
        }
        break;
      default:
        console.log(
          `Accepted notification with ID: ${notification.notification_id}`
        );
    }
  };
  const handleDecline = (notification) => {
    switch (notification.type) {
      case 'follow_request':
        // Logic to decline follow request
        try {
          // Mark the notification as read and decline the request sequentially
          const markAsRead = async () => {
            const res = await fetch(
              `/api/notifications/read/${notification.notification_id}`,
              {
                method: 'PATCH',
                credentials: 'include',
              }
            );
            const data = await res.json();
            if (data.status !== 'success') {
              throw new Error(
                data.message || 'Failed to mark notification as read'
              );
            }
          };

          const declineRequest = async () => {
            const res = await fetch(
              `/api/users/follow/request/${notification.request_id}?action=reject`,
              {
                method: 'POST',
                credentials: 'include',
              }
            );
            const data = await res.json();
            if (data.status !== 'success') {
              throw new Error(
                data.message || 'Failed to decline follow request'
              );
            }
          };

          const processDecline = async () => {
            try {
              await markAsRead();
              await declineRequest();
              // Update the notifications state or refetch notifications
              setNotifications((prev) =>
                prev.filter(
                  (n) => n.notification_id !== notification.notification_id
                )
              );
            } catch (error) {
              console.error('Error declining follow request:', error);
              toast.custom((t) => (
                <CustomToast
                  t={t}
                  type="error"
                  message={'Error declining follow request! ' + error.message}
                />
              ));
            }
          };

          processDecline();
        } catch (error) {
          console.error('Error declining follow request:', error);
          toast.custom((t) => (
            <CustomToast
              t={t}
              type="error"
              message={'Error declining follow request! ' + error.message}
            />
          ));
        }
        break;
      default:
        console.log(
          `Declined notification with ID: ${notification.notification_id}`
        );
    }
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
            key={notification.notification_id}
            notification={notification}
            onAccept={() => handleAccept(notification)}
            onDecline={() => handleDecline(notification)}
          />
        ))}
      </div>
    </div>
  );
}

export default NotificationScreen;
