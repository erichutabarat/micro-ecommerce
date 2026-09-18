export default function About() {
  return (
    <section className="container" style={{ padding: '48px 24px' }}>
      <h1>About Micromarket</h1>
      <p style={{ marginTop: 12, color: 'var(--color-ink-soft)', maxWidth: '60ch' }}>
        Micromarket is a learning project exploring how a small ecommerce app
        can be built as independent services: a frontend, a broker, auth,
        product, and logging.
      </p>
    </section>
  );
}
