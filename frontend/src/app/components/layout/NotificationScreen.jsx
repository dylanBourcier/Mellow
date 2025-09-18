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
          const acceptRequest = async () => {
            const res = await fetch(
              `/api/users/follow/request/${notification.request_id}?action=accept`,
              {
                method: 'POST',
                credentials: 'include',
              }
            );
            const data = await res.json();
            if (data.status !== 'success') {
              throw new Error(
                data.message || 'Failed to accept follow request'
              );
            }
          };

          const processAccept = async () => {
            try {
              await markAsRead(notification);
              await acceptRequest();
              // Optionally, you can update the notification to reflect the accepted state
              setNotifications((prev) =>
                prev.map((n) =>
                  n.notification_id === notification.notification_id
                    ? { ...n, seen: true }
                    : n
                )
              );
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
          };

          processAccept();
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
      case 'group_request':
        // Accept a group join request from notification
        try {
          const acceptGroupRequest = async () => {
            const res = await fetch(
              `/api/groups/${notification.group_id}/join-requests/${notification.request_id}/accept`,
              {
                method: 'POST',
                credentials: 'include',
              }
            );
            const data = await res.json();
            if (data.status !== 'success') {
              throw new Error(data.message || 'Failed to accept group request');
            }
          };
          const process = async () => {
            try {
              await markAsRead(notification);
              await acceptGroupRequest();
              setNotifications((prev) =>
                prev.map((n) =>
                  n.notification_id === notification.notification_id
                    ? { ...n, seen: true }
                    : n
                )
              );
            } catch (error) {
              console.error('Error accepting group request:', error);
              toast.custom((t) => (
                <CustomToast
                  t={t}
                  type="error"
                  message={'Error accepting group request! ' + error.message}
                />
              ));
            }
          };
          process();
        } catch (error) {
          console.error('Error accepting group request:', error);
          toast.custom((t) => (
            <CustomToast
              t={t}
              type="error"
              message={'Error accepting group request! ' + error.message}
            />
          ));
        }
        break;
      case 'group_invite':
        // Logic to accept group invite
        try {
          const acceptGroupInvite = async () => {
            const res = await fetch(
              `/api/groups/invite/answer/${notification.request_id}?action=accept`,
              {
                method: 'POST',
                credentials: 'include',
              }
            );
            const data = await res.json();
            if (data.status !== 'success') {
              throw new Error(data.message || 'Failed to accept group invite');
            }
          };
          const processGroupInvite = async () => {
            try {
              await markAsRead(notification);
              await acceptGroupInvite();
              // Optionally, you can update the notification to reflect the accepted state
              setNotifications((prev) =>
                prev.map((n) =>
                  n.notification_id === notification.notification_id
                    ? { ...n, seen: true }
                    : n
                )
              );
            } catch (error) {
              console.error('Error accepting group invite:', error);
              toast.custom((t) => (
                <CustomToast
                  t={t}
                  type="error"
                  message={'Error accepting group invite ! ' + error.message}
                />
              ));
            }
          };
          processGroupInvite();
        } catch (error) {
          console.error('Error accepting group invite:', error);
          toast.custom((t) => (
            <CustomToast
              t={t}
              type="error"
              message={'Error accepting group invite! ' + error.message}
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
              await markAsRead(notification);
              await declineRequest();
              // Update the notifications state or refetch notifications
              setNotifications((prev) =>
                prev.map((n) =>
                  n.notification_id === notification.notification_id
                    ? { ...n, seen: true }
                    : n
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
      case 'group_request':
        try {
          const declineGroupRequest = async () => {
            const res = await fetch(
              `/api/groups/${notification.group_id}/join-requests/${notification.request_id}/reject`,
              {
                method: 'POST',
                credentials: 'include',
              }
            );
            const data = await res.json();
            if (data.status !== 'success') {
              throw new Error(data.message || 'Failed to reject group request');
            }
          };
          const process = async () => {
            try {
              await markAsRead(notification);
              await declineGroupRequest();
              setNotifications((prev) =>
                prev.map((n) =>
                  n.notification_id === notification.notification_id
                    ? { ...n, seen: true }
                    : n
                )
              );
            } catch (error) {
              console.error('Error rejecting group request:', error);
              toast.custom((t) => (
                <CustomToast
                  t={t}
                  type="error"
                  message={'Error rejecting group request! ' + error.message}
                />
              ));
            }
          };
          process();
        } catch (error) {
          console.error('Error rejecting group request:', error);
          toast.custom((t) => (
            <CustomToast
              t={t}
              type="error"
              message={'Error rejecting group request! ' + error.message}
            />
          ));
        }
        break;
      case 'group_invite':
        // Logic to decline group invite
        try {
          const declineGroupInvite = async () => {
            const res = await fetch(
              `/api/groups/invite/answer/${notification.request_id}?action=reject`,
              {
                method: 'POST',
                credentials: 'include',
              }
            );
            const data = await res.json();
            if (data.status !== 'success') {
              throw new Error(data.message || 'Failed to decline group invite');
            }
          };
          const processGroupInvite = async () => {
            try {
              await markAsRead(notification);
              await declineGroupInvite();
              // Update the notifications state or refetch notifications
              setNotifications((prev) =>
                prev.map((n) =>
                  n.notification_id === notification.notification_id
                    ? { ...n, seen: true }
                    : n
                )
              );
            } catch (error) {
              console.error('Error declining group invite:', error);
              toast.custom((t) => (
                <CustomToast
                  t={t}
                  type="error"
                  message={'Error declining group invite! ' + error.message}
                />
              ));
            }
          };
          processGroupInvite();
        } catch (error) {
          console.error('Error declining group invite:', error);
          toast.custom((t) => (
            <CustomToast
              t={t}
              type="error"
              message={'Error declining group invite! ' + error.message}
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
// Mark the notification as read and decline the request sequentially
const markAsRead = async (notification) => {
  const res = await fetch(
    `/api/notifications/read/${notification.notification_id}`,
    {
      method: 'PATCH',
      credentials: 'include',
    }
  );
  const data = await res.json();
  if (data.status !== 'success') {
    throw new Error(data.message || 'Failed to mark notification as read');
  }
};

export default NotificationScreen;
