import Image from 'next/image';
import React, { useState, useRef, useEffect } from 'react';

function SelectFollowers({ followers = [], onChange, value = [] }) {
  const [isOpen, setIsOpen] = useState(false);
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedFollowers, setSelectedFollowers] = useState(value);
  const dropdownRef = useRef(null);

  // Filtrer les followers selon le terme de recherche
  const filteredFollowers = followers.filter((follower) =>
    follower.username.toLowerCase().includes(searchTerm.toLowerCase())
  );

  // Gérer la sélection/désélection d'un follower
  const toggleFollower = (followerId) => {
    const updatedSelection = selectedFollowers.includes(followerId)
      ? selectedFollowers.filter((id) => id !== followerId)
      : [...selectedFollowers, followerId];

    setSelectedFollowers(updatedSelection);

    // Appeler onChange si fourni (pour react-hook-form)
    if (onChange) {
      onChange(updatedSelection);
    }
  };

  // Vérifier si un follower est sélectionné
  const isSelected = (followerId) => {
    return selectedFollowers.includes(followerId);
  };

  // Obtenir les noms des followers sélectionnés
  const getSelectedNames = () => {
    return followers
      .filter((follower) => selectedFollowers.includes(follower.user_id))
      .map((follower) => follower.username);
  };

  // Fermer le dropdown si on clique en dehors
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
        setIsOpen(false);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  // Réinitialiser la recherche quand on ferme
  const handleToggleDropdown = () => {
    if (isOpen) {
      setSearchTerm('');
    }
    setIsOpen(!isOpen);
  };

  return (
    <div className="relative" ref={dropdownRef}>
      {/* Select hidden pour la compatibilité form */}
      <select
        name="Selected-followers"
        id="Selected-followers"
        multiple
        value={selectedFollowers}
        onChange={() => {}} // Géré par les checkboxes
        className="hidden"
      >
        {selectedFollowers.map((user_id) => (
          <option key={user_id} value={user_id} selected></option>
        ))}
      </select>

      {/* Bouton principal du selecteur */}
      <button
        type="button"
        onClick={handleToggleDropdown}
        className="relative w-full py-3 px-4 pr-10 flex items-center justify-between bg-white border border-gray-200 rounded-lg text-left text-sm focus:outline-none focus:ring-2 focus:ring-lavender-4 focus:border-lavender-4 cursor-pointer"
      >
        <span className="block truncate">
          {selectedFollowers.length === 0
            ? 'Select your followers...'
            : `${selectedFollowers.length} follower${
                selectedFollowers.length > 1 ? 's' : ''
              } selected`}
        </span>
        <div className="absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none">
          <svg
            className={`w-4 h-4 text-gray-400 transition-transform ${
              isOpen ? 'rotate-180' : ''
            }`}
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M19 9l-7 7-7-7"
            />
          </svg>
        </div>
      </button>

      {/* Liste des followers sélectionnés (en dehors du dropdown) */}
      {selectedFollowers.length > 0 && (
        <div className="mt-2 flex flex-wrap gap-2">
          {getSelectedNames().map((name, index) => (
            <span
              key={index}
              className="inline-flex items-center px-2 py-1 bg-lavender-1 text-white text-xs rounded-full"
            >
              <Image
                src={
                  followers.find((f) => f.username === name)?.image_url ||
                  '/img/DefaultAvatar.svg'
                }
                alt={name + ' avatar'}
                width={24}
                height={24}
                className="rounded-full inline-block w-6 h-6 mr-1"
              />
              {name}
            </span>
          ))}
        </div>
      )}

      {/* Dropdown */}
      {isOpen && (
        <div className="absolute z-10 mt-1 w-full bg-white border border-gray-200 rounded-lg shadow-lg">
          {/* Barre de recherche */}
          <div className="p-3 border-b border-gray-200">
            <div className="relative">
              <input
                type="text"
                placeholder="Search followers..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
                className="w-full py-2 px-3 pr-10 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-lavender-4 focus:border-lavender-4 text-sm"
              />
              <div className="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
                <svg
                  className="w-4 h-4 text-gray-400"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                  />
                </svg>
              </div>
            </div>
          </div>

          {/* Liste des followers */}
          <div className="max-h-48 overflow-y-auto">
            {filteredFollowers.length === 0 ? (
              <div className="px-3 py-2 text-dark-grey italic text-sm">
                {searchTerm ? 'No followers found' : 'No followers available'}
              </div>
            ) : (
              filteredFollowers.map((follower) => (
                <div
                  key={follower.user_id}
                  onClick={() => toggleFollower(follower.user_id)}
                  className={`flex items-center justify-between px-3 py-2 cursor-pointer transition-colors ${
                    isSelected(follower.user_id) ? 'bg-lavender-1' : ''
                  }`}
                >
                  <div className="flex items-center gap-3">
                    <input
                      type="checkbox"
                      checked={isSelected(follower.user_id)}
                      onChange={() => {}} // Géré par le onClick du parent
                      className="hidden w-4 h-4 text-lavender-5 border-gray-300 rounded focus:ring-lavender-4 focus:ring-2"
                    />
                    <span
                      className={`text-sm ${
                        isSelected(follower.user_id)
                          ? 'font-medium text-white'
                          : 'text-gray-700'
                      }`}
                    >
                      <Image
                        src={follower.image_url || '/img/DefaultAvatar.svg'}
                        alt={follower.username + ' avatar'}
                        width={24}
                        height={24}
                        className="rounded-full inline-block w-6 h-6 mr-2"
                      />
                      {follower.username}
                    </span>
                  </div>
                  {isSelected(follower.user_id) && (
                    <svg
                      className="w-4 h-4 text-lavender-5"
                      fill="currentColor"
                      viewBox="0 0 20 20"
                    >
                      <path
                        fillRule="evenodd"
                        d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                        clipRule="evenodd"
                      />
                    </svg>
                  )}
                </div>
              ))
            )}
          </div>
        </div>
      )}
    </div>
  );
}

export default SelectFollowers;
