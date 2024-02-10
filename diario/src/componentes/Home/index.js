import './home.css';

export const Home = () => {

    return (
        <>
          <main>
            <div className="div-home">
            <h1 class="title-home"> Bienvenido a Diario Cordoba</h1>
              <a className="boton-home"  href={`http://localhost:3000/inicio`}>
                <h3>¡Comenzar a leer!</h3>
              </a>
            </div>
          </main>
        </>
      )
}