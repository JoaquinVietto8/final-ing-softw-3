import './home.css';

export const Home = () => {

    return (
        <>
          <main>
            <div className="div-home">
            <h1 class="title-home"> Bienvenido a Diario Cordoba</h1>
              <a className="boton-home"  href={`https://gc-final-ing-softw-3-frontend-a2l4gusvua-uc.a.run.app/inicio`}>
                <h3>¡Comenzar a leer!</h3>
              </a>
            </div>
          </main>
        </>
      )
}