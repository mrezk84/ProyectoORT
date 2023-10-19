$(document).ready(function() {
    buildControles()
});



let formulario = null;
async function buildControles() {
    let datos = {};
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    formularioid = searchParams.get('formulario_id');
    datos.id_formulario = Number(formularioid);

    const request = await fetch("http://34.192.187.56:5000/controles/byForm", {
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
  
              let botonDelete = '<a onclick = "eliminarControl('+control.id+')"  class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span>Eliminar</a>';
              let controlHtml = '<tr><td>'+ control.id +'</td><td>' + control.descripcion + '</td><td>' + control.tipo + '</td><td>' + botonDelete + '</td></tr>';
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

async function eliminarControl(id){
    let datos = {};
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    formularioid = searchParams.get('formulario_id');
    datos.formulario_id = Number(formularioid);

    const request = await fetch("http://localhost:5000/controles/"+ id +"/formulario", {
               method: 'DELETE',
               body: JSON.stringify(datos),
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.ok) {
             alert("control borrado del formulario correctamente");
         }else{
             alert("Error eliminando el control del formulario");
         }
}