import React from "react";
import { useSpring, animated } from "@react-spring/web";
import styles from "../css/NFLLogo.module.css";

function NFLLogo() {
  const props = useSpring({
    from: { transform: "translateY(-100%)" },
    to: { transform: "translateY(50vh)" },
    config: { duration: 500 },
  });
  return (
    <div className="logo-container">
      <animated.img
        src="/nfl_logo.png"
        style={props}
        alt="NFL Logo"
        className={styles.logo}
      />
    </div>
  );
}

export default NFLLogo;
