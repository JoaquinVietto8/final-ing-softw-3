import React from "react";

export const ObjectInicio = (
    { id,
        contenido,
        imagen,
        fecha,
    }) => {
    return (
        <div class="noticia">
            <div class="datos-padre">
                <div class="div-contenido">
                    <p class="txt-contenido">{contenido}</p>
                </div>
                <div class="div-fecha">
                    <p class="fecha">{fecha}</p>
                </div>
            </div>
            <div class="contender-imagen">
                <img class="imagen" src={imagen} />
            </div>
        </div>
    )
}