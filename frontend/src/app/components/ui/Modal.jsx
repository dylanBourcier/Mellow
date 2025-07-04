import React from "react";

const Modal = ({ isOpen, onClose, message}) => {
if (!isOpen) return null;

return (
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
        <div className="bg-white rounded-lg p-6 shadow-lg">
        <h2 className="text-lg font-semibold">Confirmation</h2>
        <p className="mt-4">{message}</p>
            <div className="mt-6 flex justify-end">
                <button
                className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
                onClick={onClose}>
                    Fermer
                </button>
            </div>
        </div>
    </div>
);
};

export default Modal