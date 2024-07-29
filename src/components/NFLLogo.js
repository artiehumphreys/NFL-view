import React from "react";
import { useNavigate } from "react-router-dom";
import { useSpring, animated } from "@react-spring/web";
import styles from "../css/NFLLogo.module.css";

function NFLLogo() {
  const props = useSpring({
    from: { opacity: 0 },
    to: { opacity: 1 },
    config: { duration: 1500 },
    delay: 750,
  });
  const navigate = useNavigate();
  return (
    <div className="flex flex-col justify-start items-center h-screen py-10 mb-5 min-h-screen max-h-screen">
      <img src="/nfl_logo.png" alt="NFL Logo" className={`${styles.logo}`} />
      <span className={`${styles.text}`}>NFL View</span>
      <animated.button
        style={props}
        className={`bg-blue-500 font-bold text-white py-2 px-4 border border-blue-700 rounded ${styles.button}`}
        onClick={() => navigate("/home")}
      >
        Get Started
      </animated.button>
    </div>
  );
}

export default NFLLogo;
