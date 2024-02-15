import React, { useEffect, useState } from "react";
import { ObjectInicio } from "./ObjectInicio";
import './inicio.css';

export const Noticias = () => {

    const [noticias, setNoticias] = useState([])
    const fetchApi = async () => {
        const response = await fetch('http://localhost:8080/inicio')
            .then((response) => response.json())
        setNoticias(response)
        console.log(response);
    };
    useEffect(() => {
        fetchApi();
    }, [])
    return (
        <>
            <header className="header-noticias">
            <h1 className="titulo">
                    Diario Cordoba
                </h1>
                <a className="publicar" href="https://gc-final-ing-softw-3-frontend-a2l4gusvua-uc.a.run.app/publicar">
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
}