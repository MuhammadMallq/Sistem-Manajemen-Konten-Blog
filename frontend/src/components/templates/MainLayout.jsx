import React from 'react';
import { Toaster } from 'react-hot-toast';
import Sidebar from '../organisms/Sidebar';

export default function MainLayout({ children }) {
  return (
    <>
      <Toaster position="top-right" toastOptions={{
        style: { background: '#1a1a2e', color: '#e8e8f0', border: '1px solid #2a2a4a' }
      }} />
      <div className="app-layout">
        <Sidebar />
        <main className="main-content">
          {children}
        </main>
      </div>
    </>
  );
}
