import React, { useState } from 'react';

function FollowButton({ followStatus: initialStatus, targetID, privacy }) {
  const [status, setStatus] = useState(initialStatus);
  const [loading, setLoading] = useState(false);
  const [isHovered, setIsHovered] = useState(false);

  const handleFollow = async (e) => {
    e.preventDefault();

    if (status === 'requested') return;
    if (status === 'yourself') return;

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
          if (privacy === 'private') {
            setStatus('requested');
          } else {
            setStatus('follows');
          }

          console.log('aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa', data.status);
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
    if (status === 'follows') return isHovered ? 'Unfollow' : 'Following';
    if (status === 'not_follow') return 'Follow';
    if (status === 'requested') return 'Pending';
    if (status === 'yourself') return 'You';
    return 'Follow';
  };

  const statusClassNames = {
    follows:
      'bg-transparent text-lavender-5 border border-lavender-5 hover:bg-red-100 hover:text-red-500 hover:border-red-400 cursor-pointer',
    not_follow: 'bg-lavender-3 hover:bg-lavender-5 text-white cursor-pointer',
    requested:
      'bg-transparent text-dark-gray cursor-not-allowed border border-dark-gray',
    yourself: 'bg-gray-500 cursor-not-allowed',
  };

  return (
    <button
      type="button"
      onClick={handleFollow}
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
      className={`px-4 py-2 gap-2.5 rounded-xl shadow-(--box-shadow) ${
        statusClassNames[status] || ''
      } ${status === 'yourself' ? 'hidden' : ''}`}
      disabled={loading || status === 'requested' || status === 'yourself'}
    >
      {renderButtonText()}
    </button>
  );
}

export default FollowButton;
