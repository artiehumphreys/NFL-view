import { FaSearch } from "react-icons/fa";
import styles from "../css/HomePage.module.css";

function HomePage() {
  return (
    <div className={`${styles.container} min-h-screen flex flex-col`}>
      <header className="bg-blue-500 text-white text-center py-4">
        <h1 className="text-5xl font-bold">NFL View</h1>
      </header>
      <div className="flex flex-1 relative">
        <aside className="w-64 bg-gray-100 p-4 absolute left-0 top-0 bottom-0">
          <h2 className="text-2xl font-bold mb-4">Games</h2>
        </aside>
        <main className="flex-1 p-6 flex flex-col items-center">
          <section className="my-1 w-max">
            <div className="flex items-center mb-4">
              <input
                type="text"
                placeholder="Search..."
                className="flex-grow p-2 border border-gray-300 rounded-l"
              />
              <button className="bg-blue-500 text-white p-2 rounded-r">
                <FaSearch className="w-6 h-6" />
              </button>
            </div>
            <div className="flex gap-2 flex-row items-center mb-4">
              <h2 className="text-xl font-semibold">Search by</h2>
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
        </main>
      </div>
    </div>
  );
}

export default HomePage;
