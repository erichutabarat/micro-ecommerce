import './ProductList.css';

const PRODUCTS = [
  { id: 1, name: 'Ceramic Pour-Over Set', price: 32, rating: 4.8, color: '#e8a33d' },
  { id: 2, name: 'Canvas Tote, Undyed', price: 18, rating: 4.6, color: '#1f4b43' },
  { id: 3, name: 'Enamel Camp Mug', price: 14, rating: 4.9, color: '#b4532a' },
  { id: 4, name: 'Recycled Wool Throw', price: 54, rating: 4.7, color: '#23201b' },
  { id: 5, name: 'Beeswax Wrap, Set of 3', price: 21, rating: 4.5, color: '#c8811f' },
  { id: 6, name: 'Bamboo Cutting Board', price: 27, rating: 4.8, color: '#1f4b43' },
  { id: 7, name: 'Linen Tea Towel Pair', price: 16, rating: 4.4, color: '#b4532a' },
  { id: 8, name: 'Cast Iron Trivet', price: 22, rating: 4.9, color: '#e8a33d' },
];

function Stars({ rating }) {
  return (
    <span className="product-card__rating" aria-label={`Rated ${rating} out of 5`}>
      ★ {rating.toFixed(1)}
    </span>
  );
}

export default function ProductList() {
  return (
    <section className="products container">
      <div className="products__header">
        <h2 className="products__title">Popular right now</h2>
        <a className="products__link" href="/products">
          View all
        </a>
      </div>

      <div className="products__grid">
        {PRODUCTS.map((product) => (
          <article key={product.id} className="product-card">
            <div
              className="product-card__image"
              style={{ background: product.color }}
              aria-hidden="true"
            />
            <div className="product-card__body">
              <p className="product-card__name">{product.name}</p>
              <div className="product-card__meta">
                <span className="product-card__price">${product.price}</span>
                <Stars rating={product.rating} />
              </div>
              <button className="product-card__add">Add to cart</button>
            </div>
          </article>
        ))}
      </div>
    </section>
  );
}
