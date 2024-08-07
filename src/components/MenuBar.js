import React, { useState } from "react";
import styles from "../css/HomePage.module.css";
import { FaBars, FaTimes } from "react-icons/fa";

function MenuBar() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  return (
    <>
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
    </>
  );
}

export default MenuBar;
