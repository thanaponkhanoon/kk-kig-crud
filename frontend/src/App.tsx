import React, { useEffect, useState } from "react";

import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";
import SignIn from "./components/SignIn";
import Home from "./components/home";
import Customers from "./components/Customers";
import Products from "./components/Products"
import Orders from "./components/Orders";
import Process from "./components/Process";
import Reports from "./components/Reports";

export default function App() {
  const [token, setToken] = useState<string | null>();

  useEffect(() => {
    setToken(localStorage.getItem("token"));
  }, []);

  if (!token) {
    return <SignIn />;
  }

  return (

    <Router>

      <div>

        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/customers" element={<Customers />} />
          <Route path="/products" element={<Products />} />
          <Route path="/order" element={<Orders />} />
          <Route path="/process" element={<Process />} />
          <Route path="/report" element={<Reports />} />
        </Routes>

      </div>

    </Router>

  );

}

