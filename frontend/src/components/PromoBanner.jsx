import './PromoBanner.css';

export default function PromoBanner() {
  return (
    <section className="promo container">
      <div className="promo__text">
        <span className="promo__kicker">First order</span>
        <h1 className="promo__headline">
          20% off, no minimum, all this week
        </h1>
        <p className="promo__sub">
          Local sellers, everyday goods, delivered in two days or less.
        </p>
        <button className="promo__cta">Shop the sale</button>
      </div>

      <svg
        className="promo__art"
        viewBox="0 0 320 260"
        aria-hidden="true"
      >
        <rect x="30" y="70" width="150" height="150" rx="10" fill="var(--color-marigold)" />
        <circle cx="235" cy="95" r="65" fill="var(--color-teal)" />
        <rect x="150" y="150" width="110" height="80" rx="8" fill="var(--color-brick)" opacity="0.9" />
      </svg>
    </section>
  );
}
