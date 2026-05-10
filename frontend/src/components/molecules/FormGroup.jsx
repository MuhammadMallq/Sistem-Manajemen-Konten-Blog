import React from 'react';
import Input from '../atoms/Input';

export default function FormGroup({ label, ...inputProps }) {
  return (
    <div className="form-group">
      {label && <label>{label}</label>}
      <Input {...inputProps} />
    </div>
  );
}
