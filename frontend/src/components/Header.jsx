import { useState } from 'react';
import './Header.css';

const NAV_LINKS = ['Home', 'Categories', 'Deals', 'About'];

export default function Header() {
  const [menuOpen, setMenuOpen] = useState(false);

  return (
    <header className="header">
      <div className="container header__row">
        <a className="header__logo" href="/">
          Micro<em>market</em>
        </a>

        <nav className={`header__nav ${menuOpen ? 'header__nav--open' : ''}`}>
          {NAV_LINKS.map((link) => (
            <a key={link} href={`/${link.toLowerCase()}`} className="header__link">
              {link}
            </a>
          ))}
        </nav>

        <div className="header__actions">
          <button className="header__signin">Sign in</button>
          <button className="header__signup">Sign up</button>
          <button
            className="header__menu-toggle"
            aria-label="Toggle navigation"
            aria-expanded={menuOpen}
            onClick={() => setMenuOpen((open) => !open)}
          >
            <span />
            <span />
            <span />
          </button>
        </div>
      </div>
    </header>
  );
}
