import React from 'react';

export default function TableTemplate({ headers, children, isEmpty, emptyMessage = 'Tidak ada data' }) {
  return (
    <div className="table-container">
      <table>
        <thead>
          <tr>
            {headers.map((header, index) => (
              <th key={index}>{header}</th>
            ))}
          </tr>
        </thead>
        <tbody>
          {isEmpty ? (
            <tr>
              <td colSpan={headers.length}>
                <div className="empty-state">
                  <h3>{emptyMessage}</h3>
                </div>
              </td>
            </tr>
          ) : (
            children
          )}
        </tbody>
      </table>
    </div>
  );
}
