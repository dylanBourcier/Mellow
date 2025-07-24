'use client';

import React from 'react';
import UserInfo from './UserInfo';
import Button from './Button';
import { useState } from 'react';
import { formatDate } from '@/app/utils/date';
import toast from 'react-hot-toast';
import { useUser } from '@/app/context/UserContext';
import { useEffect } from 'react';
import {icons} from '@/app/lib/icons'; // Assurez-vous que le chemin est correct

function Event({ event }) {
  const [vote, setVote] = useState(null);
  const { user } = useUser();
  const [hasVoted, setHasVoted] = useState(false);
  const [yesCount, setYesCount] = useState(0);
  const [noCount, setNoCount] = useState(0);
  const [userVote, setUserVote] = useState(null);

  // Vérifier si l'utilisateur a déjà voté
  useEffect(() => {
    if (!event || !event.event_responses) return;

    const userResponse = event.event_responses.find(
      (r) => r.user_id === user?.user_id
    );
    if (userResponse) {
      setHasVoted(true);
      setUserVote(userResponse.vote);
    }

    const yesVotes = event.event_responses.filter(
      (r) => r.vote === true
    ).length;
    const noVotes = event.event_responses.filter(
      (r) => r.vote === false
    ).length;

    setYesCount(yesVotes);
    setNoCount(noVotes);
  }, [event.event_responses, user?.user_id]);

  const { creation_date, event_date, title, username, avatar_url } = event;
  const getButtonStyle = (type) =>
    `flex flex-1 hover:cursor-pointer px-4 py-1 rounded-xl font-light border transition-all duration-200 ${
      vote === type
        ? 'bg-lavender-1 text-white border-lavender-1'
        : 'text-lavender-5 border-lavender-5'
    }`;

  const handleVoteSubmit = () => {
    if (!vote) {
      alert('Please select Yes or No before voting.');
      return;
    }

    // 👉 Remplace ceci par ton appel API ensuite
    //Fetch API pour soumettre le vote
    const formData = new URLSearchParams();
    formData.append('vote', vote);
    fetch(`/api/groups/events/vote/${event.event_id}`, {
      method: 'POST',
      body: formData.toString(),
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },

      credentials: 'include', // pour envoyer les cookies de session
    })
      .then((res) => res.json())
      .then((data) => {
        if (data.status === 'success') {
          toast.success('Vote submitted successfully!');
        } else {
          toast.error(data.message || 'Failed to submit vote.');
        }
      })
      .catch((err) => {
        toast.error(
          err.message || 'An error occurred while submitting your vote.'
        );
      });
    setVote(null); // Réinitialiser le vote après soumission
    setHasVoted(true); // Marquer l'utilisateur comme ayant voté
    setUserVote(vote); // Enregistrer le vote de l'utilisateur
    // Mettre à jour les compteurs de votes
    if (vote === 'yes') {
      setYesCount((prev) => prev + 1);
    } else {
      setNoCount((prev) => prev + 1);
    }

    // Optionnel : désactiver les boutons ou donner un retour visuel
  };
  return (
    <div className="flex flex-col w-full bg-white shadow-(--box-shadow) rounded-lg hover:-translate-y-0.5 transition-all duration-300">
      <div className="flex  w-full p-3 gap-3 items-center justify-between">
        <UserInfo userName={username} authorAvatar={avatar_url} />
      </div>
      <div className="flex flex-col w-full px-3 mb-2 items-start justify-between">
        <h4 className="text-dark-grey p-0 text-xl font-heading leading-tight">
          {title}{' '}
        </h4>
        <span className=" flex justify-end gap-2 text-dark-grey w-full text-right items-center text-sm font-semibold ">
          {icons['events']} {formatDate(event_date)}
        </span>{' '}
      </div>
      {hasVoted && (
        <div className="flex w-full px-3 gap-3 items-start justify-between mb-4">
          <div
            className={`flex-1 border border-lavender-5 rounded-xl px-4 py-1 ${
              userVote === true
                ? 'font-medium bg-lavender-1 text-white'
                : 'text-lavender-5 font-light'
            }`}
          >
            Yes ({yesCount} vote{yesCount > 1 ? 's' : ''})
          </div>
          <div
            className={`flex-1 border border-lavender-5 rounded-xl px-4 py-1 ${
              userVote === false
                ? 'font-medium bg-lavender-1 text-white'
                : 'text-lavender-5 font-light'
            }`}
          >
            No ({noCount} vote{noCount > 1 ? 's' : ''})
          </div>
        </div>
      )}
      {!hasVoted && (
        <>
          {' '}
          <div className="flex w-full items-center gap-2 px-3 mt-1 mb-1.5">
            <button
              onClick={() => setVote('yes')}
              className={getButtonStyle('yes')}
            >
              Yes
            </button>
            <button
              onClick={() => setVote('no')}
              className={getButtonStyle('no')}
            >
              No
            </button>
          </div>
          <div className="flex w-full items-center gap-2 px-3 mb-3">
            <Button
              className="w-full"
              onClick={handleVoteSubmit}
              disabled={!vote}
            >
              Vote
            </Button>
          </div>
        </>
      )}
    </div>
  );
}

export default Event;
