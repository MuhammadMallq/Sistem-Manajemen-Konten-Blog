import React from 'react';

export default function Button({ children, variant = 'primary', icon: Icon, className = '', ...props }) {
  const baseClass = variant === 'icon' ? 'btn-icon' : `btn btn-${variant}`;
  return (
    <button className={`${baseClass} ${className}`} {...props}>
      {Icon && <Icon />}
      {children}
    </button>
  );
}
