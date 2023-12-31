$(document).ready(function() {
    getObrasByUser()

    $('#obras').DataTable();
});




async function getObrasByUser() {


    const request = await fetch("http://34.192.187.56:5000/obras", {
        method: 'GET',
    })

    const obras = await request.json();
    console.log(localStorage.getItem('email'))
    if (request.ok) {
        let listadoHtml = '';
            for (let obra of obras) {
                let botondocumentos = '<a onclick = "redirectDocumentos('+obra.id+')"  class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span>Documentos</a>';
                let botonUpdate = '<a onclick = "redirectUpdate('+obra.id+')"  class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span>Actualizar</a>';
                let botonDelete = '<a onclick = "eliminarObra('+obra.id+')"  class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span>Eliminar</a>';
                let obraHtml = '<tr><td>'+ obra.id +'</td><td>' + obra.nombre + '</td><td>' + botondocumentos + botonUpdate + botonDelete +'</td></tr>';
                listadoHtml += obraHtml;
        };

        document.querySelector('#obrasTBody').outerHTML = listadoHtml;

    }
}

function redirectPisos(id) {
    window.location.href = `altaSeguimiento.html?obra_id=${id}`;
}

function redirectDocumentos(id) {
    window.location.href = `documentosDeObra.html?obra_id=${id}`;
}

function redirectAltaObra() {
    window.location.href = `altaObra.html`;
}

async function redirectUpdate(id) {
    window.location.href = `obraUpdate.html?obra_id=${id}`;
}

async function eliminarObra(id){

    const request = await fetch("http://34.192.187.56:5000/obras/eliminar/"+ id, {
               method: 'DELETE',
               headers: {
                   'Accept': 'application/json',
                   'Content-Type': 'application/json'
               },
           })
         if (request.ok) {
             alert("Obra Eliminada");
         }else{
             alert("Error eliminando la Obra");
         }
         window.location.href = `obras.html`;
  }