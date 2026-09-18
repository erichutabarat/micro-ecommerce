import { Link } from 'react-router-dom';

export default function NotFound() {
  return (
    <section className="container" style={{ padding: '64px 24px', textAlign: 'center' }}>
      <h1>Page not found</h1>
      <p style={{ marginTop: 12, color: 'var(--color-ink-soft)' }}>
        The page you're looking for doesn't exist.
      </p>
      <Link
        to="/"
        style={{
          display: 'inline-block',
          marginTop: 24,
          fontWeight: 600,
          color: 'var(--color-teal)',
        }}
      >
        Back to home
      </Link>
    </section>
  );
}
