import SearchBar from '../components/SearchBar.jsx';
import PromoBanner from '../components/PromoBanner.jsx';
import CategoryList from '../components/CategoryList.jsx';
import ProductList from '../components/ProductList.jsx';

export default function Home() {
  function handleSearch(query) {
    console.log('Search submitted:', query);
    // TODO: route this through the Broker service once it exists
  }

  return (
    <>
      <SearchBar onSearch={handleSearch} />
      <PromoBanner />
      <CategoryList />
      <ProductList />
    </>
  );
}
