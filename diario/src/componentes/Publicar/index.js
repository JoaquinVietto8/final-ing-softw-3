import React, { useState} from 'react'
import './publicar.css';

export function Publicar(){
    
    const[contenido,setContenido] = useState("");
    const[imagen,setImagen] = useState("");

    const onChangeContenido =  (contenido)=>{
        setContenido(contenido.target.value);
    }
    
    const onChangeImagen = (imagen)=>{
    setImagen(imagen.target.value)};

    const requestOptions={
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        
        body: JSON.stringify({contenido : contenido, imagen : imagen })
    };

    const publicar = async()=>{
        fetch('https://backend-y7etararqa-uc.a.run.app/nueva-noticia',requestOptions)
        .then(response=>response.json())
        .then(response => {if (response.status == 400) {
            document.getElementById("contenido").style.border = "1px solid red";
            document.getElementById("imagen").style.border = "1px solid red";
            document.getElementById("incorrecto").style.visibility = "visible"
        }else{
            window.location.replace("/inicio")  
        }})

    }; 

    function color(id) {
        document.getElementById(id).style.border = "1px solid #c7c7c7";
        document.getElementById("incorrecto").style.visibility = "hidden"
    }
   
    const handleSubmit= (event)=>{

            event.preventDefault();
       
            publicar();
    };

    return(
        <main>
            <div class="contenedor-login">
                <h1 class="title-login">Publicar Noticia</h1>
                <form onSubmit={handleSubmit} class="form-login">
                    <div class="div-input">
                        <input id="contenido" type="text" autocomplete="off" placeholder="Contenido" class="input-datos" maxlength="350" onChange={onChangeContenido} onClick={() => color("contenido")} value ={contenido}/>
                    </div>
                    <div class="div-input">
                        <input id="imagen" type="text" autocomplete="off" placeholder="URL Imagen" class="input-datos" maxlength="350" onChange={onChangeImagen} onClick={() => color("imagen")} value ={imagen}/>
                    </div>
                    <div class="div-confirmar">
                        <button class="confirmar" type="submit">Publicar</button>
                        <p id='incorrecto'>Intentalo de nuevo.</p>
                    </div>
                </form>
            </div>
        </main>
    );   
}