import styles from "../css/HomePage.module.css";
import { FaSearch } from "react-icons/fa";

function HomePage() {
  return (
    <div className={`${styles.container} min-h-screen flex flex-col`}>
      <header className="bg-blue-500 text-white text-center py-4">
        <h1 className="text-5xl font-bold">NFL View</h1>
      </header>
      <div className="flex flex-1">
        <main className="flex-1 p-3 flex flex-col items-center">
          <section className="my-1 content-center">
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
            <div className="flex gap-2 flex-row items-center">
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
