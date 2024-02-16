import React from "react"
import { Routes, Route } from "react-router-dom";
import { Home } from "./Home";
import { Noticias } from "./Inicio";
import { Publicar } from "./Publicar";

export const Paginas = () => {
    return (
        <Routes>
            <Route path='/' exact element={<Home />} />
            <Route path='/inicio' exact element={<Noticias />} />
            <Route path='/publicar' exact element={<Publicar />} />
        </Routes>
    )
}