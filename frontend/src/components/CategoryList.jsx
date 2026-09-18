import './CategoryList.css';

const CATEGORIES = [
  { name: 'Kitchen', glyph: '◐' },
  { name: 'Fashion', glyph: '◧' },
  { name: 'Electronics', glyph: '◫' },
  { name: 'Home', glyph: '◩' },
  { name: 'Beauty', glyph: '○' },
  { name: 'Sports', glyph: '◔' },
  { name: 'Books', glyph: '▭' },
  { name: 'Toys', glyph: '△' },
];

export default function CategoryList() {
  return (
    <section className="categories container">
      <h2 className="categories__title">Shop by category</h2>
      <div className="categories__scroll">
        {CATEGORIES.map((category) => (
          <button key={category.name} className="categories__chip">
            <span className="categories__glyph" aria-hidden="true">
              {category.glyph}
            </span>
            {category.name}
          </button>
        ))}
      </div>
    </section>
  );
}
