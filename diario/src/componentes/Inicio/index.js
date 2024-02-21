import React, { useEffect, useState } from "react";
import { ObjectInicio } from "./ObjectInicio";
import './inicio.css';

export const Noticias = () => {

    const [noticias, setNoticias] = useState([])
    const fetchApi = async () => {
        const response = await fetch('https://backend-a2l4gusvua-uc.a.run.app/inicio')
            .then((response) => response.json())
        setNoticias(response)
        console.log(response);
        if (response == null) {
            vacio = true
        } else {
            vacio = false
        }
    };
    useEffect(() => {
        fetchApi();
    }, [])
    if (vacio === false) {
        return (
            <>
                <header className="header-noticias">
                    <h1 className="titulo">
                        Diario Cordoba
                    </h1>
                    <a className="publicar" href="/publicar">
                        Publicar
                    </a>
                </header>
                <main>
                    <div class="contenedor-noticias">
                        {
                            noticias.map(noticia => (
                                <ObjectInicio key={noticia.id}
                                    id={noticia.id}
                                    contenido={noticia.contenido}
                                    imagen={noticia.imagen}
                                    fecha={noticia.fecha}
                                />
                            ))
                        }
                    </div>
                </main>
            </>
        )
    } else {
        return (
            <>
                <header className="header-noticias">
                    <h1 className="titulo">
                        Diario Cordoba
                    </h1>
                    <a className="publicar" href="/publicar">
                        Publicar
                    </a>
                </header>
                <main>
                    <div class="contenedor-noticias">
                    <h2>No hay noticias para mostrar.</h2>
                    </div>
                </main>
            </>
        )

    }
}