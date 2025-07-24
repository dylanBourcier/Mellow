'use client';
import React from 'react';
import { useEffect, useState } from 'react';
import Event from '../ui/Event';
import CustomToast from '../ui/CustomToast';
import toast from 'react-hot-toast';
import Button from '../ui/Button';
import Spinner from '../ui/Spinner';

export default function EventsScreen({ groupId }) {
  //Fetch events for the group
  // EventID      uuid.UUID `json:"event_id"`
  // UserID       uuid.UUID `json:"user_id"`
  // GroupID      uuid.UUID `json:"group_id"`
  // CreationDate time.Time `json:"creation_date"`
  // EventDate    time.Time `json:"event_date"`
  // Title        string    `json:"title"`
  const [events, setEvents] = useState([]);
  const [loading, setLoading] = useState(true);
  const [isMember, setIsMember] = useState(false);

  useEffect(() => {
    if (!groupId) return;

    fetch(`/api/groups/events/${groupId}`, { credentials: 'include' }) // adapte l'URL selon ton backend
      .then((res) => {
        console.log(res);

        // if (!res.ok) throw new Error('Failed to fetch group posts');
        return res.json();
      })
      .then((data) => {
        console.log(data);
        if (data.status == 'error') {
          throw new Error(data.message || 'Failed to fetch group posts');
        }
        if (data.message === 'Not member') {
          setLoading(false);
          setIsMember(false);
          return;
        }
        setIsMember(true);
        setEvents(data.data || []);
        console.log('Events fetched successfully:', data.data);

        setLoading(false);
      })
      .catch((err) => {
        setLoading(false);
        toast.custom((t) => (
          <CustomToast
            t={t}
            type="error"
            message={err.message || 'An error occurred while fetching events.'}
          />
        ));
      });
  }, [groupId]);

  if (loading)
    return (
      <div className="flex flex-col gap-3 w-full items-center">
        <span className="flex gap-2">
          <Spinner></Spinner>Loading events...
        </span>
      </div>
    );
  if (!loading)
    return (
      <div className="flex flex-col gap-3 w-full items-center">
        {isMember ? (
          <>
            <Button
              wFull={true}
              className="w-full"
              href={`/groups/${groupId}/events/create`}
            >
              Create Event
            </Button>
            {events
              .sort(
                (a, b) => new Date(b.creation_date) - new Date(a.creation_date)
              )
              .map((evt) => (
                <Event key={evt.event_id} event={evt} />
              ))}
          </>
        ) : (
          <span className="text-dark-grey-lighter">
            You are not a member of this group. You can't see the events.
          </span>
        )}
      </div>
    );
}
