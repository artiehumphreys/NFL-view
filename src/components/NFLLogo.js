import React from "react";
import { useSpring, animated } from "@react-spring/web";
import styles from "../css/NFLLogo.module.css";

function NFLLogo() {
  const props = useSpring({
    from: { opacity: 0 },
    to: { opacity: 1 },
    config: { duration: 1000 },
  });
  return (
    <div className={styles.container}>
      <img
        src="/nfl_logo.png"
        style={props}
        alt="NFL Logo"
        className={styles.logo}
      />
      <span className={styles.text}>NFL View</span>
    </div>
  );
}

export default NFLLogo;
