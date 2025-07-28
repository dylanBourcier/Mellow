import React from "react";

const Modal = ({ isOpen, onClose,title, message }) => {
  if (!isOpen) return null;

  return (
    <div
      className="fixed inset-0 flex items-center justify-center"
      style={{
        backgroundColor: "rgba(0,0,0,0.5)",
        zIndex: 50,
      }}
    >
      <div
        style={{
          background: "var(--color-background, #fff)",
          borderRadius: "1rem",
          padding: "2rem",
          minWidth: "320px",
        }}
      >
        <h2
          style={{
            fontFamily: "var(--font-quicksand), sans-serif",
            fontSize: "1.25rem",
            fontWeight: 600,
            color: "var(--color-primary, #6C63FF)",
            margin: 0,
          }}
        >
          {title}
        </h2>
        <p
          style={{
            fontFamily: "var(--font-inter), sans-serif",
            fontSize: "1rem",
            color: "var(--color-text, #22223B)",
            marginTop: "1rem",
            marginBottom: 0,
          }}
        >
          {message}
        </p>
        <div style={{ marginTop: "2rem", display: "flex", justifyContent: "flex-end" }}>
          <button
            onClick={onClose}
            style={{
              background: "var(--color-primary, #6C63FF)",
              color: "var(--color-background, #fff)",
              padding: "0.5rem 1.5rem",
              borderRadius: "0.75rem",
              border: "none",
              fontFamily: "var(--font-inter), sans-serif",
              fontWeight: 500,
              fontSize: "1rem",
              cursor: "pointer",
              transition: "background 0.2s",
            }}
            onMouseOver={e => (e.currentTarget.style.background = "var(--color-primary-hover, #574fd6)")}
            onMouseOut={e => (e.currentTarget.style.background = "var(--color-primary, #6C63FF)")}
          >
            Fermer
          </button>
        </div>
      </div>
    </div>
  );
};

export default Modal;