import React from 'react';


function folows() {
  
    const handleFollow = async (e) => {
        e.preventDefault();
        // Implement follow logic here
        //change the text in the button to "Unfollow" after following
        const button = e.target;
        if (button.innerText === 'Follow') {
            button.innerText = 'Unfollow';
            // Call API to follow the user
            try {
                const response = await fetch('/api/follow', {
                    method: 'POST',
                    credentials: 'include',
                });
                const data = await response.json();
                if (data.status !== 'success') {
                    throw new Error(data.message || 'Failed to follow user');
                }
            } catch (error) {
                console.error('Error following user:', error);
            }
        } else {
            button.innerText = 'Follow';
            // Call API to unfollow the user
            try {
                const response = await fetch('/api/unfollow', {
                    method: 'POST',
                    credentials: 'include',
                });
                const data = await response.json();
                if (data.status !== 'success') {
                    throw new Error(data.message || 'Failed to unfollow user');
                }
            } catch (error) {
                console.error('Error unfollowing user:', error);
            }
        }

    }
    return (
        <div>
            <button type="submit" htmlFor="follow" onClick={handleFollow} className='px-4 py-2 gap-2.5 text-white border-lavender-3 bg-lavender-3 rounded-xl cursor-pointer hover:hover:bg-lavender-5 shadow-(--box-shadow)'>Follow</button>
        </div>
    );
}

export default folows;