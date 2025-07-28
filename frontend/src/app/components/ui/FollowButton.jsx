import React, { useState } from 'react';

function FollowButton({ followStatus: initialStatus, targetID }) {
  const [status, setStatus] = useState(initialStatus);
  const [loading, setLoading] = useState(false);

  const handleFollow = async (e) => {
    e.preventDefault();

    if (status === 'requested') return; // Pas d'action sur "pending"

    setLoading(true);

    try {
      if (status === 'follows') {
        const res = await fetch(`/api/users/unfollow/${targetID}`, {
          method: 'POST',
          credentials: 'include',
        });
        const data = await res.json();
        if (data.status === 'success') {
          setStatus('not_follow');
        } else {
          throw new Error(data.message || 'Failed to unfollow user');
        }
      } else if (status === 'not_follow') {
        const res = await fetch(`/api/users/follow/${targetID}`, {
          method: 'POST',
          credentials: 'include',
        });
        const data = await res.json();
        if (data.status === 'success') {
          setStatus('follows'); // ou 'requested' si tu veux une validation
        } else {
          throw new Error(data.message || 'Failed to follow user');
        }
      }
    } catch (err) {
      console.error('Follow/unfollow error:', err);
      alert('An error occurred.');
    } finally {
      setLoading(false);
    }
  };

  const renderButtonText = () => {
    if (loading) return 'Loading...';
    if (status === 'follows') return 'Unfollow';
    if (status === 'not_follow') return 'Follow';
    if (status === 'requested') return 'Pending';
    return 'Follow';
  };

  return (
    <button
      type="button"
      onClick={handleFollow}
      className="px-4 py-2 gap-2.5 text-white border-lavender-3 bg-lavender-3 rounded-xl cursor-pointer hover:bg-lavender-5 shadow-(--box-shadow)"
      disabled={loading || status === 'requested'}
    >
      {renderButtonText()}
    </button>
  );
}

export default FollowButton;
