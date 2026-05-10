import React from 'react';
import Modal from './Modal';
import Button from '../atoms/Button';

export default function ConfirmDialog({ isOpen, onClose, onConfirm, title, message }) {
  return (
    <Modal title={title} isOpen={isOpen} onClose={onClose}>
      <div className="confirm-dialog">
        <p dangerouslySetInnerHTML={{ __html: message }}></p>
        <div className="modal-actions">
          <Button variant="secondary" onClick={onClose}>Batal</Button>
          <Button variant="danger" onClick={onConfirm}>Hapus</Button>
        </div>
      </div>
    </Modal>
  );
}
