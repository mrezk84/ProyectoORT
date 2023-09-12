$(document).ready(function() {
    getObrasByUser()

    $('#obras').DataTable();
});




async function getObrasByUser() {
    const request = await fetch("http://localhost:5000/obras", {
        method: 'GET',
    })

    const obras = await request.json();
    console.log(obras)
    if (request.ok) {
        let listadoHtml = '';
            for (let obra of obras) {
                let botondocumentos = '<a onclick = "redirectDocumentos('+obra.id+')"  class="btn btn-success btn-icon-split"><span class="icon text-white-50"><i class="fas fa-check"></i></span>documentos</a>';
                let obraHtml = '<tr><td>'+ obra.id +'</td><td>' + obra.nombre + '</td><td>' + botondocumentos + '</td></tr>';
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