import React from "react"
import { Routes, Route } from "react-router-dom";
import { Home } from "./Home";
import { Noticias } from "./Inicio";
import { Publicar } from "./Publicar";

export const Paginas = () => {
    return (
        <Routes>
            <Route path='/' exact element={<Home />} />
            <Route path='https://gc-final-ing-softw-3-frontend-a2l4gusvua-uc.a.run.app/inicio' exact element={<Noticias />} />
            <Route path='https://gc-final-ing-softw-3-frontend-a2l4gusvua-uc.a.run.app/publicar' exact element={<Publicar />} />
        </Routes>
    )
}