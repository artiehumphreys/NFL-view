import React, { useState, useEffect } from "react";
import { FaSearch, FaBars, FaTimes, FaTrash } from "react-icons/fa";
import styles from "../css/HomePage.module.css";
import { useNavigate } from "react-router-dom";
import Header from "../components/Header.js";
import Footer from "../components/Footer.js";

function HomePage() {
  const navigate = useNavigate();

  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  const [visibleGameIndex, setVisibleGameIndex] = useState(null);
  const toggleVisibility = (index) => {
    setVisibleGameIndex(visibleGameIndex === index ? null : index);
  };

  const [searchTags, setSearchTags] = useState([]);
  const [displayInfo, setDisplayInfo] = useState([]);
  const [gameInfo, setGameInfo] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/tags")
      .then((response) => response.json())
      .then((data) => setSearchTags(data))
      .catch((error) => console.error("Error fetching search tags:", error));

    fetch("http://localhost:8080/info")
      .then((response) => response.json())
      .then((data) => setDisplayInfo(data))
      .catch((error) => console.error("Error fetching display info:", error));

    fetch("http://localhost:8080/games")
      .then((response) => response.json())
      .then((data) => setGameInfo(data))
      .catch((error) => console.error("Error fetching games:", error));
  }, []);
  return (
    <div className={`${styles.container} min-h-screen flex flex-col`}>
      <Header path="/"></Header>
      <div className="flex flex-1 relative overflow-auto">
        <aside className="w-64 bg-gray-100 px-3 py-4 absolute left-0 top-0 bottom-0 flex flex-col flex-1">
          <div className="flex-shrink-0 mb-4">
            <h2 className="text-2xl font-bold">Games</h2>
          </div>
          <div className="flex-1 overflow-auto">
            {gameInfo.map((game, index) => (
              // https://preline.co/docs/dropdown.html
              <div key={index} className="mb-4">
                <div
                  onClick={() => toggleVisibility(index)}
                  className={
                    "flex gap-2 items-center font-medium text-lg py-2 px-1 rounded-lg border-gray-200 hover:bg-gray-50 focus:bg-gray-50 cursor-pointer w-full"
                  }
                >
                  <div
                    className={`gap-2 flex ${
                      visibleGameIndex === index
                        ? "hover:underline decoration-black"
                        : ""
                    }`}
                    onClick={() =>
                      visibleGameIndex === index
                        ? navigate(`/games/${game.game}`)
                        : ""
                    }
                  >
                    <span>{game.game}</span>
                    <span className="font-normal">({game.events.length})</span>
                  </div>
                  <svg
                    className={`hs-dropdown-open:rotate-180 ml-auto ${
                      visibleGameIndex === index ? "rotate-180" : ""
                    }`}
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                  >
                    <path d="m6 9 6 6 6-6" />
                  </svg>
                </div>
                {visibleGameIndex === index && (
                  <ul className="gap-2 py-2 flex flex-col text-left">
                    {game.events.map((event, idx) => (
                      <button
                        key={idx}
                        className="bg-gray-200 text-gray-700 py-2 mx-2 px-2 my-1 rounded hover:bg-gray-300 text-left"
                      >
                        {event}
                      </button>
                    ))}
                  </ul>
                )}
              </div>
            ))}
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
            <ul className="flex gap-2 flex-col align-center pb-2 overflow-scroll justify-between">
              {displayInfo.map((info, index) => (
                <button
                  key={index}
                  className="flex bg-gray-200 text-gray-700 py-2 px-3 rounded hover:bg-gray-300 justify-between items-center"
                >
                  <div className="w-10"></div>
                  {info}
                  <FaTrash className="ml-2" />
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
      <Footer></Footer>
    </div>
  );
}

export default HomePage;
