import React from 'react';

export default function StatCard({ title, value, icon: Icon, index = 0 }) {
  return (
    <div className="stat-card">
      <div className="stat-icon">
        {Icon && <Icon />}
      </div>
      <div className="stat-value">{value}</div>
      <div className="stat-label">{title}</div>
    </div>
  );
}
