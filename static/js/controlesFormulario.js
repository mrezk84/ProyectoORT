$(document).ready(function() {
    buildControles()
});



let formulario = null;
async function buildControles() {
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    formularioid = searchParams.get('formulario_id');
    datos.id_formulario = Number(formularioid);

    const request = await fetch("http://localhost:5000/controles/byForm", {
               method: 'POST',
               body: JSON.stringify(datos),
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })

    const controles = await request.json();
    console.log(controles)
    if (request.ok) {
        let listadoHtml = '';
            for (let control of controles) {
  
              let controlHtml = '<tr><td>'+ control.id +'</td><td>' + control.descripcion + '</td><td>' + control.tipo + '</td></tr>';
              listadoHtml += controlHtml;
              }
            
            
            document.querySelector('#controlesTBody').outerHTML = listadoHtml;
            
            }
    }

function redirectAgregarControl(){
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    formularioid = searchParams.get('formulario_id');
    window.location.href = `agregarControlesFormulario.html?formulario=${formularioid}`;
}