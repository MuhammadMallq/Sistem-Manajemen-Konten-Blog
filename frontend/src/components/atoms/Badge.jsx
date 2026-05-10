import React from 'react';

export default function Badge({ children, type = 'category', className = '' }) {
  return (
    <span className={`badge badge-${type} ${className}`}>
      {children}
    </span>
  );
}
