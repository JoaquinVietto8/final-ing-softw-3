Feature('sample');

Scenario('test index',  ({ I }) => {
    I.amOnPage('')
    I.wait(5)
    I.see('Bienvenido a Diario Cordoba')
    I.click('a.boton-home');
    I.wait(3)
    I.see('Diario Cordoba')
});

Scenario('test home', ({ I }) => {
    I.amOnPage("inicio")
    I.wait(3)
    I.see('Diario Cordoba')
    I.seeElement('.noticia')
    I.click('a.publicar');
    I.wait(3)
    I.see('Publicar Noticia')
})

Scenario('test post noticia', ({ I }) => {
    I.amOnPage("publicar")
    I.wait(3)
    I.see('Publicar Noticia')
    I.seeElement('input#contenido.input-datos')
    I.seeElement('input#imagen.input-datos')
    I.fillField('input#contenido.input-datos', 'Noticia de Prueba-2')
    I.fillField('input#imagen.input-datos', 'https://tn.com.ar/resizer/v2/los-britanicos-eligen-hoy-el-nuevo-rumbo-del-pais-foto-efe-ANMIPDACNZB2RD22NDCO4YA2Y4.jpg?auth=4c0501428165809c12864dd4589ca16fd789d09af79ba7156f5abf2b78419ee2&width=767')
    I.wait(2)
    I.click('button.confirmar');
    I.wait(2)
    I.see('Diario Cordoba')
    I.see('Noticia de Prueba-2')
    I.seeElement('img.imagen[src="https://tn.com.ar/resizer/v2/los-britanicos-eligen-hoy-el-nuevo-rumbo-del-pais-foto-efe-ANMIPDACNZB2RD22NDCO4YA2Y4.jpg?auth=4c0501428165809c12864dd4589ca16fd789d09af79ba7156f5abf2b78419ee2&width=767"]')
})