import styles from "../css/HomePage.module.css";

function HomePage() {
  return (
    <div className={`${styles.container} min-h-screen flex flex-col`}>
      <header className="bg-blue-500 text-white text-center py-4">
        <h1 className="text-5xl font-bold">NFL View</h1>
      </header>
      <div className="flex flex-1">
        <section>
          <div className="flex gap-2">
            <h2 className="text-xl font-semibold mb-2">Search by</h2>
            <button className="bg-gray-200 text-gray-700 py-1 px-3 rounded">
              Tag 1
            </button>
            <button className="bg-gray-200 text-gray-700 py-1 px-3 rounded">
              Tag 2
            </button>
            <button className="bg-gray-200 text-gray-700 py-1 px-3 rounded">
              Tag 3
            </button>
            <button className="bg-gray-200 text-gray-700 py-1 px-3 rounded">
              Tag 4
            </button>
            <button className="bg-gray-200 text-gray-700 py-1 px-3 rounded">
              Tag 5
            </button>
          </div>
        </section>
      </div>
    </div>
  );
}

export default HomePage;
