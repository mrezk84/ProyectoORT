$(document).ready(function() {
     getDocumentos()

    $('#documentos').DataTable();
});

async function getDocumentos() {

    let obraid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    obraid = searchParams.get('obra_id');

    const request = await fetch("http://localhost:5000/document/" + obraid, {
        method: 'GET',
    })

    const documentos = await request.json();
    console.log(documentos)
    if (request.ok) {
        let listadoHtml = '';
            for (let documento of documentos) {

                if (documento.status == "DONE"){
                    listadoHtml += '<tr bgcolor="#adebad">';
                }
                if (documento.status == "WIP"){
                    listadoHtml += '<tr bgcolor="#ffff99">';
                }
                if (documento.status == "TODO"){
                    listadoHtml += '<tr bgcolor="#ffb3b3">';
                }



              let documentoHtml = '<td>' + documento.formulario.nombre + '</td><td>' + documento.formulario.informacion + '</td><td>' + documento.piso.numero + '</td></tr>';
              listadoHtml += documentoHtml;
              }
            
            
            document.querySelector('#documentosTBody').outerHTML = listadoHtml;
            
            }
    }


function redirectCrearDocumento(){

    let obraid = null;
    const url = new URL(document.URL);
    const searchParams = url.searchParams;
    obraid = searchParams.get('obra_id');

    window.location.href = `altaDePlan.html?obra_id=${obraid}`;
}