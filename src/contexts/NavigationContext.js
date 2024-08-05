import React, { createContext, useContext, useState, useCallback } from "react";
import { useNavigate } from "react-router-dom";

const NavigationContext = createContext();

export const NavigationProvider = ({ children }) => {
  const [history, setHistory] = useState([]);
  const navigate = useNavigate();
  console.log(history);

  const push = useCallback((path) => {
    setHistory((prev) => {
      if (prev.length === 0 || prev[prev.length - 1] !== path) {
        return [...prev, path];
      }
      return prev;
    });
  }, []);

  const pop = useCallback(() => {
    setHistory((prev) => {
      const newHistory = [...prev];
      newHistory.pop();
      return newHistory;
    });
  }, []);

  const goBack = useCallback(() => {
    if (history.length > 1) {
      pop();
      navigate(history[history.length - 2]);
    }
  }, [history, navigate, pop]);

  return (
    <NavigationContext.Provider value={{ history, push, goBack }}>
      {children}
    </NavigationContext.Provider>
  );
};

export const useNavigation = () => {
  return useContext(NavigationContext);
};
