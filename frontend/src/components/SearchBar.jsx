import { useState } from 'react';
import './SearchBar.css';

export default function SearchBar({ onSearch }) {
  const [query, setQuery] = useState('');

  function handleSubmit(event) {
    event.preventDefault();
    onSearch?.(query);
  }

  return (
    <form className="search-bar" onSubmit={handleSubmit} role="search">
      <div className="container search-bar__row">
        <svg
          className="search-bar__icon"
          viewBox="0 0 24 24"
          width="18"
          height="18"
          aria-hidden="true"
        >
          <circle cx="11" cy="11" r="7" fill="none" stroke="currentColor" strokeWidth="2" />
          <line x1="16.5" y1="16.5" x2="21" y2="21" stroke="currentColor" strokeWidth="2" strokeLinecap="round" />
        </svg>
        <input
          className="search-bar__input"
          type="search"
          placeholder="Search for products, brands, or categories"
          value={query}
          onChange={(event) => setQuery(event.target.value)}
          aria-label="Search products"
        />
        <button className="search-bar__button" type="submit">
          Search
        </button>
      </div>
    </form>
  );
}
