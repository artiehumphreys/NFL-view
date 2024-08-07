import React, { useState, useEffect } from "react";
import { useNavigation } from "../contexts/NavigationContext.js";
import { useLocation } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import { FaSearch, FaTrash } from "react-icons/fa";
import styles from "../css/HomePage.module.css";
import Header from "../components/Header.js";
import Footer from "../components/Footer.js";
import Modal from "../components/Modal.js";
import Success from "../components/Success.js";
import SideBar from "../components/SideBar.js";
import MenuBar from "../components/MenuBar.js";

function HomePage() {
  const navigate = useNavigate();
  const { push } = useNavigation();
  const location = useLocation();

  useEffect(() => {
    push(location.pathname);
  }, [location.pathname, push]);

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
  const [injuries, setInjuries] = useState([]);
  const [currentTag, setCurrentTag] = useState("");

  useEffect(() => {
    fetch("http://localhost:8080/tags")
      .then((response) => response.json())
      .then((data) => setSearchTags(data))
      .catch((error) => console.error("Error fetching search tags:", error));

    fetch("http://localhost:8080/injuries")
      .then((response) => response.json())
      .then((data) => setInjuries(data))
      .catch((error) => console.error("Error fetching display info:", error));
  }, []);

  useEffect(() => {
    if (localStorage.getItem("deleteSuccess") === "true") {
      toggleSuccess();
      localStorage.removeItem("deleteSuccess");
    }
  });

  const deleteInjury = async (play_id, game, fName, lName, type) => {
    const response = await fetch(
      `http://localhost:8080/injuries?game=${game}&play_id=${play_id}&fName=${fName}&lName=${lName}&type=${type}`,
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

  const handleTagClick = async (tag) => {
    setCurrentTag(tag);
    try {
      const response = await fetch(`http://localhost:8080/injuries?tag=${tag}`);
      const data = (await response.json()) || [];
      setInjuries(data);
    } catch (error) {
      console.error("Error fetching injuries:", error);
    }
  };

  const handleSearchClick = async (keyword) => {
    setCurrentTag("");
    try {
      const response = await fetch(
        `http://localhost:8080/injuries?search=${keyword}`
      );
      const data = (await response.json()) || [];
      setInjuries(data);
    } catch (error) {
      console.error("Error fetching injuries:", error);
    }
  };

  return (
    <div className={`${styles.container} min-h-screen flex flex-col`}>
      <Header></Header>
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
      <div className="flex flex-1 relative overflow-hidden">
        <SideBar></SideBar>
        <main className="flex-1 p-6 flex flex-col items-center ml-2">
          <section className="flex flex-col my-1 w-max max-h-full">
            <div className="flex items-center mb-4">
              <input
                type="text"
                placeholder="Search..."
                className="flex-grow p-2 border border-gray-300 rounded-l"
                id="input"
              />
              <button className="bg-blue-500 text-white p-2 rounded-r">
                <FaSearch
                  className="w-6 h-6"
                  onClick={() =>
                    handleSearchClick(document.getElementById("input").value)
                  }
                />
              </button>
            </div>
            <div className="flex gap-2 flex-row items-center mb-4">
              <h2 className="text-xl font-semibold">Search by</h2>
              {searchTags.map((tag, index) => (
                <button
                  key={index}
                  className={`bg-gray-200 text-gray-700 py-1 px-3 rounded hover:bg-gray-400 active:bg-gray-300 ${
                    currentTag === tag ? "bg-gray-400" : ""
                  }`}
                  onClick={() => handleTagClick(tag)}
                >
                  {tag}
                </button>
              ))}
            </div>
            <ul className="flex gap-2 flex-col align-center pb-2 overflow-scroll justify-between">
              {injuries.map((info, index) => (
                <button
                  key={index}
                  onClick={() => {
                    setCurrentEvent(info);
                    navigate(`/games/${info.Game}/plays/${info.PlayID}`);
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
        <MenuBar></MenuBar>
      </div>
      <Footer />
    </div>
  );
}

export default HomePage;
