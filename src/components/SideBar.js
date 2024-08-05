import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

function SideBar() {
  const navigate = useNavigate();

  const [gameInfo, setGameInfo] = useState([]);
  const [visibleGameIndex, setVisibleGameIndex] = useState(null);

  const toggleVisibility = (index) => {
    setVisibleGameIndex(visibleGameIndex === index ? null : index);
  };

  useEffect(() => {
    fetch("http://localhost:8080/games")
      .then((response) => response.json())
      .then((data) => setGameInfo(data))
      .catch((error) => console.error("Error fetching games:", error));
  }, []);

  return (
    <aside className="w-64 bg-gray-100 px-3 py-4 absolute left-0 top-0 bottom-0 flex flex-col flex-1">
      <div className="flex-shrink-0 mb-4">
        <h2 className="text-2xl font-bold">Games</h2>
      </div>
      <div className="flex-1 overflow-auto">
        {gameInfo.map((game, index) => (
          <div key={index} className="mb-4">
            <div
              onClick={() => toggleVisibility(index)}
              className={
                "flex gap-2 items-center font-medium text-lg py-2 px-1 rounded-lg border-gray-200 hover:bg-gray-50 focus:bg-gray-50 cursor-pointer w-full"
              }
            >
              <div
                className="gap-2 flex hover:underline decoration-black"
                onClick={() => navigate(`/games/${game.game}`)}
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
  );
}

export default SideBar;
