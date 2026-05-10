import React from 'react';

export default function Input({ type = 'text', as = 'input', className = '', ...props }) {
  if (as === 'textarea') {
    return <textarea className={`form-control ${className}`} {...props} />;
  }
  if (as === 'select') {
    return (
      <select className={`form-control ${className}`} {...props}>
        {props.children}
      </select>
    );
  }
  return <input type={type} className={`form-control ${className}`} {...props} />;
}
