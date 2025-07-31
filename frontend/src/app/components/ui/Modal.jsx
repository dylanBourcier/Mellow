import React from "react";

// Composant Modal réutilisable
// Affiche une fenêtre modale avec un titre, un message et des boutons d'action dynamiques
const Modal = ({ isOpen, onClose, title, message, actions = [] }) => {
  // Si le modal n'est pas ouvert, ne rien afficher
  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black/50 z-50">
      {/* Contenu principal du modal */}
      <div className="bg-white rounded-2xl p-8 min-w-[320px]">
        {/* Titre du modal */}
        <h2 className="font-quicksand text-xl font-semibold text-lavender-3 m-0">
          {title}
        </h2>
        {/* Message du modal */}
        <p className="font-inter text-base text-dark-blue mt-4 mb-0">
          {message}
        </p>
        {/* Zone des boutons d'action */}
        <div className="mt-8 flex justify-end gap-4">
          {actions.length > 0 ? (
            actions.map((action, idx) => (
              <button
                key={idx}
                onClick={() => {
                  action.onClick?.();
                  if (action.closeOnClick !== false) onClose();
                }}
                disabled={action.disabled}
                className={`
                  bg-lavender-3 text-white px-6 py-2 rounded-xl border-none font-inter font-medium text-base
                  transition-colors duration-200
                  ${action.disabled ? "opacity-60 cursor-not-allowed" : "hover:bg-lavender-5 cursor-pointer"}
                `}
                type="button"
              >
                {action.label}
              </button>
            ))
          ) : (
            // Sinon, afficher un bouton "Fermer" par défaut
            <button
              onClick={onClose}
              className="bg-lavender-3 text-white px-6 py-2 rounded-xl border-none font-inter font-medium text-base transition-colors duration-200 hover:bg-lavender-5 cursor-pointer"
              type="button"
            >
              Fermer
            </button>
          )}
        </div>
      </div>
    </div>
  );
};

export default Modal;