$(document).ready(function() {
    getControles()

    $('#controles').DataTable();
});




async function getControles() {
    const request = await fetch("http://localhost:5000/controles", {
        method: 'GET',
    })

    const controles = await request.json();
    console.log(controles)
    if (request.ok) {
        let listadoHtml = '';
            for (let control of controles) {
  
              let botonEliminar = '<a onclick = "eliminarControl('+control.id+')"  class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span>Delete</a>';
              let botonEditar = '<a onclick = "redirectUpdate('+control.id+')"  class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span>Update</a>';
              let controlHtml = '<tr><td>'+ control.id +'</td><td>' + control.descripcion + '</td><td>' + control.tipo + '</td><td>'+'</td><td>'+ botonEditar 
              + botonEliminar + '</td></tr>';
              listadoHtml += controlHtml;
              }
            
            
            document.querySelector('#controlesTBody').outerHTML = listadoHtml;
            
            }
    }


function redirectCrearControl(){
    window.location.href = `altaControl.html`;
}

async function redirectUpdate(id) {
    window.location.href = `controlUpdate.html?control_id=${id}`;
}

async function eliminarControl(id){

    const request = await fetch("http://localhost:5000/controles/eliminar/"+ id, {
               method: 'DELETE',
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.ok) {
             alert("Control Eliminado");
         }else{
             alert("Error eliminando el control");
         }
         window.location.href = `controles.html`;
  }