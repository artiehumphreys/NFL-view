import React, { useState, useEffect } from "react";
import { FaSearch, FaBars, FaTimes, FaTrash } from "react-icons/fa";
import styles from "../css/HomePage.module.css";
import Header from "../components/Header.js";
import Footer from "../components/Footer.js";
import Modal from "../components/Modal.js";
import Success from "../components/Success.js";
import SideBar from "../components/SideBar.js";

function HomePage() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  const [isModalOpen, setIsModalOpen] = useState(false);
  const [currentEvent, setCurrentEvent] = useState(null);

  const toggleModal = () => {
    setIsModalOpen(!isModalOpen);
  };

  const [isSuccessOpen, setIsSuccessOpen] = useState(false);
  const toggleSuccess = () => {
    setIsSuccessOpen(!isSuccessOpen);
  };

  const [searchTags, setSearchTags] = useState([]);
  const [displayInfo, setDisplayInfo] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/tags")
      .then((response) => response.json())
      .then((data) => setSearchTags(data))
      .catch((error) => console.error("Error fetching search tags:", error));

    fetch("http://localhost:8080/info")
      .then((response) => response.json())
      .then((data) => setDisplayInfo(data))
      .catch((error) => console.error("Error fetching display info:", error));
  }, []);

  useEffect(() => {
    if (localStorage.getItem("deleteSuccess") === "true") {
      toggleSuccess();
      localStorage.removeItem("deleteSuccess");
    }
  }, []);

  const deleteInjury = async (play_id, game, fName, lName, type) => {
    const response = await fetch(
      `http://localhost:8080/removeInjury?game=${game}&play_id=${play_id}&fName=${fName}&lName=${lName}&type=${type}`,
      {
        method: "DELETE",
      }
    );
    if (response.ok) {
      localStorage.setItem("deleteSuccess", "true");
      window.location.reload();
    } else {
      console.error(`Failed to remove item ${response.url}`);
    }
  };

  return (
    <div className={`${styles.container} min-h-screen flex flex-col`}>
      <Header path="/"></Header>
      <Modal
        isOpen={isModalOpen}
        onClose={toggleModal}
        onDelete={() =>
          deleteInjury(
            currentEvent.PlayID,
            currentEvent.Game,
            currentEvent.FirstName,
            currentEvent.LastName,
            currentEvent.Type
          )
        }
      />
      <Success
        isOpen={isSuccessOpen}
        onChange={toggleModal}
        message="Entry successfully removed."
      ></Success>
      <div className="flex flex-1 relative overflow-auto">
        <SideBar></SideBar>
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
                  onClick={() => {
                    setCurrentEvent(info);
                  }}
                  className="flex bg-gray-200 text-gray-700 py-2 px-3 rounded hover:bg-gray-300 justify-between items-center"
                >
                  <div className="flex-grow text-center">
                    {info.Game} - {info.FirstName} {info.LastName} -{" "}
                    {info.GamePosition} {info.Team} #{info.JerseyNumber} -{" "}
                    {info.Type}
                  </div>
                  <FaTrash
                    className="ml-2 cursor-pointer"
                    onClick={(e) => {
                      e.stopPropagation();
                      setCurrentEvent(info);
                      toggleModal();
                    }}
                  />
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
      <Footer />
    </div>
  );
}

export default HomePage;
