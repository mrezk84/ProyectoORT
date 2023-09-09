$(document).ready(function() {
    buildControles()
});



let formulario = null;
async function buildControles() {
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    formularioid = searchParams.get('formulario_id');

    const request = await fetch("http://localhost:5000/controles" + formularioid, {
        method: 'GET',
    })

    const controles = await request.json();
    console.log(controles)
    if (request.ok) {
        let listadoHtml = '';
            for (let control of controles) {
  
              let botonEliminar = '<a href="#" class="btn btn-danger btn-circle btn-sm"><i class="fas fa-trash"></i></a>' ;
              let botonEditar = '<a href="#" btn btn-info btn-circle btn-sm"><i class="fas fa-info-circle"></i></a>  | ' ;
              let controlHtml = '<tr><td>'+ control.id +'</td><td>' + control.descripcion + '</td><td>' + control.tipo + '</td><td>'+'</td><td>'+ botonEditar 
              + botonEliminar + '</td></tr>';
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