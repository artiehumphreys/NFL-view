import React, { useState, useEffect } from "react";
import { FaSearch, FaBars, FaTimes } from "react-icons/fa";
import styles from "../css/HomePage.module.css";

function HomePage() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  const [searchTags, setSearchTags] = useState([]);
  const [displayInfo, setDisplayInfo] = useState([]);
  const [gameInfo, setGameInfo] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/home/tags")
      .then((response) => response.json())
      .then((data) => setSearchTags(data))
      .catch((error) => console.error("Error fetching search tags:", error));

    fetch("http://localhost:8080/home/info")
      .then((response) => response.json())
      .then((data) => setDisplayInfo(data))
      .catch((error) => console.error("Error fetching display info:", error));

    fetch("http://localhost:8080/home/games")
      .then((response) => response.json())
      .then((data) => setGameInfo(data))
      .catch((error) => console.error("Error fetching games:", error));
  }, []);
  return (
    <div className={`${styles.container} min-h-screen flex flex-col`}>
      <header className="bg-blue-500 text-white text-center py-4">
        <h1 className="text-5xl font-bold">NFL View</h1>
      </header>
      <div className="flex flex-1 relative overflow-auto">
        <aside className="w-64 bg-gray-100 p-4 absolute left-0 top-0 bottom-0 flex flex-col">
          <div className="flex-shrink-0 mb-4">
            <h2 className="text-2xl font-bold">Games</h2>
          </div>
          <div className="flex-1 overflow-auto">
            <ul className="flex flex-col gap-2">
              {gameInfo.map((info, index) => (
                <button
                  key={index}
                  className="bg-gray-200 text-gray-700 py-1 px-3 rounded hover:bg-gray-300"
                >
                  {info}
                </button>
              ))}
            </ul>
          </div>
        </aside>
        <main className="flex-1 p-6 flex flex-col items-center">
          <section className="flex flex-col my-1 w-max max-h-full">
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
              {searchTags.map((tag, index) => (
                <button
                  key={index}
                  className="bg-gray-200 text-gray-700 py-1 px-3 rounded hover:bg-gray-300"
                >
                  {tag}
                </button>
              ))}
            </div>
            <ul className="flex gap-2 flex-col align-left mb-4 overflow-scroll">
              {displayInfo.map((info, index) => (
                <button
                  key={index}
                  className="bg-gray-200 text-gray-700 py-1 px-3 rounded hover:bg-gray-300"
                >
                  {info}
                </button>
              ))}
            </ul>
          </section>
        </main>
        <button
          className={` ${styles.button} button absolute top-4 right-4 text-blue-500 p-2 rounded`}
          onClick={toggleMenu}
        >
          {isMenuOpen ? <FaTimes /> : <FaBars />}
        </button>
        <aside
          className={`w-64 bg-gray-100 p-4 flex flex-col absolute right-0 top-0 bottom-0 transition-transform transform ${
            isMenuOpen ? "-translate-x-0" : "translate-x-full"
          }`}
        >
          <h2 className="text-2xl font-bold mb-4">Menu</h2>
          <ul>
            <li className="mb-2">
              <a href="#" className="text-blue-500">
                Profile
              </a>
            </li>
            <li className="mb-2">
              <a href="#" className="text-blue-500">
                Settings
              </a>
            </li>
            <li className="mb-2">
              <a href="#" className="text-blue-500">
                View Legacy Data
              </a>
            </li>
          </ul>
        </aside>
      </div>
    </div>
  );
}

export default HomePage;
