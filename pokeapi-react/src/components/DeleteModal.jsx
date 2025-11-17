import React from 'react';

function DeleteModal({ name, isOpen, onClose, onConfirm }) {
    if (!isOpen) return null;

    return (
        <div className="modal-backdrop">
            <div className="modal-content pokedex-style-modal">
                <h3 className="modal-title">Confirmar Exclusão</h3>
                <p className="modal-text">
                    Tem certeza que deseja DELETAR o Pokémon {name}?
                    Esta ação é irreversível.
                </p>
                <div className="modal-actions">
                    <button className="btn remove-btn" onClick={onConfirm}>
                        DELETAR
                    </button>
                    <button className="btn submit-btn" onClick={onClose}>
                        CANCELAR
                    </button>
                </div>
            </div>
        </div>
    );
}

export default DeleteModal;