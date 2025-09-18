'use client';

import { useEffect, useState } from 'react';
import Spinner from '../ui/Spinner';
import CustomToast from '../ui/CustomToast';
import { toast } from 'react-hot-toast';

export default function GroupRequestsPanel({ groupId }) {
  const [loading, setLoading] = useState(true);
  const [items, setItems] = useState([]);
  const [busy, setBusy] = useState(false);

  const fetchItems = async () => {
    setLoading(true);
    try {
      const res = await fetch(`/api/groups/${groupId}/join-requests`, {
        credentials: 'include',
      });
      const data = await res.json();
      if (res.ok && data.data) {
        setItems(data.data);
      } else {
        throw new Error(data.message || 'Failed to load requests');
      }
    } catch (e) {
      toast.custom((t) => (
        <CustomToast t={t} type="error" message={e.message} />
      ));
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchItems();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [groupId]);

  const act = async (requestId, action) => {
    setBusy(true);
    try {
      const res = await fetch(
        `/api/groups/${groupId}/join-requests/${requestId}/${action}`,
        { method: 'POST', credentials: 'include' }
      );
      const data = await res.json();
      if (!res.ok) throw new Error(data.message || 'Action failed');
      setItems((prev) => prev.filter((it) => it.id !== requestId));
      toast.success(`Request ${action}ed`);
    } catch (e) {
      toast.custom((t) => (
        <CustomToast t={t} type="error" message={e.message} />
      ));
    } finally {
      setBusy(false);
    }
  };

  if (loading) return (<div className="flex items-center gap-2"><Spinner/>Loading...</div>);

  return (
    <div className="flex flex-col gap-3">
      <h2 className="text-lg font-semibold">Join Requests</h2>
      {items.length === 0 ? (
        <div className="text-dark-grey-lighter">No pending requests.</div>
      ) : (
        items.map((it) => (
          <div key={it.id} className="flex items-center justify-between p-2 border rounded-xl bg-white">
            <div className="flex items-center gap-2">
              {it.requester?.avatar ? (
                // eslint-disable-next-line @next/next/no-img-element
                <img src={it.requester.avatar} alt="avatar" className="w-8 h-8 rounded-full" />
              ) : (
                <div className="w-8 h-8 rounded-full bg-light-grey" />
              )}
              <div className="flex flex-col">
                <span className="font-medium">{it.requester?.username || it.requester?.id}</span>
                <span className="text-xs text-dark-grey-lighter">{new Date(it.createdAt).toLocaleString()}</span>
              </div>
            </div>
            <div className="flex gap-2">
              <button disabled={busy} onClick={() => act(it.id, 'accept')} className="px-2 py-1 rounded-lg bg-lavender-1 text-white">Accept</button>
              <button disabled={busy} onClick={() => act(it.id, 'reject')} className="px-2 py-1 rounded-lg border border-dark-grey">Reject</button>
            </div>
          </div>
        ))
      )}
    </div>
  );
}

