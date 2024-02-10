import React from "react"
import { Routes, Route } from "react-router-dom";
import { Home } from "./Home";

export const Paginas = () => {
    return (
        <Routes>
            <Route path='/' exact element={<Home />} />
        </Routes>
    )
}