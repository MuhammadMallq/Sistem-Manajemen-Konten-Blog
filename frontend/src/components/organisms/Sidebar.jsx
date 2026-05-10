import React from 'react';
import { NavLink } from 'react-router-dom';
import { HiOutlineHome, HiOutlineDocumentText, HiOutlineUsers, HiOutlineTag, HiOutlineChatBubbleLeftRight } from 'react-icons/hi2';

export default function Sidebar() {
  return (
    <aside className="sidebar">
      <div className="sidebar-logo">
        <div className="logo-icon">B</div>
        <h1>BlogCMS</h1>
      </div>
      <nav>
        <div className="nav-section">
          <div className="nav-section-title">Menu Utama</div>
          <NavLink to="/" end className={({ isActive }) => `nav-link ${isActive ? 'active' : ''}`}>
            <HiOutlineHome /> Dashboard
          </NavLink>
          <NavLink to="/articles" className={({ isActive }) => `nav-link ${isActive ? 'active' : ''}`}>
            <HiOutlineDocumentText /> Artikel
          </NavLink>
          <NavLink to="/authors" className={({ isActive }) => `nav-link ${isActive ? 'active' : ''}`}>
            <HiOutlineUsers /> Penulis
          </NavLink>
          <NavLink to="/categories" className={({ isActive }) => `nav-link ${isActive ? 'active' : ''}`}>
            <HiOutlineTag /> Kategori
          </NavLink>
          <NavLink to="/comments" className={({ isActive }) => `nav-link ${isActive ? 'active' : ''}`}>
            <HiOutlineChatBubbleLeftRight /> Komentar
          </NavLink>
        </div>
      </nav>
    </aside>
  );
}
