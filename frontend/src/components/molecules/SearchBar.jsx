import React from 'react';
import { HiMagnifyingGlass } from 'react-icons/hi2';

export default function SearchBar({ value, onChange, placeholder = 'Cari...' }) {
  return (
    <div className="search-bar">
      <HiMagnifyingGlass />
      <input 
        placeholder={placeholder} 
        value={value} 
        onChange={onChange} 
      />
    </div>
  );
}
